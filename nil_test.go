package helpers

import "testing"

func TestIsNil(t *testing.T) {
	t.Parallel()

	if !IsNil(nil) {
		t.Fatalf("nil must be nil")
	}
	var m map[string]int
	if !IsNil(m) {
		t.Fatalf("nil map must be nil")
	}
	if IsNil(1) {
		t.Fatalf("int must not be nil")
	}
}
