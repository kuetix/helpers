package helpers

import "testing"

func TestIntHelpers(t *testing.T) {
	t.Parallel()

	if n, ok := IsBytesIsInt([]byte("123")); !ok || n != 123 {
		t.Fatalf("expected integer parse, got n=%d ok=%v", n, ok)
	}
	if _, ok := IsBytesIsInt([]byte("x")); ok {
		t.Fatalf("expected non-int parse to fail")
	}

	if !IsNumeric(int64(10)) || !IsNumeric(3.14) || IsNumeric("10") {
		t.Fatalf("unexpected IsNumeric behavior")
	}
}
