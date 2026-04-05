package helpers

import "testing"

func TestIsStruct(t *testing.T) {
	t.Parallel()

	type s struct{ A int }
	if !IsStruct(s{A: 1}) || !IsStruct(&s{A: 1}) {
		t.Fatalf("IsStruct expected true for struct and pointer")
	}
	var sp *s
	if IsStruct(sp) || IsStruct(1) {
		t.Fatalf("IsStruct expected false for nil pointer/non-struct")
	}
}
