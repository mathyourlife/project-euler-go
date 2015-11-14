/*
Circular primes

The number, 197, is called a circular prime because all rotations of the
digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100:
2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/

package main

import(
  "fmt"
  "github.com/mathyourlife/lt3maths/prime"
  "math"
)

func place_values(n uint64) []int {
  pv := []int{int(n)}
 	for i := 0; i < len(pv); i++ {
		if pv[i] > 9 {
			if i == len(pv)-1 {
				pv = append(pv, 0)
			}
			regroup := pv[i] / 10
			pv[i+1] += regroup
			pv[i] = pv[i] % 10
		}
	}
	return pv
}

func is_desc(pv []int) bool {
 	for i := 0; i < len(pv) - 1; i++ {
		if pv[i] > pv[i+1] {
			return false
	}}
	return true
}

func main () {
	// Initialize the generator
	is_prime := map[int]bool{}
	p := prime.NewPrimeGenerator()
	// Generate 10 primes

	circular_count := 0

	for {
		n := p.Next()
		if n > 1000000 {
			break
		}
		is_prime[int(n)] = true
		pv := place_values(n)
		// Generate any possible circular primes
		is_circular := true
		circulars := []int{}
		for a := 0; a < len(pv); a++ {
			// Generate the circular numbers and check if they are prime
			circular := 0
			for b := 0; b < len(pv); b++ {
				idx := (a+b) % (len(pv))
				circular += pv[idx] * int(math.Pow(10, float64(b)))
			}
			if ! is_prime[circular] {
				is_circular = false
				break
			}
			circulars = append(circulars, circular)
		}
		if is_circular {
			if (len(circulars) > 1) && (circulars[0] == circulars[1]) {
				circular_count += 1
			} else {
				circular_count += len(circulars)
			}
			fmt.Println(circulars)
		}
	}
	fmt.Println(circular_count)

}
