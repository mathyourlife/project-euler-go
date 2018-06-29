package problems

import (
	"fmt"
)

type GoldbachsOtherConjecture struct{}

func (p *GoldbachsOtherConjecture) ID() int {
	return 46
}

func (p *GoldbachsOtherConjecture) Text() string {
	return `It was proposed by Christian Goldbach that every
odd composite number can be written as the sum of a prime and
twice a square.

9 = 7 + 2×1^2
15 = 7 + 2×2^2
21 = 3 + 2×3^2
25 = 7 + 2×3^2
27 = 19 + 2×2^2
33 = 31 + 2×1^2

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a
prime and twice a square?
`
}

func (p *GoldbachsOtherConjecture) Solve() (string, error) {

	solution := make(chan uint64)
	composites := make(chan uint64, 1)

	go func() {
		pg := NewPrimeGenerator()
		for n := uint64(2); n < uint64(6000); n++ {
			if !pg.IsPrime(n) && n%2 == 1 {
				composites <- n
			}
		}
		fmt.Println("closing composites")
		close(composites)
	}()

	go func() {
		pg := NewPrimeGenerator()
		for c := range composites {
			n := uint64(1)
			works := false
			for {
				square := n * n
				if 2*square > c {
					break
				}
				if pg.IsPrime(c - 2*square) {
					works = true
					break
				}
				n++
			}
			if !works {
				solution <- c
				defer close(solution)
				return
			}
		}
	}()

	return fmt.Sprintf("%d", <-solution), nil
}
