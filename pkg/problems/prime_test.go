package problems

import (
	"testing"
)

func TestGetPrime(t *testing.T) {
	p := GetPrime(0)
	if p != 2 {
		t.Errorf("first prime was not 2 got: %d", p)
	}
}

func TestPrimeGenerator(t *testing.T) {
	p := NewPrimeGenerator()
	thou := p.Get(999)
	if thou != 7919 {
		t.Errorf("the 1000th prime was calculated as %d instead of 7919", thou)
	}
}
