package helpers

import (
	"reflect"
	"testing"
)

func TestObjectHelpers(t *testing.T) {
	type record struct {
		Name string
		Age  int
		Any  any
	}

	r := &record{Name: "old", Age: 1}
	if v, ok := FieldValue(r, "Name"); !ok || v != "old" {
		t.Fatalf("FieldValue struct failed: v=%v ok=%v", v, ok)
	}
	mv := map[string]interface{}{"Name": "map"}
	if v, ok := FieldValue(mv, "Name"); !ok || v != "map" {
		t.Fatalf("FieldValue map failed: v=%v ok=%v", v, ok)
	}

	if v, ok := SetFieldValueString(r, "Name", "new"); !ok || v != "new" || r.Name != "new" {
		t.Fatalf("SetFieldValueString failed: v=%v ok=%v rec=%+v", v, ok, r)
	}
	if v, ok := SetFieldValueInt(r, "Age", 9); !ok || v != 9 || r.Age != 9 {
		t.Fatalf("SetFieldValueInt failed: v=%v ok=%v rec=%+v", v, ok, r)
	}
	if v, ok := SetFieldValue(r, "Any", 123); !ok || v != 123 || r.Any != 123 {
		t.Fatalf("SetFieldValue failed: v=%v ok=%v rec=%+v", v, ok, r)
	}

	mustPanic(t, func() {
		_, _ = SetFieldValueString(map[string]interface{}{"k": "v"}, "k", "x")
	})
	mustPanic(t, func() {
		_, _ = SetFieldValueInt(map[string]interface{}{"k": 1}, "k", 2)
	})
	mustPanic(t, func() {
		_, _ = SetFieldValue(map[string]interface{}{"k": 1}, "k", 2)
	})

	clone, isPtr := CloneOf(r)
	if !isPtr {
		t.Fatalf("expected pointer input to report isPtr=true")
	}
	cl, ok := clone.(record)
	if !ok || cl.Name != r.Name || cl.Age != r.Age {
		t.Fatalf("unexpected CloneOf pointer clone: %#v", clone)
	}

	clone2, isPtr2 := CloneOf(record{Name: "x", Age: 2})
	if isPtr2 {
		t.Fatalf("expected non-pointer input to report isPtr=false")
	}
	if reflect.ValueOf(clone2).IsNil() {
		t.Fatalf("CloneOf non-pointer should not be nil")
	}
}
