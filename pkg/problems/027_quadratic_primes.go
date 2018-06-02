package problems

import (
	"fmt"
)

type QuadraticPrimes struct{}

func (p *QuadraticPrimes) ID() int {
	return 27
}

func (p *QuadraticPrimes) Text() string {
	return `Euler discovered the remarkable quadratic formula:

n² + n + 41

It turns out that the formula will produce 40 primes for the
consecutive values n = 0 to 39. However, when n = 40,
40² + 40 + 41 = 40(40 + 1) + 41
is divisible by 41, and certainly when n = 41,
41² + 41 + 41 is clearly divisible by 41.

The incredible formula  n² − 79n + 1601 was discovered, which
produces 80 primes for the consecutive values n = 0 to 79.
The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

    n² + an + b, where |a| < 1000 and |b| < 1000

    where |n| is the modulus/absolute value of n
    e.g. |11| = 11 and |−4| = 4

Find the product of the coefficients, a and b, for the
quadratic expression that produces the maximum number of
primes for consecutive values of n, starting with n = 0.
`
}

func (p *QuadraticPrimes) Solve() (string, error) {
	var args []int
	max := 0
	i := 0
	for {
		// for n² + an + b, if n=0 then b needs to be a prime number
		b := GetPrime(i)
		if b > 1000 {
			break
		}
		for a := -999; a < 1000; a++ {
			// for each |a| < 1000, determine the number of valid
			// consecutive primes, the quadratic n² + an + b can generate
			validLen := p.CheckQuadraticPrimeGenerator(a, int(b))
			if validLen > max {
				// found a new set of coefficents that generate
				// more prime numbers
				max = validLen
				args = []int{a, int(b)}
			}
		}
		i++
	}
	return fmt.Sprintf("%d", args[0]*args[1]), nil
}

func (p *QuadraticPrimes) CheckQuadraticPrimeGenerator(a int, b int) int {
	n := 0
	for {
		candidate := n*n + a*n + b
		if candidate < 2 {
			break
		}
		if !IsPrime(uint64(candidate)) {
			break
		}
		n++
	}
	return n
}
