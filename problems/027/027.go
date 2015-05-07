/*
Quadratic Primes

Euler discovered the remarkable quadratic formula:

n² + n + 41

It turns out that the formula will produce 40 primes for the
consecutive values n = 0 to 39. However, when n = 40,
40² + 40 + 41 = 40(40 + 1) + 41
is divisible by 41, and certainly when n = 41,
41² + 41 + 41 is clearly divisible by 41.

The incredible formula  n² − 79n + 1601 was discovered, which
produces 80 primes for the consecutive values n = 0 to 79.
The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

    n² + an + b, where |a| < 1000 and |b| < 1000

    where |n| is the modulus/absolute value of n
    e.g. |11| = 11 and |−4| = 4

Find the product of the coefficients, a and b, for the
quadratic expression that produces the maximum number of
primes for consecutive values of n, starting with n = 0.

*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/prime"
)

type IsPrime struct {
	p         *prime.PrimeGenerator
	pm        map[uint64]bool
	max_prime uint64
}

func NewIsPrime() *IsPrime {
	return &IsPrime{
		p:         prime.NewPrimeGenerator(),
		pm:        map[uint64]bool{},
		max_prime: 0,
	}
}

func (ip *IsPrime) IsIt(n uint64) bool {
	if n < ip.max_prime {
		return ip.pm[n]
	}
	for {
		new_prime := ip.p.Next()
		ip.pm[new_prime] = true
		if new_prime > n {
			ip.max_prime = new_prime
			break
		}
	}
	return ip.pm[n]
}

func (ip *IsPrime) Nth(n int) uint64 {
	if n < len(ip.p.Primes)-1 {
		return ip.p.Primes[n]
	}
	for {
		new_prime := ip.p.Next()
		ip.pm[new_prime] = true
		if len(ip.p.Primes) > n {
			ip.max_prime = new_prime
			break
		}
	}
	return ip.p.Primes[n]
}

func check_prime_generation(a int, b int, ip *IsPrime) int {
	n := 0
	for {
		candidate := n*n + a*n + b
		if candidate < 2 {
			break
		}
		if !ip.IsIt(uint64(candidate)) {
			break
		}
		n++
	}
	return n
}

func main() {
	ip := NewIsPrime()

	var args []int
	max_len := 0
	i := 0
	for {
		b := ip.Nth(i)
		if b > 1000 {
			break
		}
		for a := -999; a < 1000; a++ {
			pg_len := check_prime_generation(a, int(b), ip)
			// fmt.Println("a", a, "b", b, "length", pg_len)
			if pg_len > max_len {
				max_len = pg_len
				args = []int{a, int(b), pg_len}
			}
		}
		i++
	}
	fmt.Println("a", args[0], "b", args[1], "a*b", args[0]*args[1], "length", args[2])
}
