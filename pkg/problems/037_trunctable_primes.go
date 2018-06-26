package problems

import (
	"fmt"
)

type TrunctablePrimes struct{}

func (p *TrunctablePrimes) ID() int {
	return 37
}

func (p *TrunctablePrimes) Text() string {
	return `The number 3797 has an interesting property. Being prime
itself, it is possible to continuously remove digits from left to
right, and remain prime at each stage: 3797, 797, 97, and 7.
Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from
left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
`
}

func (p *TrunctablePrimes) Solve() (string, error) {

	reverse := func(n, base uint64) uint64 {
		r := uint64(0)
		for {
			if n == 0 {
				break
			}
			r = r*base + n%base
			n /= base
		}
		return r
	}

	tally := 0
	find := 11
	sum := uint64(0)

	// search the odds starting at 11
	n := uint64(11)
	for {
		right := n
		left := n

		trunctable := true
		for {
			if right == 0 {
				break
			}
			if !IsPrime(right) {
				trunctable = false
				break
			}
			if !IsPrime(left) {
				trunctable = false
				break
			}
			right /= 10
			left = reverse(reverse(left, 10)/10, 10)
		}
		if trunctable {
			sum += n
			tally++
			if tally == find {
				break
			}
		}
		n += 2
	}

	return fmt.Sprintf("%d", sum), nil
}
