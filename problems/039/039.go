/*
Integer right triangles

If p is the perimeter of a right angle triangle with integral length
sides, {a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
*/

package main

import (
	"fmt"
)

func sort(n []int) {

	for {
		done := true
		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				done = false
			}
		}
		if done {
			break
		}
	}
}

func is_ptriple(n []int) bool {
	if n[0]*n[0]+n[1]*n[1] == n[2]*n[2] {
		return true
	}
	return false
}

func main() {
	fmt.Println("Integer right triangles")

	P_MAX := 1000
	P_MIN := 3

	best_p_val := -1
	best_p_count := 0

	for p := P_MAX; p >= P_MIN; p-- {
		p_count := 0
		for a := 1; a < p/2; a++ {
			for b := a; b < (p-a)/2; b++ {
				c := p - a - b
				n := []int{a, b, c}
				sort(n)
				if n[0]+n[1] <= n[2] {
					continue
				}
				if is_ptriple(n) {
					p_count++
				}
			}
		}
		if p_count > best_p_count {
			best_p_val = p
			best_p_count = p_count
		}
	}

	fmt.Println(best_p_val, best_p_count)

}
