/*
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten
natural numbers and the square of the sum is 3025 - 385 = 2640.  Find
the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.

(1 + 2 + 3 + 4)^2 - (1^2 + 2^2 + 3^2 + 4^2)

4 : 1*4 | 2*4 | 3*4 |
3 : 1*3 | 2*3 |     | 4*3
2 : 1*2 |     | 3*2 | 4*2
1 :     | 2*1 | 3*1 | 4*1
     1     2     3     4

*/

package main

import (
	"fmt"
)

func main() {
	max := 100
	min := 1
	var a, b int

	total := 0
	for sum := max * 2; sum >= 2*min; sum-- {
		if sum > max+min {
			a, b = sum-max, max
		} else {
			a, b = min, sum-min
		}
		for {
			if a != b {
				total += a * b
			}
			a++
			b--
			if a > max || b < min {
				break
			}
		}
	}
	fmt.Println(total)
}
