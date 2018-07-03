package problems

import (
	"fmt"
	"math"
)

type PrimeDigitReplacements struct{}

func (p *PrimeDigitReplacements) ID() int {
	return 51
}

func (p *PrimeDigitReplacements) Text() string {
	return `By replacing the 1st digit of the 2-digit number *3, it
turns out that six of the nine possible values: 13, 23, 43, 53, 73,
and 83, are all prime.

By replacing the 3rd and 4th digits of 56**3 with the same digit, this
5-digit number is the first example having seven primes among the ten
generated numbers, yielding the family: 56003, 56113, 56333, 56443, 56663,
56773, and 56993. Consequently 56003, being the first member of this family,
is the smallest prime with this property.

Find the smallest prime which, by replacing part of the number (not
necessarily adjacent digits) with the same digit, is part of an
eight prime value family.
`
}

func (p *PrimeDigitReplacements) Solve() (string, error) {

	n := 2
	for {
		min, familySize := p.search(n)
		if familySize == 8 {
			return fmt.Sprintf("%d", min), nil
		}
		n++
	}

	return "", nil
}

// search the n-digit space for the largest prime family.
// If multiple prime families exist of the same size, return
// the one containing the smallest value.
func (p *PrimeDigitReplacements) search(n int) (int, int) {
	// helper function to create a number from the corresponding
	// digits int slice.
	// ex:
	//   digits = []int{3,2,4}
	//   returns 324
	sliceToNumber := func(digits []int) int {
		n := 0
		for _, d := range digits {
			n = n*10 + d
		}
		return n
	}

	lowerLimit := int(math.Pow10(n - 1))
	upperLimit := lowerLimit * 10

	bitsOnCount := func(mask uint64) int {
		n := 0
		for {
			if mask == 0 {
				return n
			}
			if mask&1 == 1 {
				n++
			}
			mask = mask >> 1
		}
		return 0
	}

	numbers := func(n int) chan []int {
		lowerLimit := int(math.Pow10(n - 1))
		upperLimit := lowerLimit * 10

		ch := make(chan []int)
		if lowerLimit == upperLimit {
			close(ch)
			return ch
		}

		numToSlice := func(n int) []int {
			v := make([]int, 0)
			for {
				if n == 0 {
					break
				}
				v = append(v, n%10)
				n /= 10
			}
			return v
		}

		go func() {
			for i := lowerLimit; i < upperLimit; i++ {
				ch <- numToSlice(i)
			}
			close(ch)
		}()

		return ch
	}

	generateFamilies := func(mask uint64, digits int) chan []int {
		ch := make(chan []int)

		go func() {
			bitsOn := bitsOnCount(mask)
			for seq := range numbers(digits - bitsOn) {
				family := make([]int, 0)
				for a := 0; a < 10; a++ {
					s := make([]int, len(seq))
					for i := 0; i < len(seq); i++ {
						s[i] = seq[i]
					}
					d := make([]int, 0, digits)
					m := mask
					for i := 0; i < digits; i++ {
						if m&1 == 1 {
							d = append(d, a)
						} else {
							d = append(d, s[0])
							s = s[1:]
						}
						m = m >> 1
					}
					member := sliceToNumber(d)
					if member < lowerLimit || member >= upperLimit {
						continue
					}
					if !IsPrime(uint64(member)) {
						continue
					}
					family = append(family, member)
				}
				if len(family) == 0 {
					continue
				}
				ch <- family
			}
			close(ch)
		}()
		return ch
	}

	smallestPrime := 0
	familySize := 0
	mask := uint64(1)<<uint64(n) - 1
	for {
		mask--
		families := generateFamilies(mask, n)
		for family := range families {
			if len(family) > familySize {
				smallestPrime = family[0]
				familySize = len(family)
			} else if len(family) > familySize && family[0] < smallestPrime {
				smallestPrime = family[0]
				familySize = len(family)
			}
		}
		if mask == 1 {
			break
		}
	}
	return smallestPrime, familySize
}
