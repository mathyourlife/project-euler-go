package problems

import (
	"fmt"
)

type PowerfulDigitSum struct{}

func (p *PowerfulDigitSum) ID() int {
	return 56
}

func (p *PowerfulDigitSum) Text() string {
	return `A googol (10^100) is a massive number: one
followed by one-hundred zeros; 100^100 is almost unimaginably
large: one followed by two-hundred zeros. Despite their
size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, a^b,
where a, b < 100, what is the maximum digital sum?
`
}

func (p *PowerfulDigitSum) Solve() (string, error) {

	numToSlice := func(n int) []int {
		slice := []int{}
		for {
			if n == 0 {
				break
			}
			slice = append(slice, n%10)
			n /= 10
		}
		return slice
	}

	scale := func(n []int, factor int) []int {
		for i, v := range n {
			n[i] = v * factor
		}
		return n
	}

	regroup := func(n []int) []int {
		r := make([]int, 0, len(n))
		carry := 0
		for _, v := range n {
			r = append(r, (v+carry)%10)
			carry = (v + carry) / 10
		}
		for {
			if carry == 0 {
				break
			}
			r = append(r, carry%10)
			carry /= 10
		}
		return r
	}

	maxSum := 0
	for a := 2; a < 100; a++ {
		n := numToSlice(a)
		for b := 2; b < 100; b++ {
			n = regroup(scale(n, a))
			sum := 0
			for _, v := range n {
				sum += v
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return fmt.Sprintf("%d", maxSum), nil
}
