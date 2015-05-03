/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package main

import (
	"fmt"
)

type BigInt struct {
	n []int
}

func NewBigInt(n int) *BigInt {
	return &BigInt{
		n: []int{n},
	}
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
}

func (b *BigInt) Print() string {
	s := ""
	for _, d := range b.n {
		s = fmt.Sprintf("%d%s", d, s)
	}
	return s
}

func main() {

	n := NewBigInt(1)

	for i := 0; i < 1000; i++ {
		n.Mul(2)
		n.Regroup()
	}

	sum := 0
	for _, d := range n.n {
		sum += d
	}
	fmt.Println(sum)
}
