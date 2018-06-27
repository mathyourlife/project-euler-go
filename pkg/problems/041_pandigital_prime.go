package problems

import (
	"fmt"
)

type PandigitalPrime struct{}

func (p *PandigitalPrime) ID() int {
	return 41
}

func (p *PandigitalPrime) Text() string {
	return `We shall say that an n-digit number is pandigital if it
makes use of all the digits 1 to n exactly once. For example, 2143
is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
`
}

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
func (p *PandigitalPrime) Solve() (string, error) {

	// helper function to create a number from the corresponding
	// digits int slice.
	// ex:
	//   index  = []int{0,1,2,3,4,5,6,8,7}
	//   digits = []int{9,8,7,6,5,4,3,2,1}
	//   returns 987654312
	indexToNumber := func(index, digits []int) uint64 {
		n := uint64(0)
		for _, l := range index {
			n = n*10 + uint64(digits[l])
		}
		return n
	}

	// search for a 7 digit solution
	index := []int{0, 1, 2, 3, 4, 5, 6}
	digits := []int{7, 6, 5, 4, 3, 2, 1}
	for {
		v := indexToNumber(index, digits)
		if IsPrime(v) {
			return fmt.Sprintf("%d", v), nil
		}

		if !LexPerm(index) {
			// Ran out of 7 digit numbers to check
			break
		}
	}

	// If couldn't find a 7 digit solution, search for a
	// 4 digit solution
	index = []int{0, 1, 2, 3}
	digits = []int{4, 3, 2, 1}
	for {
		v := indexToNumber(index, digits)
		if IsPrime(v) {
			return fmt.Sprintf("%d", v), nil
		}

		if !LexPerm(index) {
			// Ran out of 7 digit numbers to check
			break
		}
	}

	return fmt.Sprintf("%d", 0), nil
}
