package problems

import (
	"fmt"
)

type ConsecutivePrimeSum struct{}

func (p *ConsecutivePrimeSum) ID() int {
	return 50
}

func (p *ConsecutivePrimeSum) Text() string {
	return `The prime 41, can be written as the sum of six
consecutive primes:
41 = 2 + 3 + 5 + 7 + 11 + 13

This is the longest sum of consecutive primes that adds to a
prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a
prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the
most consecutive primes?
`
}

func (p *ConsecutivePrimeSum) Solve() (string, error) {

	limit := uint64(1000000)

	solution := uint64(0)
	maxCount := 0
	n := 0
	for {
		if GetPrime(n) > limit {
			break
		}
		sum := uint64(0)
		count := n
		for {
			sum += GetPrime(count)
			if sum > limit {
				break
			}
			if IsPrime(sum) {
				if count+1-n > maxCount {
					maxCount = count + 1 - n
					solution = sum
				}
			}
			count++
		}
		n++
	}

	return fmt.Sprintf("%d", solution), nil
}
