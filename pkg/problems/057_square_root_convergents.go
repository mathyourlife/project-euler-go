package problems

import (
	"fmt"
)

type SquareRootConvergents struct{}

func (p *SquareRootConvergents) ID() int {
	return 57
}

func (p *SquareRootConvergents) Text() string {
	return `It is possible to show that the square root of
two can be expressed as an infinite continued fraction.

√ 2 = 1 + 1/(2 + 1/(2 + 1/(2 + ... ))) = 1.414213...

By expanding this for the first four iterations, we get:

1 + 1/2 = 3/2 = 1.5
1 + 1/(2 + 1/2) = 7/5 = 1.4
1 + 1/(2 + 1/(2 + 1/2)) = 17/12 = 1.41666...
1 + 1/(2 + 1/(2 + 1/(2 + 1/2))) = 41/29 = 1.41379...

The next three expansions are 99/70, 239/169, and 577/408,
but the eighth expansion, 1393/985, is the first example
where the number of digits in the numerator exceeds the
number of digits in the denominator.

In the first one-thousand expansions, how many fractions
contain a numerator with more digits than denominator?
`
}

func (p *SquareRootConvergents) Solve() (string, error) {

	n := []int{1}
	d := []int{2}

	regroup := func(n []int) []int {
		for i, _ := range n {
			if i == len(n)-1 {
				break
			}
			n[i+1] += n[i] / 10
			n[i] %= 10
		}
		i := len(n) - 1
		for {
			if n[i] < 10 {
				break
			}
			n = append(n, n[i]/10)
			n[i] = n[i] % 10
			i++
		}
		return n
	}
	scale := func(n []int, factor int) []int {
		p := make([]int, len(n))
		for i, _ := range n {
			p[i] = n[i] * factor
		}
		return regroup(p)
	}
	sum := func(a, b []int) []int {
		l := len(a)
		if len(b) > l {
			l = len(b)
		}
		n := make([]int, l)
		for i := 0; i < l; i++ {
			if i < len(a) {
				n[i] += a[i]
			}
			if i < len(b) {
				n[i] += b[i]
			}
		}
		return regroup(n)
	}

	tally := 0
	for i := 2; i <= 1000; i++ {
		n = sum(n, scale(d, 2))
		n, d = d, n
		if len(sum(n, d)) > len(d) {
			tally++
		}
	}

	return fmt.Sprintf("%d", tally), nil
}
