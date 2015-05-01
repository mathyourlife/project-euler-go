/*
If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import (
	"fmt"
)

func next(factors []int, current []int) (int, []int) {

	idxs := make([]int, 0)
	min_val := -1
	for idx, factor := range current {
		if min_val == -1 || min_val > factor {
			min_val = factor
			idxs = []int{idx}
		} else if min_val == factor {
			idxs = append(idxs, idx)
		}
	}
	next_val := current[idxs[0]]
	for _, idx := range idxs {
		current[idx] += factors[idx]
	}
	return next_val, current
}

func multiple_of(factors []int, muls chan int) {
	var next_val int
	current := make([]int, 0)
	for _, factor := range factors {
		current = append(current, factor)
	}

	for {
		next_val, current = next(factors, current)
		muls <- next_val
	}
}

func main() {
	limit := 10
	factors := []int{3, 5}

	muls := make(chan int)

	go multiple_of(factors, muls)

	sum := 0
	for v := range muls {
		if v >= limit {
			break
		}
		sum += v
	}
	fmt.Printf("The sum of all the multiples of %v below %d is %d\n", factors, limit, sum)
}
