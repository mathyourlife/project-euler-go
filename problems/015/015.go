/*
Starting in the top left corner of a 2×2 grid, and only being able to
move to the right and down, there are exactly 6 routes to the
bottom right corner.

How many such routes are there through a 20×20 grid?
*/

package main

import (
	"fmt"
)

func main() {
	size := 20 + 1
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if r == 0 && c == 0 {
				grid[r][c] = 1
				continue
			}
			if r > 0 {
				grid[r][c] += grid[r-1][c]
			}
			if c > 0 {
				grid[r][c] += grid[r][c-1]
			}
		}
	}
	fmt.Println(grid[size-1][size-1])
}
