package problems

import (
	"fmt"
	"math"
)

type ChampernowneConstant struct{}

func (p *ChampernowneConstant) ID() int {
	return 40
}

func (p *ChampernowneConstant) Text() string {
	return `An irrational decimal fraction is created by
concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value
of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
`
}

func (p *ChampernowneConstant) Solve() (string, error) {
	total := 1
	d := 1
	for {
		if d == 10000000 {
			break
		}
		total *= p.character(d)
		d *= 10
	}

	return fmt.Sprintf("%d", total), nil
}

// # digits, qty,            tally
//        1,   9,       1*  9=   9
//        2,  90,   9 + 2* 90= 189
//        3, 900, 189 + 3*900=2889
func (p *ChampernowneConstant) character(pos int) int {
	size := 9
	length := 1
	for {
		if pos > size*length {
			// fmt.Printf("Filled up the %d %d-digits\n", size, numDigits(uint64(size)))
			pos -= size * length
		} else {
			// fmt.Printf("%d left in the %d-digits\n", pos, numDigits(uint64(size)))
			break
		}
		size *= 10
		length++
	}
	// fmt.Printf("need to find the %dth character in the %d digit list\n", pos, length)

	extra := (pos - 1) / length
	target := int(math.Pow10(length-1)) + extra
	idx := ((pos - 1) % length) + 1
	// fmt.Printf("pick the %dth character out of the number %d\n", idx, target)

	// fmt.Printf("shift off %d digits\n", length-idx)
	for i := 0; i < length-idx; i++ {
		target /= 10
	}
	// fmt.Printf("character: %d\n", target%10)
	return target % 10

}
