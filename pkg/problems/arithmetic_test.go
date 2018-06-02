package problems

import (
	"sort"
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

func TestBigInt_AddBigInt(t *testing.T) {
	b := NewBigInt(39)
	b.AddBigInt(NewBigInt(4123))
	if b.Print() != "4162" {
		t.Errorf("sum of big ints is not correct: %s", b.Print())
	}
}

func TestDivisors(t *testing.T) {
	d := []int{}

	for _, v := range divisors(30) {
		d = append(d, int(v))
	}
	sort.Ints(d)

	expected := []int{1, 2, 3, 5, 6, 10, 15, 30}

	if len(expected) != len(d) {
		t.Errorf("unexpected divisors for 30: %v", d)
		return
	}

	for i := range expected {
		if d[i] != expected[i] {
			t.Errorf("unexpected divisor for 30: %d", expected[i])
		}
	}
}

func TestProperDivisors(t *testing.T) {
	d := []int{}

	for _, v := range properDivisors(30) {
		d = append(d, int(v))
	}
	sort.Ints(d)

	expected := []int{1, 2, 3, 5, 6, 10, 15}

	if len(expected) != len(d) {
		t.Errorf("unexpected proper divisors for 30: %v", d)
		return
	}

	for i := range expected {
		if d[i] != expected[i] {
			t.Errorf("unexpected proper divisor for 30: %d", expected[i])
		}
	}
}
