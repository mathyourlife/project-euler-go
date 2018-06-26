package problems

import (
	"fmt"
)

type DigitFactorials struct{}

func (p *DigitFactorials) ID() int {
	return 34
}

func (p *DigitFactorials) Text() string {
	return `145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of
their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
`
}

// The search has an upper bound of a 7 digit number.  The smallest
// seven digit number 1,000,000 is possbile to be written as the
// sum of the factorials of the 7 digits. Or
//   10^(7-1) < 7(9!)
// At 8 digits, the smallest number can not be reach by the
// sum of the factorial of the digits
//   10^(8-1) > 8(9!)
func (p *DigitFactorials) Solve() (string, error) {
	digitFactorials := map[uint64]uint64{
		0: 1,
	}
	prod := uint64(1)
	for i := uint64(1); i < uint64(10); i++ {
		prod *= i
		digitFactorials[i] = prod
	}

	solution := uint64(0)
	// Search all numbers with less than 8 digits
	for n := uint64(3); n < uint64(10000000); n++ {
		value := n
		digitSum := uint64(0)
		for {
			if value == 0 {
				break
			}
			digitSum += digitFactorials[value%10]
			if digitSum > n {
				break
			}
			value /= 10
		}
		if digitSum == n {
			solution += n
		}
	}

	return fmt.Sprintf("%d", solution), nil
}
