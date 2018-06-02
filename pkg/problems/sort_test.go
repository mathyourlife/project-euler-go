package problems

import (
	"testing"
)

func TestLexPerm(t *testing.T) {
	a := []int{1, 0, 2}
	LexPerm(a)

	expected := []int{1, 2, 0}

	if len(expected) != len(a) {
		t.Errorf("unexpected lexicographic permutation: %v", a)
		return
	}

	for i := range expected {
		if a[i] != expected[i] {
			t.Errorf("unexpected lexicographic permutation: %v", a)
		}
	}
}
