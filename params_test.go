package helpers

import "testing"

func TestGetFunctionOptions(t *testing.T) {
	t.Parallel()

	if got := GetFunctionOptions("k", "def", map[string]interface{}{"k": "v"}); got != "v" {
		t.Fatalf("unexpected function option: %v", got)
	}
	if got := GetFunctionOptions("missing", "def", map[string]interface{}{"k": "v"}); got != "def" {
		t.Fatalf("expected default option: %v", got)
	}
}
