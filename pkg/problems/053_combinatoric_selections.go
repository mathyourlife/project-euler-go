package problems

import (
	"fmt"
)

type CombinatoricSelections struct{}

func (p *CombinatoricSelections) ID() int {
	return 53
}

func (p *CombinatoricSelections) Text() string {
	return `There are exactly ten ways of selecting three from
five, 12345:
123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, 5C3 = 10.

In general,
nCr = n! / r!(n−r)!
where
r ≤ n, n! = n×(n−1)×...×3×2×1, and 0! = 1.

It is not until n = 23, that a value exceeds one-million:
23C10 = 1144066.

How many, not necessarily distinct, values of  nCr,
for 1 ≤ n ≤ 100, are greater than one-million?
`
}

// Based on Pascal's triangle
//  5 C 3
// n=0       1
// n=1     1   1
// n=2    1  2  1
// n=3   1  3  3  1
// n=4  1  4  6  4  1
// n=5 1  5 10 10  5  1
func (p *CombinatoricSelections) Solve() (string, error) {

	lastRow := []uint64{1, 1}

	count := 0
	limit := uint64(1000000)
	N := 100

	for j := 1; j < N; j++ {
		nextRow := make([]uint64, len(lastRow)+1)
		nextRow[0] = 1
		nextRow[len(nextRow)-1] = 1
		for i := 0; i < len(lastRow)-1; i++ {
			val := lastRow[i] + lastRow[i+1]
			if val > limit {
				count++
			}
			nextRow[i+1] = val
		}
		lastRow = nextRow
	}

	return fmt.Sprintf("%d", count), nil
}
