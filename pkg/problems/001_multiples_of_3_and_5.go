package problems

import (
	"fmt"
)

type MultiplesOf3Or5 struct{}

func (p *MultiplesOf3Or5) ID() int {
	return 1
}

func (p *MultiplesOf3Or5) Text() string {
	return `If we list all the natural numbers below 10 that are multiples
of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
`
}

func (p *MultiplesOf3Or5) Solve() (string, error) {
	limit := 1000
	factors := []int{3, 5}

	muls := make(chan int, 100)

	go p.multiple_of(factors, muls)

	sum := 0
	for v := range muls {
		if v >= limit {
			break
		}
		sum += v
	}
	return fmt.Sprintf("%d", sum), nil
}

func (p *MultiplesOf3Or5) multiple_of(factors []int, muls chan int) {
	multiples := make([]chan int, len(factors))
	for i, factor := range factors {
		multiples[i] = make(chan int, 100)
		go func(n int, c chan int) {
			val := n
			for {
				c <- val
				val += n
			}
		}(factor, multiples[i])
	}

	current := make([]int, len(factors))
	// Prime the pump
	for i := 0; i < len(factors); i++ {
		current[i] = <-multiples[i]
	}

	for {
		min := current[0]
		max := current[0]
		for _, n := range current {
			if n > max {
				max = n
			}
			if n < min {
				min = n
			}
		}
		muls <- min

		for i := 0; i < len(factors); i++ {
			if current[i] == min {
				current[i] = <-multiples[i]
			}
		}
	}
}
