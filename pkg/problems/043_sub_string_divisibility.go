package problems

import (
	"fmt"
)

type SubStringDivisibility struct{}

func (p *SubStringDivisibility) ID() int {
	return 43
}

func (p *SubStringDivisibility) Text() string {
	return `The number, 1406357289, is a 0 to 9 pandigital number
because it is made up of each of the digits 0 to 9 in some order,
but it also has a rather interesting sub-string divisibility property.

Let d(1) be the 1st digit, d(2) be the 2nd digit, and so on. In this way,
we note the following:

    d(2)d(3)d(4)=406 is divisible by 2
    d(3)d(4)d(5)=063 is divisible by 3
    d(4)d(5)d(6)=635 is divisible by 5
    d(5)d(6)d(7)=357 is divisible by 7
    d(6)d(7)d(8)=572 is divisible by 11
    d(7)d(8)d(9)=728 is divisible by 13
    d(8)d(9)d(10)=289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.
`
}

func (p *SubStringDivisibility) Solve() (string, error) {

	// helper function to create a number from the corresponding
	// digits int slice.
	// ex:
	//   digits = []int{3,2,4}
	//   returns 324
	sliceToNumber := func(digits []int) uint64 {
		n := uint64(0)
		for _, d := range digits {
			n = n*10 + uint64(d)
		}
		return n
	}

	check := func(digits []int) bool {
		if digits[3]%2 != 0 {
			return false
		}
		if (digits[2]+digits[3]+digits[4])%3 != 0 {
			return false
		}
		if digits[5] != 0 && digits[5] != 5 {
			return false
		}
		if sliceToNumber(digits[4:7])%7 != 0 {
			return false
		}
		if sliceToNumber(digits[5:8])%11 != 0 {
			return false
		}
		if sliceToNumber(digits[6:9])%13 != 0 {
			return false
		}
		if sliceToNumber(digits[7:10])%17 != 0 {
			return false
		}

		return true
	}

	sum := uint64(0)
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for {
		if check(digits) {
			sum += sliceToNumber(digits)
		}
		if !LexPerm(digits) {
			break
		}
	}
	return fmt.Sprintf("%d", sum), nil
}
