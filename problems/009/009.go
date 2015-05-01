/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.

Find the product abc.


c = 1000 - a - b

a^2 + b^2 = (1000 - a - b)^2

0 < a < 1000
0 < b < 1000

0 < a^2 < 1000000
0 < b^2 < 1000000

*/

package main

import (
	"fmt"
)

func main() {
	N := 1000
Search:
	for c := N; c > 1; c-- {
		for b := 1; b <= c; b++ {
			if b + c >= N {
				break
			}
			a := N - b - c
			if a > b {
				continue
			}
			if a*a+b*b == c*c {
				fmt.Printf("%d*%d*%d=%d\n",a, b, c, a*b*c)
				break Search
			}
		}
	}

}
