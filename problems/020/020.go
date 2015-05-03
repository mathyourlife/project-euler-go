/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
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
	N := 100

	b := NewBigInt(1)

	for i := 1; i <= N; i++ {
		b.Mul(i)
	}
	fmt.Println(b.Print())
	sum := 0
	for _, d := range b.n {
		sum += d
	}
	fmt.Println(sum)

}
