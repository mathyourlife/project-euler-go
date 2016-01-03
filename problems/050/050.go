/*
Consecutive prime sum

The prime 41, can be written as the sum of six consecutive primes:
41 = 2 + 3 + 5 + 7 + 11 + 13

This is the longest sum of consecutive primes that adds to a
prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a
prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the
most consecutive primes?
*/

package main

import (
	"github.com/mathyourlife/lt3maths/prime"
	"log"
	"os"
)

func search(is_prime map[uint64]bool, p *prime.PrimeGenerator) (int, uint64, uint64, uint64) {
	max_idx := len(is_prime) - 1
	max_length := 0

	var low, high, solution uint64

	for start_idx := 0; start_idx <= max_idx; start_idx++ {
		sum := uint64(0)
		for i := start_idx; i < max_idx; i++ {
			sum += p.Primes[i]
			if is_prime[sum] && i-start_idx+1 > max_length {
				max_length = i - start_idx + 1
				low = p.Primes[start_idx]
				high = p.Primes[i]
				solution = sum
			}
			if sum > 1000000 {
				break
			}
		}
	}
	return max_length, low, high, solution
}

func main() {

	log.SetOutput(os.Stdout)
	log.Println("Consecutive prime sum")

	// Create the prime number generator
	p := prime.NewPrimeGenerator()

	is_prime := map[uint64]bool{}
	for {
		prime := p.Next()
		if prime > 1000000 {
			break
		}
		is_prime[prime] = true
	}

	max_length, low, high, solution := search(is_prime, p)

	log.Printf("The sum of %d consecutive primes from %d to %d is %d",
		max_length, low, high, solution)

}
