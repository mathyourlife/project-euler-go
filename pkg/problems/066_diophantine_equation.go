package problems

import (
	"fmt"
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

func (p *DiophantineEquation) next(n, i, b, c int) (int, int, int) {
	m := (n - b*b) / c
	d := (i + b) / m * m

	return (c * d) / (n - b*b), d - b, (n - b*b) / c
}

func (p *DiophantineEquation) getContinuedFraction(n int) (int, []int) {
	parameters := map[string]bool{}
	var i int
	for i = 1; i < n/2; i++ {
		if (i+1)*(i+1) > n {
			break
		}
	}

	var coeff []int
	var a, b, c int

	b = i
	c = 1

	for {
		a, b, c = p.next(n, i, b, c)
		if parameters[fmt.Sprintf("%d-%d", b, c)] {
			break
		}
		parameters[fmt.Sprintf("%d-%d", b, c)] = true
		coeff = append(coeff, a)
	}
	return i, coeff
}

func (p *DiophantineEquation) Solve() (string, error) {
	// Pell's Equation
	// x^2 - n y^2 = 1
	// where n is not a perfect square
	//
	// Joseph Louis Lagrange proved that, as long as n is not a perfect square,
	// Pell's equation has infinitely many distinct integer solutions. These
	// solutions may be used to accurately approximate the square root of n by
	// rational numbers of the form x/y.

	target := NewBigInt(1)
	maxX := NewBigInt(0)
	solution := uint64(0)

	for D := uint64(1); D <= uint64(1000); D++ {
		if math.Pow(math.Round(math.Sqrt(float64(D))), 2) == float64(D) {
			continue
		}
		i, coeff := p.getContinuedFraction(int(D))

		j := 0
		for {
			x := NewBigInt(1)
			y := NewBigInt(0)
			var s int
			for k := j; k >= 0; k-- {
				x, y = y, x
				if k == 0 {
					s = i
				} else {
					s = coeff[(k-1)%len(coeff)]
				}
				yCopy := y.Copy()
				yCopy.Mul(s)
				x.AddBigInt(yCopy)
				x.Regroup()
			}

			x2 := x.MulBigInt(x)
			y2 := y.MulBigInt(y)

			y2.Mul(int(D))
			if x2.LessThan(y2) {
				j++
				continue
			}
			x2.SubBigInt(y2)

			if x2.Equals(target) {
				if maxX.LessThan(x) {
					maxX = x
					solution = D
				}
				break
			}
			j++
		}
	}
	return fmt.Sprintf("%d", solution), nil
}
