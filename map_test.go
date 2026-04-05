package helpers

import "testing"

func TestMapHelpers(t *testing.T) {
	t.Parallel()

	type rec struct {
		Name string
		Age  int
	}
	m, err := ToMap(rec{Name: "ana", Age: 20})
	if err != nil {
		t.Fatalf("ToMap: %v", err)
	}
	if m["Name"] != "ana" || m["Age"] != 20 {
		t.Fatalf("unexpected map conversion: %#v", m)
	}

	mr, err := ToMapRecursive(3)
	if err != nil {
		t.Fatalf("ToMapRecursive: %v", err)
	}
	if mr["value"] != 3 {
		t.Fatalf("expected wrapped value map, got %#v", mr)
	}

	pathMap := map[string]interface{}{"a": map[string]interface{}{"b": 1}}
	if !IsPathExists(pathMap, []string{"a", "b"}) {
		t.Fatalf("path should exist")
	}
	if IsPathExists(pathMap, []string{"a", "c"}) {
		t.Fatalf("path should not exist")
	}

	root := map[string]interface{}{}
	n := MapKey(&root, "nest")
	n["x"] = "y"
	if root["nest"].(map[string]interface{})["x"] != "y" {
		t.Fatalf("MapKey did not create nested map")
	}

	mustPanic(t, func() {
		_ = MapPtrKey(&root, "bad")
	})

	decoded, err := DecodeToMap(rec{Name: "sam", Age: 30})
	if err != nil {
		t.Fatalf("DecodeToMap: %v", err)
	}
	if decoded["Name"] != "sam" {
		t.Fatalf("unexpected decoded value: %#v", decoded)
	}

	type obj struct {
		Keep string
		Drop string
	}
	merged := MergeObjectsToMap(map[string]interface{}{"Drop": true}, obj{Keep: "a", Drop: "x"}, obj{Keep: "b", Drop: "y"})
	if (*merged)["Keep"] != "b" {
		t.Fatalf("expected Keep to be overwritten, got %#v", *merged)
	}
	if (*merged)["Drop"] != "x" {
		t.Fatalf("expected Drop to stay excluded from second merge, got %#v", *merged)
	}
}
