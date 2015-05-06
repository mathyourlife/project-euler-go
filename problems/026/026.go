/*
Reciprocal cycles

A unit fraction contains 1 in the numerator. The decimal
representation of the unit fractions with denominators
2 to 10 are given:

	1/2	= 	0.5
	1/3	= 	0.(3)
	1/4	= 	0.25
	1/5	= 	0.2
	1/6	= 	0.1(6)
	1/7	= 	0.(142857)
	1/8	= 	0.125
	1/9	= 	0.(1)
	1/10	= 	0.1

Where 0.1(6) means 0.166666..., and has a 1-digit recurring
cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest
recurring cycle in its decimal fraction part.
*/

package main

import (
	"fmt"
)

func decimal_nutation(num int, den int) (string, int) {
	digits := []int{0}

	terminating := false
	remainders_map := map[int]bool{num: true}
	remainders_list := []int{num}
	for {
		digits = append(digits, num*10/den)
		num = (num * 10) % den
		remainders_list = append(remainders_list, num)
		if num == 0 {
			terminating = true
			break
		}
		if remainders_map[num] {
			break
		}
		remainders_map[num] = true
	}

	s := ""
	var rlen int
	pos := 0
	repeating := false
	for i := 0; i < len(remainders_list); i++ {
		s += fmt.Sprintf("%d", digits[i])
		if remainders_list[i] == remainders_list[len(remainders_list)-1] && !repeating {
			if !terminating {
				s += "("
				pos = i
			}
			repeating = true
		}
	}
	if !terminating {
		s += ")"
		rlen = len(remainders_list) - pos - 1
	} else {
		rlen = 0
	}
	return fmt.Sprintf("%s.%s", s[0:1], s[1:]), rlen

}

func main() {

	N := 1000
	max_rlen := 0
	max_den := 0

	num := 1
	for den := 2; den < N; den++ {
		_, rlen := decimal_nutation(num, den)
		if rlen > max_rlen {
			max_rlen = rlen
			max_den = den
		}
	}
	s, rlen := decimal_nutation(num, max_den)
	fmt.Println(fmt.Sprintf("%d/%d =", num, max_den), s, rlen)

}
