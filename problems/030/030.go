/*
Digit fifth powers

Surprisingly there are only three numbers that can be written
as the sum of fourth powers of their digits:

    1634 = 1^4 + 6^4 + 3^4 + 4^4
    8208 = 8^4 + 2^4 + 0^4 + 8^4
    9474 = 9^4 + 4^4 + 7^4 + 4^4

As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum
of fifth powers of their digits
*/

package main

import (
	"fmt"
	"math"
)

func num_digits(n uint) uint {
	return uint(math.Log10(float64(n)) + 1)
}

// Determine if a specific combination of digits works
// if so add to the solutions map
func check_comb(digits []uint, powers map[uint]uint, solutions map[uint]bool) {

	// Calculate the goal sum for these set of digits
	goal_sum := uint(0)
	for _, d := range digits {
		goal_sum += powers[d]
	}

	// Generate an idx array 0,1,2,..,len(digits)
	// idx values will be permuted to match with the sum
	idx := []int{}
	for i, _ := range digits {
		idx = append(idx, i)
	}

	// Generate permutations of the digits and compare to goal sum
	lp := LexPerm{}
	for {
		more := lp.Next(idx)
		num := ""
		for _, i := range idx {
			num += fmt.Sprintf("%d", digits[i])
		}
		if num == fmt.Sprintf("%d", goal_sum) {
			solutions[goal_sum] = true
		}
		if !more {
			break
		}
	}
}

type CombWithRepl struct {
	items []uint
	idx   []int
}

func NewCombWithRepl(items []uint, length uint) *CombWithRepl {
	idx := []int{}
	for i := 0; i < int(length); i++ {
		idx = append(idx, 0)
	}
	return &CombWithRepl{
		items: items,
		idx:   idx,
	}
}

// Iterate to the next arrangement of indicies
// 0000, 1000, ..., 9000, 1100,
func (c *CombWithRepl) next_idx() bool {
	if c.idx[0]+1 == len(c.items) {
		pos := 1
		for {
			if pos == len(c.idx) {
				return false
			}
			if c.idx[pos] != len(c.items)-1 {
				break
			}
			pos++
		}
		c.idx[pos]++
		for i := 0; i < pos; i++ {
			c.idx[i] = c.idx[pos]
		}
	} else {
		c.idx[0]++
	}
	return true
}

// Generate the next combination.  Duplications occur with
// repeated digits.
func (c *CombWithRepl) Next() ([]uint, bool) {
	r := []uint{}
	for _, i := range c.idx {
		r = append(r, c.items[i])
	}
	more := c.next_idx()
	return r, more
}

// Lexicographic Permutations
type LexPerm struct{}

/*
Find the largest index k such that a[k] < a[k + 1].
If no such index exists, the permutation is the last permutation.

Return -1 if items are in descending order.
*/
func (lp LexPerm) get_k(a []int) int {
	k := -1
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			k = i
		}
	}
	return k
}

/*
Find the largest index l greater than k such that a[k] < a[l].
*/
func (lp LexPerm) get_l(a []int, k int) int {
	var l int
	for i := k + 1; i < len(a); i++ {
		if a[k] < a[i] {
			l = i
		}
	}
	return l
}

/*
Modify the provided slice in place, and return false if the slice
is already in descending order (no more permutations exist).
*/
func (lp LexPerm) Next(a []int) bool {
	k := lp.get_k(a)
	if k < 0 {
		return false
	}
	l := lp.get_l(a, k)
	// Step 3 swap k and l
	a[k], a[l] = a[l], a[k]
	// Step 4 revers items for elements > k+1
	for i, j := k+1, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return true
}

func main() {

	N := 5

	// Cache the Nth powers of 0-9
	powers := map[uint]uint{
		uint(0): 0,
		uint(1): 1,
	}
	for b := 2; b < 10; b++ {
		powers[uint(b)] = uint(math.Pow(float64(b), float64(N)))
	}

	// Calculate an upper limit for the number of digits
	// e.g. For the exponent N=4, at 99,999 the sum of the digits^4 is
	// eclipsed by the number the digits represent.
	//       9 => 9^4 = 6561
	//      99 => 9^4 + 9^4 = 13122
	//     999 => 9^4 + 9^4 + 9^4 = 19683
	//   9,999 => 9^4 + 9^4 + 9^4 + 9^4 = 26244
	//  99,999 => 9^4 + 9^4 + 9^4 + 9^4 + 9^4 = 32805
	// 999,999 => 9^4 + 9^4 + 9^4 + 9^4 + 9^4 + 9^4 = 39366

	max_length := uint(1)
	for {
		if num_digits(powers[9]*max_length) < max_length {
			break
		}
		max_length++
	}
	max_length--

	solutions := map[uint]bool{}
	digits := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for length := uint(2); length <= max_length; length++ {
		c := NewCombWithRepl(digits, length)
		for {
			comb, more := c.Next()
			check_comb(comb, powers, solutions)
			if !more {
				break
			}
		}
	}

	sum := uint(0)
	for k, _ := range solutions {
		sum += k
	}
	fmt.Println(sum)
}
