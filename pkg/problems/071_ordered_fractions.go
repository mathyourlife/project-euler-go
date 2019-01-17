package problems

import (
	"fmt"
)

type OrderedFractions struct{}

func (p *OrderedFractions) ID() int {
	return 71
}

func (p *OrderedFractions) Text() string {
	return `Consider the fraction, n/d, where n and d are positive integers.
If n<d and HCF(n,d)=1, it is called a reduced proper fraction.

If we list the set of reduced proper fractions for d ≤ 8 in ascending
order of size, we get:

1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8

It can be seen that 2/5 is the fraction immediately to the left of 3/7.

By listing the set of reduced proper fractions for d ≤ 1,000,000 in
ascending order of size, find the numerator of the fraction immediately
to the left of 3/7.
`
}

// compare fractions
// if n2/d2 is larger return 1
// if n2/d2 is smaller return -1
// return 0 if equal
func (p *OrderedFractions) compareFractions(n1, d1, n2, d2 int) int {
	left := n1 * d2
	right := n2 * d1
	if left < right {
		return 1
	} else if left > right {
		return -1
	}
	return 0
}

func (p *OrderedFractions) Solve() (string, error) {
	limit := 1000000
	target := []int{0, 1}
	for d := 2; d <= limit; d++ {
		// log.Printf("**********  %d", p.compareFractions(3, 7, d*3/7+1, d))
		// log.Printf("3/7 of %d = %d %g", d, d*3/7+1, float64(d)*3.0/7.0)
		for n := d*3/7 + 1; n > 0; n-- {
			// log.Printf("%d/%d", n, d)
			compare := p.compareFractions(n, d, target[0], target[1])
			if compare == 1 {
				break
			}
			if p.compareFractions(n, d, 3, 7) == 1 && compare == -1 {
				target = []int{n, d}
			}
		}
	}

	return fmt.Sprintf("%d", target[0]), nil
}
