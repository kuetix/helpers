package helpers

import (
	"errors"
	"testing"
)

func TestMustHelpers(t *testing.T) {
	t.Parallel()

	if v, typ := MustInt("12"); v != 12 || typ != "string" {
		t.Fatalf("MustInt string parse failed: v=%d typ=%s", v, typ)
	}
	if v, typ := MustInt(nil, 9); v != 9 || typ != "nil" {
		t.Fatalf("MustInt nil default behavior changed: v=%d typ=%s", v, typ)
	}
	if v, typ := MustInt("bad", 8); v != 8 || typ != "default" {
		t.Fatalf("MustInt default fallback failed: v=%d typ=%s", v, typ)
	}

	errVal := errors.New("boom")
	if s, typ := MustString(errVal); s != "boom" || typ != "error" {
		t.Fatalf("MustString error conversion failed: s=%q typ=%s", s, typ)
	}
	if s, typ := MustString(7); s != "7" || typ != "int" {
		t.Fatalf("MustString int conversion failed: s=%q typ=%s", s, typ)
	}

	if b, typ := MustBool(false); b || typ != "bool" {
		t.Fatalf("MustBool bool conversion failed: b=%v typ=%s", b, typ)
	}
	if b, typ := MustBool(nil, true); b || typ != "nil" {
		t.Fatalf("MustBool nil behavior changed: b=%v typ=%s", b, typ)
	}

	arr, typ := MustArray("x", []interface{}{})
	if typ != "array" || len(arr) != 1 || arr[0] != "x" {
		t.Fatalf("MustArray string conversion failed: %#v (%s)", arr, typ)
	}
	arr, typ = MustArray(10, []interface{}{"fallback"})
	if typ != "default" || len(arr) != 1 || arr[0] != "fallback" {
		t.Fatalf("MustArray default fallback failed: %#v (%s)", arr, typ)
	}
}
