package helpers

import "testing"

func TestIsPointer(t *testing.T) {
	t.Parallel()

	x := 1
	if !IsPointer(&x) || IsPointer(x) {
		t.Fatalf("unexpected IsPointer behavior")
	}
	mustPanic(t, func() {
		_ = IsPointer(nil)
	})
}
