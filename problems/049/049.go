/*
Prime permutations

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms
increases by 3330, is unusual in two ways:

(i) each of the three terms are prime, and,
(ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes,
exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three
terms in this sequence?
*/

package main

import (
	"github.com/mathyourlife/lt3maths/lexperm"
	"github.com/mathyourlife/lt3maths/prime"
	"log"
)

func main() {
	log.Println("Prime permutations")

	// Create the prime number generator
	p := prime.NewPrimeGenerator()

	// Instantiate a Lexicographic Permutation Iterator
	lp := lexperm.LexPerm{}

	// four digit primes
	fdp := map[int]bool{}

	prime := 0
GENERATE_PRIMES:
	for {
		prime = int(p.Next())
		if prime >= 1000 && prime < 10000 {
			fdp[prime] = true
		} else if prime >= 10000 {
			break GENERATE_PRIMES
		}
	}

	for n, _ := range fdp {
		digits := make([]int, 4)

		place := 10
		for i := 0; i < 4; i++ {
			digit := n - n/place*place
			digits[3-i] = digit * 10 / place
			n -= digit
			place *= 10
		}

		prime_perms := []int{}
		for {
			perm := 0
			for _, i := range digits {
				perm = perm*10 + i
			}
			if fdp[perm] {
				prime_perms = append(prime_perms, perm)
			}
			more := lp.Next(digits)
			if !more {
				break
			}
		}
		if len(prime_perms) == 3 && prime_perms[1]-prime_perms[0] == prime_perms[2]-prime_perms[1] {
			log.Println(prime_perms)
		}
	}
}
