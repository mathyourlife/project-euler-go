package problems

import (
	"fmt"
)

type PermutedMultiples struct{}

func (p *PermutedMultiples) ID() int {
	return 52
}

func (p *PermutedMultiples) Text() string {
	return `It can be seen that the number, 125874, and its
double, 251748, contain exactly the same digits, but in a
different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x,
5x, and 6x, contain the same digits.
`
}

func (p *PermutedMultiples) Solve() (string, error) {

	tallyDigits := func(num int) map[int]int {
		tally := map[int]int{}
		for {
			if num == 0 {
				break
			}
			tally[num%10]++
			num /= 10
		}
		return tally
	}

	compareMaps := func(m1, m2 map[int]int) bool {
		if len(m1) != len(m2) {
			return false
		}
		for k, v := range m1 {
			if m2[k] != v {
				return false
			}
		}
		return true
	}

	num := 0
	for {
		num++
		tally := tallyDigits(num)

		found := true
		for multiple := 2; multiple <= 6; multiple++ {
			if !compareMaps(tally, tallyDigits(num*multiple)) {
				found = false
				break
			}
		}
		if found {
			return fmt.Sprintf("%d", num), nil
		}
	}

	return fmt.Sprintf("%d", 0), nil
}
