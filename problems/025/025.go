/*
1000-digit Fibonacci number

The Fibonacci sequence is defined by the recurrence relation:

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
*/

package main

import (
	"fmt"
	"math"
)

type BigInt struct {
	n []int
}

func NewBigInt(n int) *BigInt {
	b := &BigInt{
		n: []int{n},
	}
	if n >= 10 {
		b.Regroup()
	}
	return b
}

func (b *BigInt) Regroup() {
	for i := 0; i < len(b.n); i++ {
		if b.n[i] > 9 {
			if i == len(b.n)-1 {
				b.n = append(b.n, 0)
			}
			regroup := b.n[i] / 10
			b.n[i+1] += regroup
			b.n[i] = b.n[i] % 10
		}
	}
}

func (b *BigInt) Mul(f int) {
	for i, _ := range b.n {
		b.n[i] *= f
	}
	b.Regroup()
}

func (b *BigInt) Add(f int) {
	b.n[0] += f
	b.Regroup()
}

func (b *BigInt) AddBig(a *BigInt) {
	for i, n := range a.n {
		b.n[i] += n
	}
	b.Regroup()
}

func (b *BigInt) Print() string {
	s := ""
	for _, d := range b.n {
		s = fmt.Sprintf("%d%s", d, s)
	}
	return s
}

func (b *BigInt) Name() []string {

	var a int
	for i := 0; i < 3 && i < len(b.n); i++ {
		a += b.n[i] * int(math.Pow(10, float64(i)))
	}

	return b.hundreds_name(a)
}

func (b *BigInt) hundreds_name(a int) []string {
	if len(b.n) == 1 && b.n[0] == 0 {
		return []string{"zero"}
	}

	digit := map[int]string{
		0:  "",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}
	tens := map[int]string{
		2: "twenty",
		3: "thirty",
		4: "forty",
		5: "fifty",
		6: "sixty",
		7: "seventy",
		8: "eighty",
		9: "ninety",
	}

	name := []string{}
	c := a % 100
	if c < 20 && c > 0 {
		name = append(name, digit[c])
	} else if b.n[0] == 0 && b.n[1] == 0 {
	} else {
		if !(b.n[0] == 0) {
			name = append(name, digit[b.n[0]])
		}
		if !(b.n[1] == 0) {
			name = append(name, tens[b.n[1]])
		}
	}

	if a > 99 {
		if !(b.n[0] == 0 && b.n[1] == 0) {
			name = append(name, "and")
		}
		name = append(name, "hundred")
		name = append(name, digit[b.n[2]])
	}
	return name
}

func main() {

	N := 1000

	a := NewBigInt(1)
	b := NewBigInt(1)

	j := 2
	for {
		j++
		for i := 0; i < len(a.n); i++ {
			a.n[i], b.n[i] = b.n[i], b.n[i]+a.n[i]
		}
		for i := len(a.n); i < len(b.n); i++ {
			a.n = append(a.n, b.n[i])
		}
		a.Regroup()
		b.Regroup()
		if len(b.n) >= N {
			break
		}
	}

	fmt.Println(j)
}
