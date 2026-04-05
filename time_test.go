package helpers

import (
	"testing"
	"time"
)

func TestTimeHelpers(t *testing.T) {
	t.Parallel()

	tm := time.Date(2026, time.April, 5, 12, 0, 0, 0, time.UTC)
	if got := TimeToString(tm); got != "2026-04-05T12:00:00Z" {
		t.Fatalf("TimeToString unexpected: %q", got)
	}
	if _, err := time.Parse(time.RFC3339, NowAsString()); err != nil {
		t.Fatalf("NowAsString is not RFC3339: %v", err)
	}
}
