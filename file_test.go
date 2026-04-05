package helpers

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func mustEvalSymlinks(t *testing.T, p string) string {
	t.Helper()
	resolved, err := filepath.EvalSymlinks(p)
	if err != nil {
		return p
	}
	return resolved
}

func withChdir(t *testing.T, dir string) {
	t.Helper()
	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir %s: %v", dir, err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(old)
	})
}

func TestTouchFile(t *testing.T) {
	t.Parallel()

	t.Run("creates file when missing", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "created.txt")

		if err := TouchFile(path); err != nil {
			t.Fatalf("TouchFile: %v", err)
		}
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("stat touched file: %v", err)
		}
	})

	t.Run("returns nil when file exists", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "existing.txt")
		if err := os.WriteFile(path, []byte("content"), 0o644); err != nil {
			t.Fatalf("write file: %v", err)
		}

		before, err := os.Stat(path)
		if err != nil {
			t.Fatalf("stat before: %v", err)
		}
		beforeMod := before.ModTime()
		time.Sleep(10 * time.Millisecond)

		if err := TouchFile(path); err != nil {
			t.Fatalf("TouchFile existing: %v", err)
		}

		after, err := os.Stat(path)
		if err != nil {
			t.Fatalf("stat after: %v", err)
		}
		if after.ModTime().Before(beforeMod) {
			t.Fatalf("mod time went backwards: before=%v after=%v", beforeMod, after.ModTime())
		}
	})
}

func TestFprintfAndFprintln(t *testing.T) {
	t.Parallel()

	t.Run("writes content and returns bytes", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "out.txt")
		f, err := os.Create(path)
		if err != nil {
			t.Fatalf("create file: %v", err)
		}
		defer func(f *os.File) {
			err = f.Close()
			if err != nil && !errors.Is(err, fs.ErrClosed) {
				t.Fatalf("close file in defer: %v", err)
			}
		}(f)

		n1 := Fprintf(f, "hello %s", "world")
		n2 := Fprintln(f, "!")

		if n1 <= 0 {
			t.Fatalf("Fprintf returned %d", n1)
		}
		if n2 <= 0 {
			t.Fatalf("Fprintln returned %d", n2)
		}

		if err := f.Close(); err != nil {
			t.Fatalf("close file: %v", err)
		}

		b, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read file: %v", err)
		}
		if string(b) != "hello world!\n" {
			t.Fatalf("unexpected content: %q", string(b))
		}
	})

	t.Run("returns -1 on write error", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "closed.txt")
		f, err := os.Create(path)
		if err != nil {
			t.Fatalf("create file: %v", err)
		}
		if err := f.Close(); err != nil {
			t.Fatalf("close file: %v", err)
		}

		if got := Fprintf(f, "x"); got != -1 {
			t.Fatalf("Fprintf expected -1, got %d", got)
		}
		if got := Fprintln(f, "x"); got != -1 {
			t.Fatalf("Fprintln expected -1, got %d", got)
		}
	})
}

func TestFindFileUp(t *testing.T) {
	root := t.TempDir()
	level1 := filepath.Join(root, "a")
	level2 := filepath.Join(level1, "b")
	if err := os.MkdirAll(level2, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	target := filepath.Join(root, "marker.txt")
	if err := os.WriteFile(target, []byte("ok"), 0o644); err != nil {
		t.Fatalf("write marker: %v", err)
	}

	withChdir(t, level2)
	found, err := FindFileUp("marker.txt")
	if err != nil {
		t.Fatalf("FindFileUp: %v", err)
	}
	if mustEvalSymlinks(t, found) != mustEvalSymlinks(t, target) {
		t.Fatalf("expected %s, got %s", target, found)
	}

	_, err = FindFileUp("missing.txt")
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
}

func TestFindFileDown(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	nested := filepath.Join(dir, "x", "y")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	target := filepath.Join(nested, "target.txt")
	if err := os.WriteFile(target, []byte("ok"), 0o644); err != nil {
		t.Fatalf("write target: %v", err)
	}

	found, err := FindFileDown(dir, "target.txt")
	if err != nil {
		t.Fatalf("FindFileDown: %v", err)
	}
	if found != target {
		t.Fatalf("expected %s, got %s", target, found)
	}

	_, err = FindFileDown(dir, "missing.txt")
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
}

func TestGetModuleFromGoMod(t *testing.T) {
	t.Run("reads module from ancestor go.mod", func(t *testing.T) {
		root := t.TempDir()
		nested := filepath.Join(root, "pkg", "sub")
		if err := os.MkdirAll(nested, 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module github.com/example/project\n"), 0o644); err != nil {
			t.Fatalf("write go.mod: %v", err)
		}

		withChdir(t, nested)
		got, err := GetModuleFromGoMod()
		if err != nil {
			t.Fatalf("GetModuleFromGoMod: %v", err)
		}
		if got != "github.com/example/project" {
			t.Fatalf("unexpected module: %q", got)
		}
	})

	t.Run("returns os.ErrNotExist for commented or missing module", func(t *testing.T) {
		root := t.TempDir()
		if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("// module ignored\ngo 1.22\n"), 0o644); err != nil {
			t.Fatalf("write go.mod: %v", err)
		}
		withChdir(t, root)

		_, err := GetModuleFromGoMod()
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected os.ErrNotExist, got %v", err)
		}
	})
}

func TestGetRootPathGoMod(t *testing.T) {
	t.Run("returns directory containing nearest go.mod", func(t *testing.T) {
		root := t.TempDir()
		nested := filepath.Join(root, "a", "b")
		if err := os.MkdirAll(nested, 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module example.com/root\n"), 0o644); err != nil {
			t.Fatalf("write go.mod: %v", err)
		}

		withChdir(t, nested)
		got, err := GetRootPathGoMod()
		if err != nil {
			t.Fatalf("GetRootPathGoMod: %v", err)
		}
		if mustEvalSymlinks(t, got) != mustEvalSymlinks(t, root) {
			t.Fatalf("expected %s, got %s", root, got)
		}
	})

	t.Run("returns os.ErrNotExist when go.mod is not found", func(t *testing.T) {
		root := t.TempDir()
		withChdir(t, root)

		_, err := GetRootPathGoMod()
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected os.ErrNotExist, got %v", err)
		}
	})
}

func TestCleanSlashes_CurrentBehavior(t *testing.T) {
	t.Parallel()

	in := []string{"/a/", "b/", "/c"}
	got := CleanSlashes(in...)

	// Current implementation does not modify the incoming elements.
	if got[0] != "/a/" || got[1] != "b/" || got[2] != "/c" {
		t.Fatalf("unexpected CleanSlashes behavior: %#v", got)
	}
}
