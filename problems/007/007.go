/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
package main

import (
	"fmt"
	"math"
)

type PrimeGenerator struct {
	primes    []uint64
	current   []uint64
	n         uint64
	threshold uint64
}

func NewPrimeGenerator() *PrimeGenerator {
	return &PrimeGenerator{
		primes:    make([]uint64, 0),
		current:   make([]uint64, 0),
		n:         2,
		threshold: uint64(math.Sqrt(2)),
	}
}

func (p *PrimeGenerator) Next() uint64 {
	if p.n == 2 {
		p.primes = append(p.primes, 2)
		p.current = append(p.current, 2)
		p.n++
		return 2
	}

	for {
		t := p.loop_next()
		if t {
			break
		}
	}
	return p.n - 2
}

func (p *PrimeGenerator) loop_next() bool {
	prime := true
CurrentLoop:
	for i, c := range p.current {
		for {
			if c == p.n {
				prime = false
				p.current[i] = c
				break CurrentLoop

			} else if p.primes[i] > p.threshold {
				break CurrentLoop

			} else if c < p.n {
				c += p.primes[i]
				continue

			} else {
				p.current[i] = c
				break
			}
		}
	}
	if prime {
		p.threshold = uint64(math.Sqrt(float64(p.n))) + 1
		p.primes = append(p.primes, p.n)
		p.current = append(p.current, p.n)
	}
	p.n += 2
	return prime
}

func main() {

	var n uint64
	nth := 10001

	pg := NewPrimeGenerator()

	for i := 0; i < nth; i++ {
		n = pg.Next()
	}
	fmt.Printf("%dth prime number is %d\n", nth, n)
}
