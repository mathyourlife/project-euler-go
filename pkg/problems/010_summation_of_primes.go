package problems

import (
	"fmt"
)

type SummationOfPrimes struct{}

func (p *SummationOfPrimes) ID() int {
	return 10
}

func (p *SummationOfPrimes) Text() string {
	return `The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
`
}

func (p *SummationOfPrimes) Solve() (string, error) {
	s := uint64(0)
	i := 0
	limit := uint64(2000000)
	for {
		n := pg.Get(i)
		if n > limit {
			break
		}
		s += n
		i++
	}
	return fmt.Sprintf("%d", s), nil
}
