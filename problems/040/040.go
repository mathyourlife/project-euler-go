/*
Champernowne's constant

An irrational decimal fraction is created by concatenating the
positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value
of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
*/

package main

import (
	"fmt"
	"math"
)

func champernowne_sub(n int, l int) int {
	base := int(math.Pow(10, float64(l-1))) - 1
	for {
		if n <= l {
			break
		}
		n -= l
		base++
	}

	base++

	return int(math.Mod(float64(base/int(math.Pow(float64(10), float64(int(math.Log10(float64(base)))+1-n)))), 10))
}

func champernowne_digit(n int) int {

	l := float64(1)
	for {
		shift := 9 * int(math.Pow(10, l-1)) * int(l)
		if shift >= n {
			break
		} else {
			n -= shift
			l++
		}
	}

	return champernowne_sub(n, int(l))
}

func main() {
	fmt.Println("Champernowne's constant")

	// d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000

	fmt.Println(
		champernowne_digit(1) *
			champernowne_digit(10) *
			champernowne_digit(100) *
			champernowne_digit(1000) *
			champernowne_digit(10000) *
			champernowne_digit(100000) *
			champernowne_digit(1000000))
}
