package problems

import (
	"fmt"
)

type PandigitalProducts struct{}

func (p *PandigitalProducts) ID() int {
	return 32
}

func (p *PandigitalProducts) Text() string {
	return `We shall say that an n-digit number is pandigital if it makes
use of all the digits 1 to n exactly once; for example, the
5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254,
containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product
identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be
sure to only include it once in your sum.
`
}

func (p *PandigitalProducts) Solve() (string, error) {

	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	split := func(i, j int) int {
		n := 0
		for k := i; k < j; k++ {
			n *= 10
			n += d[k]
		}
		return n
	}

	products := map[int]bool{}
	var a, b, c, prod int
	for {
		for i := 1; i < len(d)-1; i++ {
			for j := i + 1; j < len(d); j++ {
				a = split(0, i)
				b = split(i, j)
				c = split(j, len(d))
				prod = a * b
				if prod == c {
					products[c] = true
				} else if prod > c {
					break
				}
			}
		}
		if !LexPerm(d) {
			break
		}
	}
	sum := 0
	for k, _ := range products {
		sum += k
	}
	return fmt.Sprintf("%d", sum), nil
}
