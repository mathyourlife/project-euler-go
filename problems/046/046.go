/*
Goldbach's other conjecture

It was proposed by Christian Goldbach that every odd composite number
can be written as the sum of a prime and twice a square.

9 = 7 + 2×1^2
15 = 7 + 2×2^2
21 = 3 + 2×3^2
25 = 7 + 2×3^2
27 = 19 + 2×2^2
33 = 31 + 2×1^2

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a
prime and twice a square?
*/

package main

import (
	"github.com/mathyourlife/lt3maths/prime"
	"log"
)

func NewDoubleSquares() (*DoubleSquares, error) {
	return &DoubleSquares{
		n:      make(chan uint64, 0),
		Values: make([]uint64, 0),
	}, nil
}

type DoubleSquares struct {
	n      chan uint64
	Values []uint64
}

func (ds *DoubleSquares) Generate() {
	i := uint64(1)
	var val uint64
	for {
		val = uint64(2) * i * i
		ds.Values = append(ds.Values, val)
		ds.n <- val
		i++
	}
}

func main() {
	log.Println("Goldbach's other conjecture")

	// Create the double-square generator
	ds, err := NewDoubleSquares()
	if err != nil {
		panic(err)
	}
	go ds.Generate()

	// Create the prime number generator
	p := prime.NewPrimeGenerator()

	ds_val := <-ds.n
	p_val := p.Next()
	p_val = p.Next()

	// Keep track of all the odds that can be written as the
	// sum of a prime and twice a square
	odds := map[uint64]bool{}

	// Track the next candidate for an odd composite
	odd_comp_candidate := uint64(3)

	// Double Squares: 2, 8, 18, 32, ...
	// Primes:         3, 5, 7, 11, ...
	// Use the generators to track either the next double square or prime
	// merging the series: 2, 3, 5, 7, 8, 11, ...
	// If the next value is a double square, track the sum of the new
	// double square and all previous primes.
SEARCH_LOOP:
	for {

		if ds_val < p_val {
			// Next value is a double square
			for _, n := range p.Primes {
				odds[n+ds_val] = true
			}
			ds_val = <-ds.n
		} else {
			// Next value is a prime
			for _, n := range ds.Values {
				odds[n+p_val] = true
			}
			for {
				if odd_comp_candidate > p_val {
					break
				}
				if odd_comp_candidate < p_val {
					if !odds[odd_comp_candidate] {
						log.Println(odd_comp_candidate)
						break SEARCH_LOOP
					}
				}
				odd_comp_candidate += 2
			}
			p_val = p.Next()
		}
	}
}
