package problems

import (
	"fmt"
)

type LexicographicPermutations struct{}

func (p *LexicographicPermutations) ID() int {
	return 24
}

func (p *LexicographicPermutations) Text() string {
	return `A permutation is an ordered arrangement of objects. For
example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If
all of the permutations are listed numerically or alphabetically, we call
it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits
0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
`
}

// Solve use lexicographic permutations to order digits 0-9
// Based on http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
//
// 1) Find the largest index k such that a[k] < a[k + 1].
//    If no such index exists, the permutation is the last permutation.
// 2) Find the largest index l greater than k such that a[k] < a[l].
// 3) Swap the value of a[k] with that of a[l].
// 4) Reverse the sequence from a[k + 1] up to and including the final
//    element a[n].
func (p *LexicographicPermutations) Solve() (string, error) {
	N := 1000000
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// LexPerm has no maintained state so items are assumed in ascending
	// order upon entry of the permutation generator.
	for i := 1; i < N; i++ {
		more := LexPerm(a)
		if !more {
			break
		}
	}

	n := ""
	for _, v := range a {
		n = n + fmt.Sprintf("%d", v)
	}

	return n, nil
}
