package problems

import (
	"fmt"
)

type LychrelNumbers struct{}

func (p *LychrelNumbers) ID() int {
	return 55
}

func (p *LychrelNumbers) Text() string {
	return `If we take 47, reverse and add,
47 + 74 = 121, which is palindromic.

Not all numbers produce palindromes so quickly. For example,

349 + 943 = 1292,
1292 + 2921 = 4213
4213 + 3124 = 7337

That is, 349 took three iterations to arrive at a palindrome.

Although no one has proved it yet, it is thought that some
numbers, like 196, never produce a palindrome. A number that
never forms a palindrome through the reverse and add process
is called a Lychrel number. Due to the theoretical nature of
these numbers, and for the purpose of this problem, we shall
assume that a number is Lychrel until proven otherwise. In
addition you are given that for every number below ten-thousand,
it will either (i) become a palindrome in less than fifty
iterations, or, (ii) no one, with all the computing power that
exists, has managed so far to map it to a palindrome. In fact,
10677 is the first number to be shown to require over fifty
iterations before producing a palindrome:
4668731596684224866951378664 (53 iterations, 28-digits).

Surprisingly, there are palindromic numbers that are
themselves Lychrel numbers; the first example is 4994.

How many Lychrel numbers are there below ten-thousand?
`
}

func (p *LychrelNumbers) Solve() (string, error) {
	numToSlice := func(n uint64) []int {
		slice := []int{}
		for {
			if n == 0 {
				break
			}
			slice = append(slice, int(n%10))
			n /= 10
		}
		return slice
	}

	flipSum := func(n []int) []int {
		sum := make([]int, len(n))
		for i, v := range n {
			sum[len(n)-i-1] = v
		}
		for i, v := range n {
			sum[i] += v
		}
		return sum
	}

	regroup := func(n []int) []int {
		r := make([]int, 0, len(n))
		carry := 0
		for _, v := range n {
			r = append(r, (v+carry)%10)
			carry = (v + carry) / 10
		}
		if carry > 0 {
			r = append(r, carry)
		}
		return r
	}

	isPalindrome := func(n []int) bool {
		for i := 0; i < len(n)/2; i++ {
			if n[i] != n[len(n)-i-1] {
				return false
			}
		}
		return true
	}

	count := 0
	for n := uint64(1); n < uint64(10000); n++ {
		s := numToSlice(n)
		for i := 0; i < 50; i++ {
			s = regroup(flipSum(s))
			if isPalindrome(s) {
				break
			}
		}
		if !isPalindrome(s) {
			count++
		}
	}

	return fmt.Sprintf("%d", count), nil
}
