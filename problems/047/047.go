/*
Distinct primes factors

The first two consecutive numbers to have two distinct prime factors are:
14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:
644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors.
What is the first of these numbers?
*/

package main

import (
	"github.com/mathyourlife/lt3maths/primefactorization"
	"log"
)

func main() {
	log.Println("Distinct primes factors")

	pf := primefactorization.NewPrimeFactorization()

	target := 4
	consecutive := 0
	n := uint64(2)

	// Should reduce the workload here by jumping n by the target size, but
	// current solution is "good enough" (comes in just about a minute)
SEARCH_LOOP:
	for {
		g := pf.Of(n)
		if len(g) == target {
			consecutive++
			if consecutive == target {
				log.Println(n - uint64(target) + uint64(1))
				break SEARCH_LOOP
			}
		} else {
			consecutive = 0
		}
		n++
	}
}
