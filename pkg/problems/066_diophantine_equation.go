package problems

import (
	"fmt"
	"log"
	"math"
)

type DiophantineEquation struct{}

func (p *DiophantineEquation) ID() int {
	return 66
}

func (p *DiophantineEquation) Text() string {
	return `Consider quadratic Diophantine equations of the form:

x^2 – Dy^2 = 1

For example, when D=13, the minimal solution in x is 649^2 – 13×180^2 = 1.

It can be assumed that there are no solutions in positive integers when D is square.

By finding minimal solutions in x for D = {2, 3, 5, 6, 7}, we obtain the following:

3^2 – 2×2^2 = 1
2^2 – 3×1^2 = 1
9^2 – 5×4^2 = 1
5^2 – 6×2^2 = 1
8^2 – 7×3^2 = 1

Hence, by considering minimal solutions in x for D ≤ 7, the largest x is
obtained when D=5.

Find the value of D ≤ 1000 in minimal solutions of x for which the largest
value of x is obtained.
`
}

func (p *DiophantineEquation) trial1() (string, error) {
	var squares []uint64
	isSquare := map[uint64]bool{}
	minSolutions := map[uint64]uint64{}
	maxD := uint64(60)
	dCount := int(maxD - uint64(math.Sqrt(float64(maxD))))

	for n := uint64(1); n < uint64(1000000); n++ {
		squares = append(squares, n*n)
		isSquare[n*n] = true
	}

	found := 0
	dx := make([]uint64, maxD+1)
	for i := 0; i < len(dx); i++ {
		if i*i < len(dx)-1 {
			dx[i*i] = 1
		}
	}
	for _, x2 := range squares {
		if x2 == 1 {
			continue
		}
		// log.Println("x2", x2)
		for d, check := range dx {
			if d == 0 || check != 0 {
				continue
			}
			y2 := (x2 - 1) / uint64(d)
			if x2/uint64(d) == 1 {
				break
			}
			if isSquare[y2] && x2-uint64(d)*y2 == 1 {
				found++
				dx[d] = x2
				log.Printf("%3d: %d - %dx%d = 1", d, x2, d, y2)
			}
		}
		if x2 < maxD+1 && dx[x2-1] == 0 {
			found++
			d := x2 - 1
			dx[d] = x2
			log.Printf("%3d: %d - %dx%d = 1", d, x2, d, 1)
		}
		if found == dCount {
			log.Println("found 'em")
			break
		}
	}
	log.Println(dx)
	missed := 0
	for _, x := range dx {
		if x == 0 {
			missed++
		}
	}
	log.Println("missed", missed)
	return "", nil

	for i, x := range squares {
		for j := i + 1; j < len(squares); j++ {
			y := squares[j]
			// log.Println(i, j, x, y)
			if (y-1)%x == 0 {
				D := (y - 1) / x
				if D > maxD || minSolutions[D] > 0 {
					break
				}
				minSolutions[D] = x
				log.Printf("%d - %dx%d = 1", y, (y-1)/x, x)
			}
		}
		if len(minSolutions) == dCount {
			log.Println("found 'em")
			break
		}
	}
	log.Println(len(minSolutions), minSolutions)
	return "", nil
}

func (p *DiophantineEquation) Solve() (string, error) {
	d := 10

	x := 1
	for {
		c := (x*x - 1)
		if c%d != 0 {
			continue
		}
		c /= d
		// math.Sqrt()
		x++
	}
	return fmt.Sprintf("%d", 0), nil
}
