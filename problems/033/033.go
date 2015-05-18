/*
Digit cancelling fractions

The fraction 49/98 is a curious fraction, as an inexperienced
mathematician in attempting to simplify it may incorrectly
believe that 49/98 = 4/8, which is correct, is obtained by
cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of
fraction, less than one in value, and containing two digits
in the numerator and denominator.

If the product of these four fractions is given in its lowest
common terms, find the value of the denominator.
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/primefactorization"
	"math"
)

func GCF(n []uint64, pf *primefactorization.PrimeFactorization) map[uint64]uint64 {

	gcf := map[uint64]uint64{}
	for i, a := range n {
		pf_n := pf.Of(a)
		if i == 0 {
			for k, v := range pf_n {
				gcf[k] = v
			}
		} else {
			for k, v := range pf_n {
				if v < gcf[k] {
					gcf[k] = v
				}
			}
			for k, _ := range gcf {
				if pf_n[k] == 0 {
					delete(gcf, k)
				}
			}
		}
	}
	return gcf
}

func main() {
	pf := primefactorization.NewPrimeFactorization()

	prod_num := uint64(1)
	prod_den := uint64(1)
	for num := uint64(11); num < 100; num++ {
		for den := num + 1; den < 100; den++ {
			if num%10 == 0 || den%10 == 0 {
				continue
			}
			curious := false
			dec := float64(num) / float64(den)
			if num/10 == den/10 && float64(num%10)/float64(den%10) == dec {
				curious = true
			}
			if num/10 == den%10 && float64(num%10)/float64(den/10) == dec {
				curious = true
			}
			if num%10 == den/10 && float64(num/10)/float64(den%10) == dec {
				curious = true
			}
			if num%10 == den%10 && float64(num/10)/float64(den/10) == dec {
				curious = true
			}
			if curious {
				fmt.Println(num, "/", den)
				prod_num *= num
				prod_den *= den
			}
		}
	}

	den := pf.Of(prod_den)

	gcf := GCF([]uint64{prod_num, prod_den}, pf)
	sol := float64(1)
	for k, v := range den {
		sol *= math.Pow(float64(k), float64(v-gcf[k]))
	}
	fmt.Println(sol)
}
