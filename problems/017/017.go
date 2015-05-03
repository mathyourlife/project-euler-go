/*
If the numbers 1 to 5 are written out in words: one, two, three, four,
five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written
out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
forty-two) contains 23 letters and 115 (one hundred and fifteen) contains
20 letters. The use of "and" when writing out numbers is in compliance
with British usage.

Note: Current implementation of BigInt.Name() only intended for 3 digit numbers
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

	words := map[string]int{
		"one":      1,
		"thousand": 1,
	}
	n := NewBigInt(1)
	for i := 0; i < 999; i++ {
		name := n.Name()
		for _, word := range name {
			words[word]++
		}
		n.Add(1)
	}
	length := 0
	for word, count := range words {
		length += len(word) * count
	}
	fmt.Println(length)

}
