package helpers

import "testing"

func TestAsserts(t *testing.T) {
	t.Parallel()

	if !AssertInteger("gt", 5, 3) || AssertInteger("lt", 5, 3) {
		t.Fatalf("AssertInteger operators failed")
	}
	if AssertInteger("unknown", 1, 1) {
		t.Fatalf("AssertInteger should return false for unknown operator")
	}

	if !AssertString("eq", "a", "a") || !AssertString("ne", "a", "b") {
		t.Fatalf("AssertString operators failed")
	}
	if AssertString("bad", "a", "a") {
		t.Fatalf("AssertString should return false for unknown operator")
	}

	if got := AssertSwitch("ok", map[string]interface{}{"ok": "value"}); got != "value" {
		t.Fatalf("unexpected AssertSwitch value: %q", got)
	}
	if got := AssertSwitch("missing", map[string]interface{}{"ok": "value"}); got != "" {
		t.Fatalf("unexpected AssertSwitch fallback: %q", got)
	}
}
