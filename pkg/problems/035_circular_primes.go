package problems

import (
	"fmt"
	"math"
)

type CircularPrimes struct{}

func (p *CircularPrimes) ID() int {
	return 35
}

func (p *CircularPrimes) Text() string {
	return `The number, 197, is called a circular prime because all
rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100:
2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
`
}

func (p *CircularPrimes) Solve() (string, error) {

	rotate := func(n uint64) uint64 {
		mod := n % 10
		digits := numDigits(n)
		return (mod * uint64(math.Pow(10, float64(digits-1)))) + n/10
	}

	tally := 0
	for n := uint64(2); n < uint64(1000000); n++ {
		circular := true
		value := n
		for i := 0; i < numDigits(value); i++ {
			if !IsPrime(value) {
				circular = false
				break
			}
			value = rotate(value)
		}
		if circular {
			tally++
		}
	}
	return fmt.Sprintf("%d", tally), nil
}
