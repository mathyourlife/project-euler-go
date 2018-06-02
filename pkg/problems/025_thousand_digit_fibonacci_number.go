package problems

import (
	"fmt"
)

type ThousandDigitFibonacciNumber struct{}

func (p *ThousandDigitFibonacciNumber) ID() int {
	return 25
}

func (p *ThousandDigitFibonacciNumber) Text() string {
	return `The Fibonacci sequence is defined by the recurrence relation:

F(n) = F(n−1) + F(n−2), where F(1) = 1 and F(2) = 1.

Hence the first 12 terms will be:

    F(1) = 1
    F(2) = 1
    F(3) = 2
    F(4) = 3
    F(5) = 5
    F(6) = 8
    F(7) = 13
    F(8) = 21
    F(9) = 34
    F(10) = 55
    F(11) = 89
    F(12) = 144

The 12th term, F(12), is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to
contain 1000 digits?
`
}

func (p *ThousandDigitFibonacciNumber) Solve() (string, error) {
	N := 1000

	a := NewBigInt(1)
	b := NewBigInt(1)

	// Go back and forth adding a to b and b to a until
	// one of them reaching N digits in length.
	i := 3
	for {
		if i%2 == 0 {
			a.AddBigInt(b)
			a.Regroup()
			if len(a.n) >= N {
				break
			}
		} else {
			b.AddBigInt(a)
			b.Regroup()
			if len(b.n) >= N {
				break
			}
		}
		i++
	}

	// var s string
	// if i%2 == 0 {
	// 	s = a.Print()
	// } else {
	// 	s = b.Print()
	// }
	// fmt.Println(i, s)
	return fmt.Sprintf("%d", i), nil
}
