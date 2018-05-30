package problems

import (
	"testing"
)

func TestBigInt(t *testing.T) {
	b := NewBigInt(99999)
	b.Mul(2)
	b.Regroup()
	if b.Print() != "199998" {
		t.Errorf("BigInt didn't calculate a product correctly expected 199998, got: %s", b.Print())
	}
}
