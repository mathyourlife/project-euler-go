package problems

import (
	"fmt"
)

type SpecialPythagoreanTriplet struct{}

func (p *SpecialPythagoreanTriplet) ID() int {
	return 9
}

func (p *SpecialPythagoreanTriplet) Text() string {
	return `A Pythagorean triplet is a set of three natural
numbers, a < b < c, for which,

a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
`
}

func (p *SpecialPythagoreanTriplet) Solve() (string, error) {
	var a, b, c int
	N := 1000
Search:
	for c = N; c > 1; c-- {
		for b = 1; b <= c; b++ {
			if b+c >= N {
				break
			}
			a = N - b - c
			if a > b {
				continue
			}
			if a*a+b*b == c*c {
				break Search
			}
		}
	}
	return fmt.Sprintf("%d", a*b*c), nil
}
