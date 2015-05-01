/*
2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?
*/
package main

import (
	"fmt"
)

func is_a_factor(composite uint64, factor uint64) (bool, uint64) {
	if composite < 2 {
		return false, composite
	}
	attempt := composite / factor
	if attempt*factor == composite {
		return true, attempt
	}
	return false, composite
}

func prime_factors(n uint64) map[uint64]int {
	var t bool
	factor := uint64(2)
	pf := map[uint64]int{}

	for {
		t, n = is_a_factor(n, factor)
		if t {
			pf[factor]++
			if n <= 1 {
				break
			}
		} else {
			factor++
		}
		if n <= 1 {
			break
		}
	}
	return pf
}

func main() {

	pfs := map[uint64]int{}

	for n := uint64(2); n <= 20; n++ {
		pf := prime_factors(n)
		for f, c := range pf {
			if pfs[f] < c {
				pfs[f] = c
			}
		}
	}

	prod := uint64(1)
	for f, c := range pfs {
		for i := 0; i < c; i++ {
			prod *= f
		}
	}
	fmt.Println(prod)
}
