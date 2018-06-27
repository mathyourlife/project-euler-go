package problems

import (
	"fmt"
	"math"
)

type Sign int

const (
	SignNone = iota
	SignPositive
	SignNegative
)

// Rational represent a number in rational form
type Rational struct {
	Sign        Sign
	Numerator   uint64
	Denominator uint64
}

// Decimal conver a Rational number to a decimal format
func (r *Rational) Decimal() *Decimal {
	divisor := r.Denominator
	whole := r.Numerator / divisor // TODO: this needs to be updated for non-base 10
	base := uint64(10)

	remainder := r.Numerator - (whole * divisor)

	decimals := []uint64{}
	remainders := []uint64{remainder}
	i := 0
	for {
		// Long division steps
		remainder *= base
		digit := remainder / divisor
		decimals = append(decimals, digit)
		remainder = remainder - (digit * divisor)

		// Check for repeating decimals if the same remainder
		// has been seen before
		for i, r := range remainders {
			if r == remainder {
				// Repeating decimal found.  Return Decimal
				// with the whole number, non-repeating and
				// repeating decimal portions
				return &Decimal{
					whole:     whole,
					decimals:  decimals[:i],
					repetends: decimals[i:],
				}
			}
		}

		// No remainder is a terminating decimal
		if remainder == 0 {
			break
		}
		remainders = append(remainders, remainder)
		i++
	}
	// Terminating decimal
	return &Decimal{
		whole:    whole,
		decimals: decimals,
	}
}

// Decimal represent a number in decimal format with whole
// number, non-repeating decimal, and repeating decimal portions
type Decimal struct {
	whole     uint64
	decimals  []uint64
	repetends []uint64
}

// String convert the decimal number to a string format.  Repeating
// decmials use a parentheses format.  ex: 1/12 = "0.08(3)"
func (d Decimal) String() string {
	decimalStr := ""
	for _, decimal := range d.decimals {
		decimalStr += fmt.Sprintf("%d", decimal)
	}
	if d.repetends == nil || len(d.repetends) == 0 {
		return fmt.Sprintf("%d.%s", d.whole, decimalStr)
	}
	repetendStr := ""
	for _, repetend := range d.repetends {
		repetendStr += fmt.Sprintf("%d", repetend)
	}

	return fmt.Sprintf("%d.%s(%s)", d.whole, decimalStr, repetendStr)
}

func numDigits(n uint64) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(n)) + 1)
}

// IsPandigital determine if a number is pandigital up to
// the number of digits.  A four digit number should
// contain the digits 1,2,3,and 4 only once
// and in any order.
func IsPandigital(n uint64) bool {
	var bit uint64
	count := uint64(0)
	tmp := uint64(0)
	digits := uint64(0)
	for {
		if n == 0 {
			break
		}
		bit = 1 << (n%10 - 1)
		digits |= bit
		count++
		// If digits hasn't changed since last time,
		// we've got a repeat digit, so not pandigital
		if tmp == digits {
			return false
		}
		tmp = digits
		n /= 10
	}
	return digits == (1<<count)-1
}
