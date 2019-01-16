package problems

import (
	"fmt"
	"math"
)

type TotientPermutation struct{}

func (p *TotientPermutation) ID() int {
	return 70
}

func (p *TotientPermutation) Text() string {
	return `Euler's Totient function, φ(n) [sometimes called the phi function],
is used to determine the number of positive numbers less than or equal to n
which are relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are
all less than nine and relatively prime to nine, φ(9)=6.
The number 1 is considered to be relatively prime to every positive number,
so φ(1)=1.

Interestingly, φ(87109)=79180, and it can be seen that 87109 is a
permutation of 79180.

Find the value of n, 1 < n < 107, for which φ(n) is a permutation of n and
the ratio n/φ(n) produces a minimum.
`
}

// n/φ(n) is a minimum for the fewer distict prime factors
//
// φ(p) = p-1 a single prime won't work since x and x-1 aren't permatuans
// Value for a prime power argument
// φ(p) = p^(k-1)(p-1)
// φ(p1p2) = p1p2(1-1/p1)(1-1/p2) = (p1-1)(p2-1)
func (p *TotientPermutation) Solve() (string, error) {
	countDigits := func(n uint64) map[uint64]int {
		m := map[uint64]int{}
		for {
			m[n%10]++
			n /= 10
			if n == 0 {
				return m
			}
		}
		return nil
	}

	compare := func(a, b map[uint64]int) bool {
		if len(a) != len(b) {
			return false
		}
		for k, v := range a {
			if b[k] != v {
				return false
			}
		}
		return true
	}

	l := uint64(10000000)
	i := 0
	for {
		n := pg.Get(i)
		if n >= uint64(math.Sqrt(float64(l))) {
			break
		}
		i++
	}

	minRatio := 0.0
	minN := uint64(0)
	for a := i; a > 0; a-- {
		p1 := pg.Get(a)
		b := a + 1

		// To limit the search space, for each p1 and the smallest n/φ(n)
		// ratio found so far, calcuate the corresponding p2
		// r = n / φ(n)
		// r = p1 p2 / (p1 - 1) * (p2 - 1)
		// r = p1 / (p1 - 1) * p2 / (p2 - 1)
		//
		// Since p / (p - 1) > 1,
		// p1 / (p1-1) * p2 / (p2-1) must be greater than p1 / (p1-1)
		//
		// Stop search conditions:
		//   Since we're searching from sqrt(n) -> 2 for pi,
		//   in each loop, p1/(p1-1) gets bigger.  Tracking the minimum ratio
		//   allows us to stop the search once p1/(p1-1) is larger that the
		//   minimum ratio found so far.  Multiplying by any p2/(p2-1) will
		//   only increase the value, and continuing to search smaller primes
		//   for p1 will result in larger values of p1/(p1-1)
		if minRatio != 0 && float64(p1)/float64(p1-1) > minRatio {
			// end
			break
		}
		for {
			p2 := pg.Get(b)
			n := p1 * p2
			if n > l {
				// n has exceeded our upper limit in the search space
				break
			}
			phi := (p1 - 1) * (p2 - 1)
			if compare(countDigits(n), countDigits(phi)) {
				if minRatio == 0 || float64(n)/float64(phi) < minRatio {
					minRatio = float64(n) / float64(phi)
					minN = n
				}
			}
			b++
		}
	}

	return fmt.Sprintf("%d", minN), nil
}
