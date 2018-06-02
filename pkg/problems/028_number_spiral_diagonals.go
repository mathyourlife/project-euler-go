package problems

import (
	"fmt"
)

type NumberSpiralDiagonals struct{}

func (p *NumberSpiralDiagonals) ID() int {
	return 28
}

func (p *NumberSpiralDiagonals) Text() string {
	return `Number spiral diagonals

Starting with the number 1 and moving to the right in a
clockwise direction a 5 by 5 spiral is formed as follows:

[21] 22  23  24 [25]
 20 [ 7]  8 [ 9] 10
 19   6 [ 1]  2  11
 18 [ 5]  4 [ 3] 12
[17] 16  15  14 [13]

It can be verified that the sum of the numbers on the
diagonals is 101.

What is the sum of the numbers on the diagonals in a
1001 by 1001 spiral formed in the same way?
`
}

// Solve - The numbers along the diagonals increase in a pattern.
// first 3x3 spiral increasing by 2:   3,  5,  7,  9
// second 5x5 spiral increasing by 4: 13, 17, 21, 25
// third 7x7 spiral increasing by 6:  31, 37, 43, 49
//
// The spirals add the following vaules respectively
// 24, 76, 160, 276, 424, ...
//
// which can be rewritten as:
// 20+4(1)², 40+4(3)², 60+4(5)², 80+4(7)², ...
// or
// s(n) = 20n+4(2n-1)² for n >= 1 and s(0) = 1
// or
// s(n) = 16n²+4n+4 for n >= 1 and s(0) = 1
func (p *NumberSpiralDiagonals) Solve() (string, error) {
	spiralDiagonalSum := func(n uint64) uint64 {
		return 16*n*n + 4*n + 4
	}
	sum := uint64(1)
	// Need to add 500 spirals to increase up to a 1001x1001
	for n := uint64(1); n <= uint64(500); n++ {
		sum += spiralDiagonalSum(n)
	}

	return fmt.Sprintf("%d", sum), nil
}
