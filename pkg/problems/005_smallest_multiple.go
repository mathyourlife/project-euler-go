package problems

import (
	"fmt"
)

type SmallestMultiple struct{}

func (p *SmallestMultiple) ID() int {
	return 5
}

func (p *SmallestMultiple) Text() string {
	return `2520 is the smallest number that can be divided by each of the
numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of
the numbers from 1 to 20?
`
}

func (p *SmallestMultiple) Solve() (string, error) {
	ns := make([]uint64, 20)
	for i := uint64(1); i <= 20; i++ {
		ns[i-1] = i
	}
	return fmt.Sprintf("%d", lcm(ns)), nil
}
