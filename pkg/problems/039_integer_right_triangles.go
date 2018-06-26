package problems

import (
	"fmt"
)

type IntegerRightTriangles struct{}

func (p *IntegerRightTriangles) ID() int {
	return 39
}

func (p *IntegerRightTriangles) Text() string {
	return `If p is the perimeter of a right angle triangle with
integral length sides, {a,b,c}, there are exactly three solutions
for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
`
}

func (p *IntegerRightTriangles) Solve() (string, error) {
	limit := 1000

	squares := map[int]int{}
	for i := 0; i <= limit/2; i++ {
		squares[i] = i * i
	}

	maxPerim := 0
	maxPerimCount := 0
	for perim := 3; perim < limit; perim++ {
		solutions := 0
		for a := 1; a <= perim/2; a++ {
			for b := a + 1; b < perim/2; b++ {
				c := perim - a - b
				if squares[a]+squares[b] == squares[c] {
					solutions++
				}
			}
		}
		if solutions > maxPerimCount {
			maxPerimCount = solutions
			maxPerim = perim
		}
	}
	return fmt.Sprintf("%d", maxPerim), nil
}
