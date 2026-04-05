package helpers

import "testing"

func TestArgsReorg(t *testing.T) {
	t.Parallel()

	got := ArgsReorg("-v", "input.txt", "other")
	if len(got) != 2 || got[0] != "-v" || got[1] != "input.txt" {
		t.Fatalf("unexpected reordered args: %#v", got)
	}
}
