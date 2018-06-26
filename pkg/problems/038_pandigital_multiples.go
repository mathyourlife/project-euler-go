package problems

import (
	"fmt"
	"math"
)

type PandigitalMultiples struct{}

func (p *PandigitalMultiples) ID() int {
	return 38
}

func (p *PandigitalMultiples) Text() string {
	return `Take the number 192 and multiply it by each of 1, 2, and 3:

    192 × 1 = 192
    192 × 2 = 384
    192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576.
We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3,
4, and 5, giving the pandigital, 918273645, which is the concatenated
product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed
as the concatenated product of an integer with (1,2, ... , n) where n > 1?
`
}

// not the prettiest code, but... it works
func (p *PandigitalMultiples) Solve() (string, error) {

	// digits is a int slice that stands for the digit in a
	// particular place value.  digits are initially arranged
	// to represent the largest pandigital number 987654321
	// and will decrease as the index slice increases.
	digits := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	index := make([]int, 0, len(digits))
	for i, _ := range digits {
		index = append(index, i)
	}

	// helper function to create a number from the corresponding
	// digits int slice.
	// ex:
	//   index = []int{0,1,2,3,4,5,6,8,7}
	//   returns 987654312
	indexToNumber := func(index []int) uint64 {
		n := uint64(0)
		for _, l := range index {
			n = n*10 + uint64(digits[l])
		}
		return n
	}

	// loop until a solution is found or LexPerm runs out
	// of permutations and the index is sorted in decreasing order
	for {
		// The result can be a concatenation of the product of up
		// to the first 4 digits.  The pandigital 192384576
		// could have a base value of 1, 19, 192, or 1923
		for length := 1; length <= 4; length++ {
			base := uint64(0)
			for _, l := range index[:length] {
				base = base*10 + uint64(digits[l])
			}

			// append multiple of the base until creating at least a 9
			// digit number.  A base of 19 would create
			// 19 38 57 76 95 or 1,938,577,695
			concat := base
			i := uint64(2)
			for {
				// break once we reach at least a 9 digit number
				if concat >= 99999999 {
					break
				}
				next := base * i
				concat = concat*uint64(math.Pow10(numDigits(next))) + next
				i++
			}
			// If the concatenated number doesn't match the original,
			// we've found an answer, just not the largest.
			if concat == indexToNumber(index) {
				return fmt.Sprintf("%d", concat), nil
			}
		}
		if !LexPerm(index) {
			break

		}
	}

	return fmt.Sprintf("%d", 0), nil
}
