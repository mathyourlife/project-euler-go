/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143?
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
	if attempt * factor == composite {
		return true, attempt
	}
	return false, composite
}

func main() {


	var t bool
	val := uint64(600851475143)
	factor := uint64(2)

	fmt.Printf("Factors of %d:\n", val)
	for{
		t, val = is_a_factor(val, factor)
		if t {
			if val <= 1 {
				fmt.Printf("%d", factor)
				break
			} else {
				fmt.Printf("%d * ", factor)
			}
		} else {
			factor++
		}
		if val <= 1 {
			break
		}
	}
	fmt.Printf("\n")

}