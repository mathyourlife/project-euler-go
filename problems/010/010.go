/*
The sum of the primes below 10 is
2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/prime"
)

func main() {

	limit := uint64(2000000)
	p := prime.NewPrimeGenerator()
	sum := uint64(0)

	for {
		n := p.Next()
		if n > limit {
			break
		}
		sum += n
	}
	fmt.Println(sum)
}
