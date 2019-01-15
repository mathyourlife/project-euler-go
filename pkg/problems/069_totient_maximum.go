package problems

import (
	"fmt"
)

type TotientMaxium struct {
	pfs map[uint64]map[uint64]int
}

func (p *TotientMaxium) ID() int {
	return 69
}

func (p *TotientMaxium) Text() string {
	return `Euler's Totient function, φ(n) [sometimes called the phi function],
is used to determine the number of numbers less than n which are relatively
prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all less than nine and
relatively prime to nine, φ(9)=6.

n 	Relatively Prime 	φ(n)  n/φ(n)
2 	1                 1     2
3 	1,2               2     1.5
4 	1,3               2     2
5 	1,2,3,4           4     1.25
6 	1,5               2     3
7 	1,2,3,4,5,6       6     1.1666...
8 	1,3,5,7           4     2
9 	1,2,4,5,7,8       6     1.5
10 	1,3,7,9           4     2.5

It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10.

Find the value of n ≤ 1,000,000 for which n/φ(n) is a maximum.
`
}

// In order to maximize n/φ(n), we want a large number with
// the fewest relative primes.  To achieve this, a prime factorization
// of n should have the most possible distinct prime factors such as:
//
// 6    = 2*3
// 30   = 2*3*5
// 210  = 2*3*5*7
// 2310 = 2*3*5*7*11
//
func (p *TotientMaxium) Solve() (string, error) {
	l := uint64(1000000)
	total := uint64(1)
	n := 0
	for {
		p := pg.Get(n)
		if p*total > l {
			break
		}
		total *= p
		n++
	}

	return fmt.Sprintf("%d", total), nil
}
