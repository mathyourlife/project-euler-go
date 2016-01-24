/*
Prime digit replacements

By replacing the 1st digit of the 2-digit number *3, it turns out that six
of the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.

By replacing the 3rd and 4th digits of 56**3 with the same digit, this
5-digit number is the first example having seven primes among the ten
generated numbers, yielding the family: 56003, 56113, 56333, 56443, 56663,
56773, and 56993. Consequently 56003, being the first member of this family,
is the smallest prime with this property.

Find the smallest prime which, by replacing part of the number (not
necessarily adjacent digits) with the same digit, is part of an
eight prime value family.
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/prime"
	"log"
	"math"
	"os"
)

// NewMaskedNum create a MaskedNum instance from a uint64
func NewMaskedNum(num uint64) *MaskedNum {
	p := &MaskedNum{
		Value: num,
	}

	place := uint64(10)
	length := int(math.Log10(float64(num))) + 1
	digits := make([]int, length)

	for i := 0; i < length; i++ {
		digit := num - num/place*place
		digits[length-1-i] = int(digit * uint64(10) / place)
		num -= digit
		place *= 10
	}

	p.Digits = digits
	p.Length = length
	return p
}

// MaskedNum - allows for bit masking of a base 10 positive integer
type MaskedNum struct {
	Value  uint64
	Digits []int
	Length int
}

// Mask - mask out the digits of the number according to a bit mask
// ex: 59212 with mask 13 (b1101)
//   returns 5**1* and the masked digits 922
func (d *MaskedNum) Mask(n int) (string, string) {

	mask := ""
	masked := ""

	for i := len(d.Digits) - 1; i >= 0; i-- {
		if n&1 == 1 {
			mask = fmt.Sprintf("%d", d.Digits[i]) + mask
		} else {
			mask = "*" + mask
			masked = fmt.Sprintf("%d", d.Digits[i]) + masked
		}
		n = n >> 1
	}
	return mask, masked
}

// GeneratePrimes - Generate arrays of primes in groups of all 2 digits,
// then all 3 digits...
func GeneratePrimes(chPrimeGroups chan []*MaskedNum) {

	// Create the prime number generator
	p := prime.NewPrimeGenerator()

	limit := uint64(100)

	var prime uint64

	// Skip single digit primes
	for {
		prime = p.Next()
		if prime >= 10 {
			break
		}
	}

	primeGroup := []*MaskedNum{}
	for {
		if prime > limit {
			chPrimeGroups <- primeGroup
			primeGroup = []*MaskedNum{}
			limit *= 10
		}
		primeGroup = append(primeGroup, NewMaskedNum(prime))
		prime = p.Next()
	}
}

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("Prime digit replacements")

	chPrimeGroups := make(chan []*MaskedNum, 4)

	// Separate goroutine to generate sets of n-digit prime numbers
	// starting with all the 2-digit primes
	go GeneratePrimes(chPrimeGroups)

	// Quick map lookup if the masked digits are 11 or 555 or ...
	sames := map[string]bool{}

SEARCH_LOOP:
	for {
		primeGroup := <-chPrimeGroups

		// Make sure the sames map has all entries needed.  For 4 digit primes,
		// we should have it populated up to "999" = true
		for _, n := range primeGroup {
			for i := 0; i < 10; i++ {
				key := ""
				for j := 0; j < n.Length-1; j++ {
					key += fmt.Sprintf("%d", i)
				}
				sames[key] = true
			}
			break
		}

		// Track the number of occurrences for a prime value family's
		// mask like 56**3
		familyMask := map[string]int{}
		// Generate prime value families like 56**3 = [56003, 56113, 56333,
		// 56443, 56663, 56773, 56993]
		familyMembers := map[string][]uint64{}

		for _, n := range primeGroup {
			// For each prime in the current group such as 56003, cycle throug all
			// bit masks from 1 to 30 (b11110). (Don't need to mask the entire number)
			for m := 1; float64(m) <= math.Pow(2, float64(n.Length))-2; m++ {
				mask, masked := n.Mask(m)
				if sames[masked] {
					familyMask[mask]++
					familyMembers[mask] = append(familyMembers[mask], n.Value)
				}
			}
		}

		// See if we found the eight prime value family
		for mask, count := range familyMask {
			if count == 8 {
				minVal := uint64(0)
				for _, val := range familyMembers[mask] {
					if minVal == 0 || val < minVal {
						minVal = val
					}
				}
				log.Println("Prime Value Family:", mask)
				log.Println("Family Members:", familyMembers[mask])
				log.Println("Youngest Member:", minVal)
				break SEARCH_LOOP
			}
		}
	}
}
