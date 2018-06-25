package problems

import (
	"fmt"
	"math"
)

type DigitFifthPowers struct {
	powers map[int]int
}

func (p *DigitFifthPowers) ID() int {
	return 30
}

func (p *DigitFifthPowers) Text() string {
	return `Surprisingly there are only three numbers that can be written
as the sum of fourth powers of their digits:

    1634 = 1^4 + 6^4 + 3^4 + 4^4
    8208 = 8^4 + 2^4 + 0^4 + 8^4
    9474 = 9^4 + 4^4 + 7^4 + 4^4

As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum
of fifth powers of their digits
`
}

func (p *DigitFifthPowers) Solve() (string, error) {

	exp := 5

	// Cache the Nth powers of the digits
	p.powers = map[int]int{}
	for digit := 0; digit < 10; digit++ {
		p.powers[digit] = int(math.Pow(float64(digit), float64(exp)))
	}

	// Calculate an upper limit for the number of digits
	// e.g. For the exponent exp=4, at 99,999 the sum of the digits^4 is
	// eclipsed by the number the digits represent.
	//       9 => 9^4 = 6561
	//      99 => 9^4 + 9^4 = 13122
	//     999 => 9^4 + 9^4 + 9^4 = 19683
	//   9,999 => 9^4 + 9^4 + 9^4 + 9^4 = 26244
	//  99,999 => 9^4 + 9^4 + 9^4 + 9^4 + 9^4 = 32805
	// 999,999 => 9^4 + 9^4 + 9^4 + 9^4 + 9^4 + 9^4 = 39366

	// Simple function to return how many digits in an integer
	numDigits := func(n int) int {
		return int(math.Log10(float64(n)) + 1)
	}

	maxLen := 1
	for {
		if numDigits(p.powers[9]*maxLen) < maxLen {
			break
		}
		maxLen++
	}
	maxLen--

	// given a number, break it into it's digits, find the sum
	// of the digits^exp, and determine if it equals the original
	// number.
	checkNumber := func(n int) bool {
		mod := n
		sum := 0
		for {
			sum += p.powers[mod%10]
			// fmt.Println(mod%10, p.powers[mod%10], sum)
			if sum > n {
				// fmt.Println("sum exceeds number, discarding")
				return false
			}
			mod = mod / 10
			if mod <= 0 {
				if sum == n {
					return true
				}
				break
			}
		}
		return false
	}

	sum := 0
	for n := 10; n < int(math.Pow10(maxLen)); n++ {
		if checkNumber(n) {
			sum += n
		}
	}
	return fmt.Sprintf("%d", sum), nil
}
