/*
Pandigital prime

We shall say that an n-digit number is pandigital if it makes use of all the
digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital
and is also prime.

What is the largest n-digit pandigital prime that exists?
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/prime"
)

func is_pandigital(prime uint) bool {
	var last_digit, bit uint
	count := uint(0)
	tmp := uint(0)
	digits := uint(0)
	for {
		if prime <= 0 {
			break
		}
		last_digit = prime - ((prime / 10) * 10)
		bit = 1 << (last_digit - 1)
		digits |= bit
		count++
		if tmp == digits {
			break
		}
		tmp = digits
		prime /= 10
	}
	return digits == (1<<count)-1
}

func main() {
	fmt.Println("Pandigital prime")

	p := prime.NewPrimeGenerator()

	for {
		n := p.Next()
		if n > 1000000000 {
			// if n > 7652413 {
			break
		}
		if is_pandigital(uint(n)) {
			fmt.Printf("%d\n", n)
		}
	}
}
