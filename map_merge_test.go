package helpers

import "testing"

func TestMapMergeHelpers(t *testing.T) {
	t.Parallel()

	m := MergeMapsLevel0(map[string]interface{}{"a": 1}, map[string]interface{}{"a": 2, "b": 3})
	if m["a"] != 2 || m["b"] != 3 {
		t.Fatalf("unexpected MergeMapsLevel0 result: %#v", m)
	}

	dst := map[string]interface{}{
		"m": map[string]interface{}{"x": 1},
		"s": []string{"a"},
	}
	src := map[string]interface{}{
		"m": map[string]interface{}{"y": 2},
		"s": []interface{}{"b"},
	}
	got := MergeMaps(dst, src)
	if gotM := got["m"].(map[string]interface{}); gotM["x"] != 1 || gotM["y"] != 2 {
		t.Fatalf("unexpected nested MergeMaps result: %#v", got)
	}
	if gotS := got["s"].([]string); len(gotS) != 2 || gotS[1] != "b" {
		t.Fatalf("unexpected slice MergeMaps result: %#v", gotS)
	}

	base := map[string]interface{}{"k": map[string]interface{}{"n": 1}}
	UpdateMap(&base, map[string]interface{}{"k": map[string]interface{}{"m": 2}})
	if nested := base["k"].(map[string]interface{}); nested["n"] != 1 || nested["m"] != 2 {
		t.Fatalf("unexpected UpdateMap result: %#v", base)
	}

	var nilDst *map[string]interface{}
	u := UpdateMaps(nilDst, &map[string]interface{}{"a": 1}, &map[string]interface{}{"b": 2})
	if (*u)["a"] != 1 || (*u)["b"] != 2 {
		t.Fatalf("unexpected UpdateMaps result: %#v", *u)
	}
}
