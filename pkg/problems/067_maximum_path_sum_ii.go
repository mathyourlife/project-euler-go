package problems

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type MaximumPathSumII struct {
	gridFile string
	grid     [][]int64
}

func (p *MaximumPathSumII) ID() int {
	return 67
}

func (p *MaximumPathSumII) Text() string {
	return `By starting at the top of the triangle below and moving to adjacent
numbers on the row below, the maximum total from top to bottom is 23.

   3
  7 4
 2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and
'Save Link/Target As...'), a 15K text file containing a triangle with
one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible
to try every route to solve this problem, as there are 299 altogether! If you
could check one trillion (1012) routes every second it would take over twenty
billion years to check them all. There is an efficient algorithm to solve
it. ;o)
`
}

func (p *MaximumPathSumII) Solve() (string, error) {

	file, err := os.Open(p.gridFile)
	if err != nil {
		return "", err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	p.load(strings.TrimSpace(string(b)))

	// Collapse rows from the bottom up using the larger number
	// For the bottom left corner
	//   63
	// 04  62
	// The 62 is larger so gets added to the 63 above
	for row := 0; row < len(p.grid)-1; row++ {
		for col := 0; col < len(p.grid[row])-1; col++ {
			if p.grid[row][col] > p.grid[row][col+1] {
				p.grid[row+1][col] += p.grid[row][col]
			} else {
				p.grid[row+1][col] += p.grid[row][col+1]
			}
		}
	}

	return fmt.Sprintf("%d", p.grid[len(p.grid)-1][0]), nil
}

func (p *MaximumPathSumII) load(s string) {
	rows := strings.Split(s, "\n")

	p.grid = make([][]int64, 0, len(rows))
	for i := len(rows) - 1; i >= 0; i-- {
		r := []int64{}
		rnums := strings.Split(rows[i], " ")
		for _, ns := range rnums {
			n, err := strconv.ParseInt(ns, 10, 16)
			if err != nil {
				panic(err)
			}
			r = append(r, n)
		}
		p.grid = append(p.grid, r)
	}
}
