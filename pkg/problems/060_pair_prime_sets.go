package problems

import (
	"fmt"
)

type PairPrimeSets struct{}

func (p *PairPrimeSets) ID() int {
	return 60
}

func (p *PairPrimeSets) Text() string {
	return `The primes 3, 7, 109, and 673, are quite remarkable.
By taking any two primes and concatenating them in any order
the result will always be prime. For example, taking 7 and
109, both 7109 and 1097 are prime. The sum of these four
primes, 792, represents the lowest sum for a set of four
primes with this property.

Find the lowest sum for a set of five primes for which any
two primes concatenate to produce another prime.
`
}

func (p *PairPrimeSets) Solve() (string, error) {

	scale := func(n int) uint64 {
		s := uint64(1)
		for i := uint64(0); i < uint64(n); i++ {
			s *= uint64(10)
		}
		return s
	}

	pairs := map[uint64]map[uint64]bool{}

	for i := 0; i < 2000; i++ {
		n := GetPrime(i)
		for m, _ := range pairs {
			if !IsPrime(n*scale(numDigits(m))+m) ||
				!IsPrime(m*scale(numDigits(n))+n) {
				continue
			}
			pairs[m][n] = true
		}
		pairs[n] = map[uint64]bool{}
	}

	loopLimit := 10000
	for a := 1; a < loopLimit; a++ {
		primeA := GetPrime(a)
		for b := a + 1; b < loopLimit; b++ {
			primeB := GetPrime(b)
			if !pairs[primeA][primeB] {
				continue
			}
			for c := b + 1; c < loopLimit; c++ {
				primeC := GetPrime(c)
				if !pairs[primeA][primeC] ||
					!pairs[primeB][primeC] {
					continue
				}
				for d := c + 1; d < loopLimit; d++ {
					primeD := GetPrime(d)
					if !pairs[primeA][primeD] ||
						!pairs[primeB][primeD] ||
						!pairs[primeC][primeD] {
						continue
					}
					for e := d + 1; e < loopLimit; e++ {
						primeE := GetPrime(e)
						if !pairs[primeA][primeE] ||
							!pairs[primeB][primeE] ||
							!pairs[primeC][primeE] ||
							!pairs[primeD][primeE] {
							continue
						}
						sum := primeA + primeB + primeC + primeD + primeE
						return fmt.Sprintf("%d", sum), nil
					}
				}
			}
		}
	}

	return fmt.Sprintf("%d", 0), nil
}
