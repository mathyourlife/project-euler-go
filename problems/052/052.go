/*
Permuted multiples

It can be seen that the number, 125874, and its double, 251748, contain
exactly the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
contain the same digits.
*/
package main

import (
	"log"
	"math"
	"os"
)

// CountDigits - Break a number out into a map containing the frequency
// of its digits.
func CountDigits(num int) map[int]int {
	place := 10
	length := int(math.Log10(float64(num))) + 1
	digits := map[int]int{}

	for i := 0; i < length; i++ {
		digit := num - num/place*place
		digits[int(digit*10/place)]++
		num -= digit
		place *= 10
	}
	return digits
}

// EqualMaps - simple map comparison
func EqualMaps(m1, m2 map[int]int) bool {
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

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("Permuted multiples")

	numbers := map[int]map[int]int{}

	num := 0

SEARCH_LOOP:
	for {
		num++
		if len(numbers[num]) == 0 {
			numbers[num] = CountDigits(num)
		}
		numbers[2*num] = CountDigits(2 * num)
		if !EqualMaps(numbers[num], numbers[2*num]) {
			continue
		}
		numbers[3*num] = CountDigits(3 * num)
		if !EqualMaps(numbers[num], numbers[3*num]) {
			continue
		}
		numbers[4*num] = CountDigits(4 * num)
		if !EqualMaps(numbers[num], numbers[4*num]) {
			continue
		}
		numbers[5*num] = CountDigits(5 * num)
		if !EqualMaps(numbers[num], numbers[5*num]) {
			continue
		}
		numbers[6*num] = CountDigits(6 * num)
		if !EqualMaps(numbers[num], numbers[6*num]) {
			continue
		}

		log.Println(num, 2*num, 3*num, 4*num, 5*num, 6*num)
		break SEARCH_LOOP
	}
}
