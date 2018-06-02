package problems

import (
	"fmt"
)

type NonAbundantSums struct{}

func (p *NonAbundantSums) ID() int {
	return 23
}

func (p *NonAbundantSums) Text() string {
	return `A perfect number is a number for which the sum of its proper
divisors is exactly equal to the number. For example, the sum of the
proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means
that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is
less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the
smallest number that can be written as the sum of two abundant
numbers is 24. By mathematical analysis, it can be shown that all
integers greater than 28123 can be written as the sum of two
abundant numbers. However, this upper limit cannot be reduced any
further by analysis even though it is known that the greatest
number that cannot be expressed as the sum of two abundant numbers
is less than this limit.

Find the sum of all the positive integers which cannot be written
as the sum of two abundant numbers.
`
}

func (p *NonAbundantSums) Solve() (string, error) {

	max := uint64(30000)
	abundants := []uint64{}

	for n := uint64(1); n <= max; n++ {
		primeDivisors := properDivisors(n)
		s := uint64(0)
		for _, primeDivisor := range primeDivisors {
			s += primeDivisor
		}
		if s > n {
			abundants = append(abundants, n)
		}
	}

	// numbers that *can* be written as the sum of 2 abundant numbers
	// use a map in case the sum of 2 abundant numbers is duplicated
	abundantPairs := map[uint64]bool{}
	for i, _ := range abundants {
		for j := i; j < len(abundants); j++ {
			abundantPairs[abundants[i]+abundants[j]] = true
		}
	}

	// find numbers that are not in the list of the sums of 2
	// abundant numbers
	sum := uint64(0)
	for i := uint64(1); i <= max; i++ {
		if !abundantPairs[i] {
			sum += i
		}
	}

	return fmt.Sprintf("%d", sum), nil
}
