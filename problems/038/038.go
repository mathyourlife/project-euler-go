/*
Pandigital multiples

Take the number 192 and multiply it by each of 1, 2, and 3:

    192 × 1 = 192
    192 × 2 = 384
    192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576.
We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3,
4, and 5, giving the pandigital, 918273645, which is the concatenated
product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed
as the concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/lexperm"
	"strings"
)

func check_pan(base int, pandigits string) bool {

	pan_check := ""
	i := 1
	for {
		pan_check += fmt.Sprintf("%d", base*i)
		if !strings.Contains(pandigits, pan_check) {
			return false
		}
		if len(pan_check) == 9 {
			return true
		}
		if len(pan_check) > 9 {
			break
		}
		i++
	}
	return true
}

func is_concat_prod(digits []int) bool {

	pandigits := ""
	for _, d := range digits {
		pandigits += fmt.Sprintf("%d", d)
	}

	var base int
	for i := 1; i < 6; i++ {
		if i == 1 {
			base = digits[0]
		} else if i == 2 {
			base = digits[0]*10 + digits[1]
		} else if i == 3 {
			base = digits[0]*100 + digits[1]*10 + digits[2]
		} else if i == 4 {
			base = digits[0]*1000 + digits[1]*100 + digits[2]*10 + digits[3]
		} else if i == 5 {
			base = digits[0]*10000 + digits[1]*1000 + digits[2]*100 + digits[3]*10 + digits[4]
		}
		if check_pan(base, pandigits) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Pandigital multiples")
	lp := lexperm.LexPerm{}
	digits := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for {

		if is_concat_prod(digits) {
			fmt.Println("Found it")
			fmt.Println(digits)
			break
		}

		more := lp.Prev(digits)
		if !more {
			break
		}
	}
}
