/*
Truncatable primes

The number 3797 has an interesting property. Being prime itself, it is
possible to continuously remove digits from left to right, and remain
prime at each stage: 3797, 797, 97, and 7. Similarly we can work from
right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from
left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/prime"
	"math"
)

func drop_right(n uint64, is_prime map[uint64]bool) bool {
	for {
		if n < 10 {
			break
		}
		n /= 10
		if !is_prime[n] {
			return false
		}
	}
	return true
}

func drop_left(n uint64, is_prime map[uint64]bool) bool {

	mod := uint64(math.Pow(10, float64(int(math.Log10(float64(n))))))

	for {
		if !is_prime[n%mod] {
			return false
		}
		mod /= 10
		if mod <= 1 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println("Truncatable Primes")
	is_prime := map[uint64]bool{}
	p := prime.NewPrimeGenerator()

	tp_count := 0
	tp_sum := uint64(0)
	for {
		n := p.Next()
		is_prime[n] = true

		if drop_right(n, is_prime) && drop_left(n, is_prime) {
			fmt.Println(n)
			tp_count += 1
			tp_sum += n
		}

		if tp_count >= 11 {
			break
		}
	}
	fmt.Println(tp_sum)
}
