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

func is_pandigital(prime uint64) bool {
	var last_digit, bit uint64
	count := uint64(0)
	tmp := uint64(0)
	digits := uint64(0)
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

func go_primes(ch chan uint64) {

	p := prime.NewPrimeGenerator()

	for {
		n := p.Next()
		ch <- n
	}

}

func go_is_pandigital(ch_primes chan uint64, ch_ans chan uint64) {

	var n uint64
	var pandigital uint64

	for {
		n = <- ch_primes
		// Due to divisibility by 3 rule (sum of digits x of 3), pandigital
		// primes can't be:
		// 12=3         2 digits
		// 123=6        3 digits
		// 12345=15     5 digits
		// 124356=21    6 digits
		// 12435678=36  8 digits
		// 124356789=45 9 digits
		//
		// Leaving 4 and 7 digit primes
		if n <= 7654321 && n >= 1234567 {
			if is_pandigital(n) {
				pandigital = n
			}
		} else if n <= 4321 && n >= 1234 {
			if is_pandigital(n) {
				pandigital = n
			}
		} else if n > 7654321 {
			break
		}

	}
	ch_ans <- pandigital
}

func main() {
	fmt.Println("Pandigital prime")

	ch_primes := make(chan uint64)
	ch_ans := make(chan uint64)

	go go_primes(ch_primes)
	go go_is_pandigital(ch_primes, ch_ans)

	n := <- ch_ans
	fmt.Println(n)
}
