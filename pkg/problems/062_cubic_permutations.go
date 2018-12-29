package problems

import (
	"fmt"
	"log"
	"sort"
)

type CubicPermutations struct{}

func (p *CubicPermutations) ID() int {
	return 62
}

func (p *CubicPermutations) Text() string {
	return `The cube 41063625 (345^3) can be permuted to produce two other cubes:
56623104 (384^3) and 66430125 (405^3). In fact, 41063625 is the smallest cube
which has exactly three permutations of its digits which are also cube.

Find the smallest cube for which exactly five permutations of its digits
are cube.`
}

func (p *CubicPermutations) orderDigits(n uint64) uint64 {
	digits := []uint64{}

	for {
		if n <= 0 {
			break
		}
		digits = append(digits, n%10)
		n /= 10
	}
	sort.Sort(UInt64List(digits))

	ordered := uint64(0)
	for i := len(digits) - 1; i >= 0; i-- {
		ordered += digits[i]
		if i > 0 {
			ordered *= 10
		}
	}

	return ordered
}

func (p *CubicPermutations) Solve() (string, error) {
	orderedSets := map[uint64][]uint64{}
	cubeRoot := map[uint64]uint64{}
	// calculating cubes up to 1000 takes 0.007s
	for i := uint64(0); i < uint64(10000); i++ {
		cube := i * i * i
		cubeRoot[cube] = i
		ordered := p.orderDigits(cube)
		cubes := orderedSets[ordered]
		if cubes == nil {
			cubes = []uint64{}
		}
		orderedSets[ordered] = append(cubes, cube)
	}
	min := uint64(0)
	for ordered, cubes := range orderedSets {
		if len(cubes) == 5 {
			if cubes[0] < min || min == uint64(0) {
				min = cubes[0]
			}
			log.Println(min, ordered, cubes)
		}
	}
	return fmt.Sprintf("%d", min), nil
}
