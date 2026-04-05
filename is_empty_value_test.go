package helpers

import (
	"reflect"
	"testing"
)

func TestEmptyValueHelpers(t *testing.T) {
	t.Parallel()

	if !IsEmptyValue(nil) {
		t.Fatalf("nil should be empty")
	}
	if !IsEmptyValue("") {
		t.Fatalf("empty string should be empty")
	}
	if IsEmptyValue(1) {
		t.Fatalf("non-zero int should not be empty")
	}

	type s struct{ N int }
	if IsEmptyValue(s{}) {
		t.Fatalf("current behavior reports zero-value structs as non-empty")
	}

	if !IsEmptyReflectValue(reflect.ValueOf(false)) {
		t.Fatalf("false bool should be empty")
	}
}
