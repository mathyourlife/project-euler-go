/*
Pandigital products

We shall say that an n-digit number is pandigital if it makes
use of all the digits 1 to n exactly once; for example, the
5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254,
containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product
identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be
sure to only include it once in your sum.
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/lexperm"
	"math"
)

func conv(ds []int) int {
	n := 0
	for k, v := range ds {
		n += v * int(math.Pow(10, float64(len(ds)-k-1)))
	}
	return n
}

func check(d []int, products map[int]bool) {
	for prod := 1; prod < len(d); prod++ {
		for equal := prod + 1; equal < len(d); equal++ {
			if conv(d[0:prod])*conv(d[prod:equal]) == conv(d[equal:len(d)]) {
				products[conv(d[equal:len(d)])] = true
				fmt.Println(conv(d[0:prod]), "*", conv(d[prod:equal]), "=", conv(d[equal:len(d)]))
			}
		}
	}
}

func main() {
	products := map[int]bool{}
	lp := lexperm.LexPerm{}
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for {
		check(d, products)
		more := lp.Next(d)
		if !more {
			break
		}
	}
	sum := 0
	for k, _ := range products {
		sum += k
	}
	fmt.Println(sum)
}
