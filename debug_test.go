package helpers

import (
	"strings"
	"testing"
)

func TestDebugAsPrettyJsonToBytes(t *testing.T) {
	t.Parallel()

	type payload struct {
		Name string `json:"name"`
		Num  int    `json:"num"`
	}
	out, errJSON, errMap := DebugAsPrettyJsonToBytes(payload{Name: "dev", Num: 7})
	if errJSON != nil || errMap != nil {
		t.Fatalf("unexpected errors: json=%v map=%v", errJSON, errMap)
	}
	text := string(out)
	if !strings.Contains(text, "\"name\": \"dev\"") || !strings.Contains(text, "\"num\": 7") {
		t.Fatalf("unexpected json output: %s", text)
	}
}
