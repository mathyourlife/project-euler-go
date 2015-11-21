/*
Sub-string divisibility

The number, 1406357289, is a 0 to 9 pandigital number because it is made up
of each of the digits 0 to 9 in some order, but it also has a rather
interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way,
we note the following:

    d2d3d4=406 is divisible by 2
    d3d4d5=063 is divisible by 3
    d4d5d6=635 is divisible by 5
    d5d6d7=357 is divisible by 7
    d6d7d8=572 is divisible by 11
    d7d8d9=728 is divisible by 13
    d8d9d10=289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/lexperm"
)

func check_divisibility(sub_n int, divisor int) bool {
	// fmt.Println("sub_n", sub_n, "divisor", divisor)
	if sub_n == sub_n/divisor*divisor {
		return true
	}
	return false
}

func main() {
	fmt.Println("Sub-string divisibility")

	ans := 0

	lp := lexperm.LexPerm{}
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for {
		n := 0
		for _, d := range digits {
			n = (n * 10) + d
		}

		pass := true
		if !check_divisibility(n-(n/1000*1000), 17) {
			pass = false
		} else if !check_divisibility(n/10-(n/10000*1000), 13) {
			pass = false
		} else if !check_divisibility(n/100-(n/100000*1000), 11) {
			pass = false
		} else if !check_divisibility(n/1000-(n/1000000*1000), 7) {
			pass = false
		} else if !check_divisibility(n/10000-(n/10000000*1000), 5) {
			pass = false
		} else if !check_divisibility(n/100000-(n/100000000*1000), 3) {
			pass = false
		} else if !check_divisibility(n/1000000-(n/1000000000*1000), 2) {
			pass = false
		}
		if pass {
			fmt.Println(n)
			ans += n
		}
		more := lp.Next(digits)
		if !more {
			break
		}
	}
	fmt.Println("ans", ans)

}
