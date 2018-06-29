package problems

import (
	"testing"
)

func TestPrimeGet(t *testing.T) {
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

func TestPrimeGenerator_IsPrime(t *testing.T) {
	p := NewPrimeGenerator()
	if !p.IsPrime(19) {
		t.Errorf("did not determine that 19 was prime")
	}
	if p.IsPrime(100) {
		t.Errorf("thought that 100 was a prime")
	}
}

var primeResult uint64

func benchmarkPrime(i int, b *testing.B) {
	var r uint64
	for n := 0; n < b.N; n++ {
		pg = NewPrimeGenerator()
		// always record the result to prevent the compiler
		// eliminating the function call.
		r = pg.Get(i)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	primeResult = r
}

func BenchmarkPrime1(b *testing.B)     { benchmarkPrime(1, b) }
func BenchmarkPrime10(b *testing.B)    { benchmarkPrime(10, b) }
func BenchmarkPrime100(b *testing.B)   { benchmarkPrime(100, b) }
func BenchmarkPrime1000(b *testing.B)  { benchmarkPrime(1000, b) }
func BenchmarkPrime10000(b *testing.B) { benchmarkPrime(10000, b) }
