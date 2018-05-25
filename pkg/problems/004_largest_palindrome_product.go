package problems

import (
	"fmt"
)

type LargestPalindromeProduct struct{}

func (p *LargestPalindromeProduct) ID() int {
	return 4
}

func (p *LargestPalindromeProduct) Text() string {
	return `A palindromic number reads the same both ways. The largest palindrome
made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
`
}

func (p *LargestPalindromeProduct) Solve() (string, error) {
	max := 999
	min := 100
	var a, b int

	for sum := max * 2; sum >= 2*min; sum-- {
		if sum > max+min {
			a, b = sum-max, max
		} else {
			a, b = min, sum-min
		}
		for {
			t, c, d := p.check(a, b)
			if t {
				return fmt.Sprintf("%d", c*d), nil
			}
			a++
			b--
			if a > max || b < min {
				break
			}
		}
	}
	return fmt.Sprintf("%d", 0), nil
}

func (p *LargestPalindromeProduct) isPalindrome(n string) bool {
	for i := 0; i < (len(n)+1)/2; i++ {
		if n[i:i+1] != n[len(n)-i-1:len(n)-i] {
			return false
		}
	}
	return true
}

func (p *LargestPalindromeProduct) check(a int, b int) (bool, int, int) {
	product := fmt.Sprintf("%d", a*b)
	if p.isPalindrome(product) {
		return true, a, b
	}
	return false, 0, 0
}
