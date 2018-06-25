package problems

import (
	"reflect"
	"testing"
)

func TestRemoveFactor(t *testing.T) {
	tests := []struct {
		dividend uint64
		divisor  uint64
		status   bool
		quotient uint64
	}{
		{10, 2, true, 5},
		{10, 1, true, 10},
		{1, 10, false, 1},
	}

	for _, test := range tests {
		ok, res := removeFactor(test.dividend, test.divisor)
		if ok != test.status {
			t.Errorf("%d/%d didn't evaluate the status correctly", test.dividend, test.divisor)
		}
		if test.quotient != res {
			t.Errorf("%d/%d didn't evaluate the quotient correctly", test.dividend, test.divisor)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		n  uint64
		pf map[uint64]int
	}{
		{0, map[uint64]int{}},
		{1, map[uint64]int{}},
		{2, map[uint64]int{2: 1}},
		{12, map[uint64]int{2: 2, 3: 1}},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.pf, primeFactors(test.n)) {
			t.Errorf("incorrect prime factorization of %d to %v", test.n, test.pf)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		ns  []uint64
		lcm uint64
	}{
		{[]uint64{1, 1}, 1},
		{[]uint64{2, 3}, 6},
		{[]uint64{12, 45}, 180},
	}

	for _, test := range tests {
		v := lcm(test.ns)
		if test.lcm != v {
			t.Errorf("incorrect lcm %d of %v", v, test.ns)
		}
	}
}

func TestGCF(t *testing.T) {
	tests := []struct {
		ns  []uint64
		gcf uint64
	}{
		{[]uint64{10, 25}, 5},
		{[]uint64{25, 10}, 5},
		{[]uint64{27, 10}, 1},
		{[]uint64{36, 27}, 9},
		{[]uint64{36}, 36},
	}

	for _, test := range tests {
		v := gcf(test.ns)
		if test.gcf != v {
			t.Errorf("incorrect gcf %d of %v", v, test.ns)
		}
	}
}
