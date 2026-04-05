package helpers

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"testing"
)

func TestStringHelpers(t *testing.T) {
	t.Parallel()

	if idx := FindStringIndex([]string{"a", "b"}, "b"); idx != 1 {
		t.Fatalf("FindStringIndex expected 1, got %d", idx)
	}
	if idx := FindStringIndex([]string{"a"}, "z"); idx != -1 {
		t.Fatalf("FindStringIndex missing expected -1, got %d", idx)
	}

	list := NaturalSort{"item2", "item10", "item1"}
	sort.Sort(list)
	if list[0] != "item1" || list[1] != "item2" || list[2] != "item10" {
		t.Fatalf("natural sort failed: %#v", list)
	}
	if !NaturalLess("a2", "a10") {
		t.Fatalf("NaturalLess numeric segment compare failed")
	}

	if got := EscapeChars("a+b", "+", "\\"); got != "a\\+b" {
		t.Fatalf("EscapeChars unexpected: %q", got)
	}
	if got := EscapeRedisValue("x-y"); got != "x\\-y" {
		t.Fatalf("EscapeRedisValue unexpected: %q", got)
	}

	if !IsString("x") || IsString(1) {
		t.Fatalf("unexpected IsString behavior")
	}
	if a, b := SplitString("a:b", ":"); a != "a" || b != "b" {
		t.Fatalf("SplitString failed: %q %q", a, b)
	}

	if got := ToCamelCase("my_service-name"); got != "MyServiceName" {
		t.Fatalf("ToCamelCase unexpected: %q", got)
	}
	if got := SanitizeServiceName("a.b-c/d\\e"); got != "a_b_c_d_e" {
		t.Fatalf("SanitizeServiceName unexpected: %q", got)
	}
	if got := ToSnakeCase("My HTTPService"); got != "my__h_t_t_p_service" {
		t.Fatalf("ToSnakeCase unexpected: %q", got)
	}

	expr, err := parser.ParseExpr("map[string]*pkg.Type")
	if err != nil {
		t.Fatalf("ParseExpr: %v", err)
	}
	if got := ExprToString(expr); got != "map[string]*pkg.Type" {
		t.Fatalf("ExprToString unexpected: %q", got)
	}
	if got := ExprToString(&ast.FuncType{}); got != "*ast.FuncType" {
		t.Fatalf("ExprToString fallback unexpected: %q", got)
	}

	exprIdent := &ast.Ident{Name: token.STRING.String()}
	if got := ExprToString(exprIdent); got != "STRING" {
		t.Fatalf("ExprToString ident unexpected: %q", got)
	}
}
