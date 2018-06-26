package problems

import (
	"fmt"
)

type DoubleBasePalindromes struct{}

func (p *DoubleBasePalindromes) ID() int {
	return 36
}

func (p *DoubleBasePalindromes) Text() string {
	return `The decimal number, 585 = 1001001001_2 (binary), is
palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in
base 10 and base 2.

(Please note that the palindromic number, in either base, may not include
leading zeros.)
`
}

// To avoid working with string manipulation, a palindrome can be
// compared by value if you pop off a modulus and shift it onto
// the reverse value.  Such as the reverse of 684 in base 10 is
//
// 684 mod 10 = 4  =>   0*10 + 4 = 4
//  68 mod 10 = 8  =>   4*10 + 8 = 48
//   6 mod 10 = 6  =>  48*10 + 6 = 486
//
// same applies for base 2.  The reversal of 13 (1101 in base 2) is
// 11 (1011 in base2)
//
// 13 mod 2 = 1  =>  0*2 + 1 = 1
//  6 mod 2 = 0  =>  1*2 + 0 = 2
//  3 mod 2 = 1  =>  2*2 + 1 = 5
//  1 mod 2 = 1  =>  5*2 + 1 = 11
func (p *DoubleBasePalindromes) Solve() (string, error) {

	reverse := func(n, base uint64) uint64 {
		r := uint64(0)
		for {
			if n == 0 {
				break
			}
			r = r*base + n%base
			n /= base
		}
		return r
	}

	sum := uint64(0)
	for n := uint64(1); n < uint64(1000000); n++ {
		if reverse(n, 10) == n && reverse(n, 2) == n {
			sum += n
		}
	}

	return fmt.Sprintf("%d", sum), nil
}
