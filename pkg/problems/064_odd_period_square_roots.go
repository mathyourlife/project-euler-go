package problems

import (
	"fmt"
)

type OddPeriodSquareRoots struct{}

func (p *OddPeriodSquareRoots) ID() int {
	return 64
}

func (p *OddPeriodSquareRoots) Text() string {
	return `All square roots are periodic when written as continued fractions
and can be written in the form:

√N = a_0 + 1/(a_1 + 1/(a_2 + 1/(a_3 + ...)))

For example, let us consider √23:

√23 = 4 + √23 - 4 = 4 + 1/(1/(√23-4))
                            = 4 + 1/(1 + (√23-3)/7)

if we continue, we would get the following expansion:

√23 = 4 +            1
          --------------------
					1 +           1
					    ----------------
							3 +         1
							    ------------
									1 +      1
									    --------
											8 + ...

where a_0 = 4, a_1 = 1, a_2 = 3, a_3 = 1, a_4 = 8, a_5 = 1, a_6 = 3, a_7 = 1

It can be seen that the sequence is repeating.  For concicseness, we use the
notation √23 = [4;(1,3,1,8)] to indicate that the block (1,3,1,8) repeats
indefinitely.

The first ten continued fraction representitians of (irrational) square roots are:

√2=[1;(2)],          period=1
√3=[1;(1,2)],        period=2
√5=[2;(4)],          period=1
√6=[2;(2,4)],        period=2
√7=[2;(1,1,1,4)],    period=4
√8=[2;(1,4)],        period=2
√10=[3;(6)],         period=1
√11=[3;(3,6)],       period=2
√12= [3;(2,6)],      period=2
√13=[3;(1,1,1,1,6)], period=5

Exactly four continued fractions, for N ≤ 13, have an odd period.

How many continued fractions for N ≤ 10000 have an odd period?
`
}

func (p *OddPeriodSquareRoots) next(N, i, b, c int) (int, int, int) {
	m := (N - b*b) / c
	d := (i + b) / m * m

	return (c * d) / (N - b*b), d - b, (N - b*b) / c
}

func (p *OddPeriodSquareRoots) getContinuedFraction(N int) (int, []int) {
	parameters := map[string]bool{}
	var i int
	for i = 1; i < N/2; i++ {
		if (i+1)*(i+1) > N {
			break
		}
	}

	var coeff []int
	var a, b, c int

	b = i
	c = 1

	for {
		a, b, c = p.next(N, i, b, c)
		if parameters[fmt.Sprintf("%d-%d", b, c)] {
			break
		}
		parameters[fmt.Sprintf("%d-%d", b, c)] = true
		coeff = append(coeff, a)
	}
	return i, coeff
}

func (p *OddPeriodSquareRoots) Solve() (string, error) {

	odds := 0
	s := 2
	for N := 2; N <= 10000; N++ {
		// skip over perfect squares
		if N == s*s {
			s++
			continue
		}
		_, coeff := p.getContinuedFraction(N)
		if len(coeff)%2 == 1 {
			odds++
		}
	}

	return fmt.Sprintf("%d", odds), nil
}
