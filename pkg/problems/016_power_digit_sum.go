package problems

import (
	"fmt"
)

type PowerDigitSum struct{}

func (p *PowerDigitSum) ID() int {
	return 16
}

func (p *PowerDigitSum) Text() string {
	return `2^15 = 32768 and the sum of its digits is
3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
`
}

func (p *PowerDigitSum) Solve() (string, error) {
	n := NewBigInt(1)

	for i := 0; i < 1000; i++ {
		n.Mul(2)
		n.Regroup()
	}

	sum := 0
	for _, d := range n.n {
		sum += d
	}

	return fmt.Sprintf("%d", sum), nil
}
