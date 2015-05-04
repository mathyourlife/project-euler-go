/*
A perfect number is a number for which the sum of its proper divisors
is exactly equal to the number. For example, the sum of the proper
divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means
that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is
less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the
smallest number that can be written as the sum of two abundant
numbers is 24. By mathematical analysis, it can be shown that all
integers greater than 28123 can be written as the sum of two
abundant numbers. However, this upper limit cannot be reduced any
further by analysis even though it is known that the greatest
number that cannot be expressed as the sum of two abundant numbers
is less than this limit.

Find the sum of all the positive integers which cannot be written
as the sum of two abundant numbers.
*/

package main

import (
	"fmt"
	"github.com/mathyourlife/lt3maths/primefactorization"
	"math"
)

func divisors(n uint64, pf *primefactorization.PrimeFactorization) []uint64 {
	ds := map[uint64]bool{
		1: true,
	}
	for k, v := range pf.Of(n) {
		new_divisors := []uint64{}
		for divisor, _ := range ds {
			for i := uint64(0); i < v; i++ {
				multiple := k * uint64(math.Pow(float64(k), float64(i)))
				new_divisors = append(new_divisors, divisor*multiple)
			}
		}
		for _, divisor := range new_divisors {
			ds[divisor] = true
		}
	}

	list := []uint64{}
	for n, _ := range ds {
		list = append(list, n)
	}
	return list
}

func main() {

	N := uint64(43)
	N = uint64(28123)

	pf := primefactorization.NewPrimeFactorization()

	// Populate divisors for numbers 2 to N
	ds := map[uint64][]uint64{}
	for i := uint64(2); i < N; i++ {
		pf.Of(i)
		ds[i] = divisors(i, pf)
	}

	// Determine if numbers between 2 and N are abundant
	abundant := []uint64{}
	for n, ndivisors := range ds {
		sum := uint64(0)
		for _, d := range ndivisors {
			sum += d
		}
		sum -= n
		if n < sum {
			abundant = append(abundant, n)
		}
	}

	// Sums of 2 abundant numbers
	sums := map[uint64]bool{}
	for i := 0; i < len(abundant)-1; i++ {
		for j := i; j < len(abundant); j++ {
			sums[abundant[i]+abundant[j]] = true
		}
	}

	// Sum of all numbers that can be written as the sum of 2 abundant numbers
	sum_abundants := uint64(0)
	for n, _ := range sums {
		sum_abundants += n
	}

	final := uint64(0)
	for i := uint64(1); i < N; i++ {
		if !sums[i] {
			final += i
		}
	}
	fmt.Println(final)
}
