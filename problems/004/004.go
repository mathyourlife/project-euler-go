/*
A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

6 : 1+6 | 2+6 | 3+6 | 4+6 | 5+6 | 6+6
5 : 1+5 | 2+5 | 3+5 | 4+5 | 5+5 | 6+5
4 : 1+4 | 2+4 | 3+4 | 4+4 | 5+4 | 6+4
3 : 1+3 | 2+3 | 3+3 | 4+3 | 5+3 | 6+3
2 : 1+2 | 2+2 | 3+2 | 4+2 | 5+2 | 6+2
1 : 1+1 | 2+1 | 3+1 | 4+1 | 5+1 | 6+1
     1     2     3     4     5     6


*/
package main

import (
	"fmt"
)

func is_palindrome(n string) bool {
	for i := 0; i < (len(n)+1)/2; i++ {
		if n[i:i+1] != n[len(n)-i-1:len(n)-i] {
			return false
		}
	}
	return true
}

func check(a int, b int) (bool, int, int) {
	product := fmt.Sprintf("%d", a*b)
	if is_palindrome(product) {
		return true, a, b
	}
	return false, 0, 0
}

func main() {
	max := 999
	min := 100
	var a, b int

Loop:
	for sum := max * 2; sum >= 2*min; sum-- {
		if sum > max+min {
			a, b = sum-max, max
		} else {
			a, b = min, sum-min
		}
		for {
			t, c, d := check(a, b)
			if t {
				fmt.Printf("%d*%d=%d\n", c, d, c*d)
				break Loop
			}
			a++
			b--
			if a > max || b < min {
				break
			}
		}
	}
}
