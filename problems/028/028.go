/*
Number spiral diagonals

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

*/
package main

import (
	"fmt"
)

func move(r int, c int, direction int) (int, int) {
	if direction == 0 {
		c += 1
	} else if direction == 1 {
		r += 1
	} else if direction == 2 {
		c -= 1
	} else {
		r -= 1
	}
	return r, c
}

func create_grid() [1001][1001]int {
	size := 1001 // needs to be odd for a center
	grid := [1001][1001]int{}

	direction := 0 // 0=right, 1=down, 2=left, 3=up

	r := ((size + 1) / 2) - 1
	c := ((size + 1) / 2) - 1
	grid[r][c] = 1

	count := 2
	steps := 1
Spiral:
	for {
		for s := 0; s < steps; s++ {
			r, c = move(r, c, direction)
			if r >= size || c >= size || r < 0 || c < 0 {
				break Spiral
			}
			grid[r][c] = count
			count++
		}
		direction = (direction + 1) % 4

		for s := 0; s < steps; s++ {
			r, c = move(r, c, direction)
			if r >= size || c >= size || r < 0 || c < 0 {
				break Spiral
			}
			grid[r][c] = count
			count++
		}
		direction = (direction + 1) % 4
		steps++
	}

	return grid
}

func main() {
	grid := create_grid()

	sum := 0
	for r, c := 0, 0; r < 1001 || c < 1001; r, c = r+1, c+1 {
		sum += grid[r][c]
		sum += grid[1001-r-1][c]
	}
	sum-- // double count center
	fmt.Println(sum)
}
