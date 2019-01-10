package problems

import (
	"fmt"
	"math"
)

type Magic5GonRing struct{}

func (p *Magic5GonRing) ID() int {
	return 68
}

func (p *Magic5GonRing) Text() string {
	return `Consider the following "magic" 3-gon ring, filled with the numbers
1 to 6, and each line adding to nine.

                 4
								  \
									 3
									/ \
								 1 - 2 - 6
								/
							 5

Working clockwise, and starting from the group of three with the numerically
lowest external node (4,3,2 in this example), each solution can be described
uniquely. For example, the above solution can be described by the set:
4,3,2; 6,2,1; 5,1,3.

It is possible to complete the ring with four different totals: 9, 10, 11,
and 12. There are eight solutions in total.

Total	Solution Set
9 	4,2,3; 5,3,1; 6,1,2
9 	4,3,2; 6,2,1; 5,1,3
10	2,3,5; 4,5,1; 6,1,3
10	2,5,3; 6,3,1; 4,1,5
11	1,4,6; 3,6,2; 5,2,4
11	1,6,4; 5,4,2; 3,2,6
12	1,5,6; 2,6,4; 3,4,5
12	1,6,5; 3,5,4; 2,4,6

By concatenating each group it is possible to form 9-digit strings; the
maximum string for a 3-gon ring is 432621513.

Using the numbers 1 to 10, and depending on arrangements, it is possible to
form 16- and 17-digit strings. What is the maximum 16-digit string for
a "magic" 5-gon ring?
`
}

func (p *Magic5GonRing) check(values []int) bool {
	spokes := [][]int{
		{5, 0, 1},
		{6, 1, 2},
		{7, 2, 3},
		{8, 3, 4},
		{9, 4, 0},
	}
	// spokes := [][]int{
	// 	{3, 0, 1},
	// 	{4, 1, 2},
	// 	{5, 2, 0},
	// }
	check := 0
	for _, spoke := range spokes {
		s := 0
		for _, idx := range spoke {
			s += values[idx]
		}
		if check == 0 {
			check = s
			continue
		}
		if s != check {
			return false
		}
	}
	return true
}

func (p *Magic5GonRing) valuesToSpokes(values []int) [][]int {
	// spokes := [][]int{
	// 	{values[3], values[0], values[1]},
	// 	{values[4], values[1], values[2]},
	// 	{values[5], values[2], values[0]},
	// }
	spokes := [][]int{
		{values[5], values[0], values[1]},
		{values[6], values[1], values[2]},
		{values[7], values[2], values[3]},
		{values[8], values[3], values[4]},
		{values[9], values[4], values[0]},
	}

	minEnd := 0
	minSpoke := 0
	for i, spoke := range spokes {
		if minEnd == 0 || spoke[0] < minEnd {
			minEnd = spoke[0]
			minSpoke = i
		}
	}

	sortedSpokes := [][]int{}

	for i := 0; i < len(spokes); i++ {
		sortedSpokes = append(sortedSpokes, spokes[(i+minSpoke)%len(spokes)])
	}
	return sortedSpokes
}

func (p *Magic5GonRing) Solve() (string, error) {

	maxSolution := uint64(0)

	// Limit to 16 digit solutions
	maxLen := uint64(10000000000000000)

	// values := []int{1, 2, 3, 4, 5, 6}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for {
		if p.check(values) {
			sorted := p.valuesToSpokes(values)
			concat := uint64(0)
			for i := 0; i < len(sorted)*len(sorted[0]); i++ {
				v := uint64(sorted[i/3][i%3])
				for j := numDigits(v) - 1; j >= 0; j-- {
					base := uint64(math.Pow(10, float64(j)))
					concat *= uint64(10)
					concat += uint64((v / base) % 10)
				}
			}
			if concat > maxSolution && concat < maxLen {
				maxSolution = concat
			}
		}

		if !LexPerm(values) {
			break
		}
	}

	return fmt.Sprintf("%d", maxSolution), nil
}
