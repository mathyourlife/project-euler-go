package problems

import (
	"fmt"
)

type FactorialDigitSum struct{}

func (p *FactorialDigitSum) ID() int {
	return 20
}

func (p *FactorialDigitSum) Text() string {
	return `n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
`
}

func (p *FactorialDigitSum) Solve() (string, error) {
	N := 100

	b := NewBigInt(1)

	for i := 1; i <= N; i++ {
		b.Mul(i)
		b.Regroup()
	}
	sum := 0
	for _, d := range b.n {
		sum += d
	}
	return fmt.Sprintf("%d", sum), nil
}
