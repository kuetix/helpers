package helpers

import "testing"

func TestCalculateMemoryUsage(t *testing.T) {
	t.Parallel()

	if got := CalculateMemoryUsage(nil, map[uintptr]bool{}); got != 0 {
		t.Fatalf("nil should be size 0, got %d", got)
	}
	if got := CalculateMemoryUsage("abc", map[uintptr]bool{}); got == 0 {
		t.Fatalf("string should report non-zero size")
	}

	type node struct {
		Next *node
		Val  int
	}
	n := &node{Val: 1}
	n.Next = n
	if got := CalculateMemoryUsage(n, map[uintptr]bool{}); got == 0 {
		t.Fatalf("cyclic pointer should still report size")
	}
}
