package problems

import (
	"fmt"
)

type TenThousandFirstPrime struct{}

func (p *TenThousandFirstPrime) ID() int {
	return 7
}

func (p *TenThousandFirstPrime) Text() string {
	return `By listing the first six prime numbers: 2, 3, 5, 7, 11,
and 13, we can see that the 6th prime is 13.

What is the 10,001st prime number?
`
}

func (p *TenThousandFirstPrime) Solve() (string, error) {
	return fmt.Sprintf("%d", pg.Get(10000)), nil
}
