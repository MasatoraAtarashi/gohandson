package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

// Walk のテスト
func TestWalk(t *testing.T) {
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 0
	for v := range ch {
		if v != expected[i] {
			t.Errorf("Walk failed: expected %d, got %d", expected[i], v)
		}
		i++
	}
	if i != len(expected) {
		t.Errorf("Walk output length mismatch: expected %d values, got %d", len(expected), i)
	}
}

// Same のテスト
func TestSame(t *testing.T) {
	// 同じ木（値が同じ）→ true を期待
	if !Same(tree.New(1), tree.New(1)) {
		t.Errorf("Same(tree.New(1), tree.New(1)) should be true")
	}

	// 異なる木（値が異なる）→ false を期待
	if Same(tree.New(1), tree.New(2)) {
		t.Errorf("Same(tree.New(1), tree.New(2)) should be false")
	}
}
