package helpers

import "testing"

func TestSliceHelpers(t *testing.T) {
	t.Parallel()

	if got := Len([]int{1, 2}); got != 2 {
		t.Fatalf("Len slice got %d", got)
	}
	if got := Len(10); got != -1 {
		t.Fatalf("Len default got %d", got)
	}

	if !IsSlice([]string{"x"}) || IsSlice("x") {
		t.Fatalf("unexpected IsSlice behavior")
	}

	if got := AppendStringUnique([]string{"a"}, "a"); len(got) != 1 {
		t.Fatalf("AppendStringUnique duplicated existing value: %#v", got)
	}
	if got := AppendUnique([]string{"a"}, []string{"a", "b", "c", "b"}); len(got) != 3 {
		t.Fatalf("AppendUnique unexpected result: %#v", got)
	}
}
