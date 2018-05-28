package problems

import (
	"fmt"
)

type LargestPrimeFactor struct{}

func (p *LargestPrimeFactor) ID() int {
	return 3
}

func (p *LargestPrimeFactor) Text() string {
	return `The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
`
}

func (p *LargestPrimeFactor) Solve() (string, error) {
	var t bool
	val := uint64(600851475143)
	factor := uint64(2)

	for {
		t, val = removeFactor(val, factor)
		if !t {
			factor++
		}
		if val <= 1 {
			break
		}
	}
	return fmt.Sprintf("%d", factor), nil
}

// Given a composite number, determine if factor is a factor.
// If so return true and the quotient of composite/factor
// If not return false and the original composite
func (p *LargestPrimeFactor) isFactor(composite uint64, factor uint64) (bool, uint64) {
	if composite < 2 {
		return false, composite
	}
	attempt := composite / factor
	if attempt*factor == composite {
		return true, attempt
	}
	return false, composite
}
