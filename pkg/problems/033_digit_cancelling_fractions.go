package problems

import (
	"fmt"
)

type DigitCancellingFractions struct{}

func (p *DigitCancellingFractions) ID() int {
	return 33
}

func (p *DigitCancellingFractions) Text() string {
	return `The fraction 49/98 is a curious fraction, as an inexperienced
mathematician in attempting to simplify it may incorrectly
believe that 49/98 = 4/8, which is correct, is obtained by
cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of
fraction, less than one in value, and containing two digits
in the numerator and denominator.

If the product of these four fractions is given in its lowest
common terms, find the value of the denominator.
`
}

// The dropped digit is shared between numerator and denominator.
// Loop through adding a second digit and check for equality.
// ex:
//   If n=1,d=4,drop=6 the checks are 16/46, 16/64, 61/46, 61/64
func (p *DigitCancellingFractions) Solve() (string, error) {

	checkFraction := func(n, d, num, den uint64) bool {
		if num < 10 || den < 10 {
			return false
		}
		if num >= den {
			return false
		}
		if n*10 == num {
			return false
		}
		gcf1 := gcf([]uint64{num, den})
		gcf2 := gcf([]uint64{n, d})
		if num/gcf1 == n/gcf2 && den/gcf1 == d/gcf2 {
			return true
		}
		return false
	}

	solutionNum := uint64(1)
	solutionDen := uint64(1)
	for n := uint64(0); n < uint64(10); n++ {
		for d := uint64(1); d < uint64(10); d++ {
			if n >= d {
				continue
			}
			for drop := uint64(0); drop < uint64(10); drop++ {
				// If n=1,d=4,drop=6 the checks
				// are 16/46, 16/64, 61/46, 61/64
				checks := []struct {
					n   uint64
					d   uint64
					num uint64
					den uint64
				}{
					{n, d, n*10 + drop, d*10 + drop},
					{n, d, n*10 + drop, drop*10 + d},
					{n, d, drop*10 + n, d*10 + drop},
					{n, d, drop*10 + n, drop*10 + d},
				}

				for _, check := range checks {
					if checkFraction(check.n, check.d, check.num, check.den) {
						// Found!
						solutionNum *= check.num
						solutionDen *= check.den
					}
				}
			}
		}
	}
	return fmt.Sprintf("%d", solutionDen/gcf([]uint64{solutionNum, solutionDen})), nil
}
