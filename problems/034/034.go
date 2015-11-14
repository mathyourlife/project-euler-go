/*
Digit factorials

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of
their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/

package main

import (
	"fmt"
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

func (b *BigInt) Compare(a *BigInt) bool {
	if len(a.n) != len(b.n) {
		return false
	}
	for i := 0; i < len(b.n); i++ {
		if b.n[i] != a.n[i] {
			return false
		}
		// fmt.Println(b.n[i])
		// fmt.Println(check_n.n[i])
	}
	return true
}

func main() {

	f := map[int]int{
		0: 1,
	}
	prod := 1
	for i := 1; i < 10; i++ {
		prod *= i
		f[i] = prod
	}

	ans := 0
	for i := 3; i < 1000000; i++ {
		b := NewBigInt(i)
		check := 0
		for _, j := range(b.n) {
			check += f[j]
		}
		if b.Compare(NewBigInt(check)) {
			ans += i
			fmt.Println(b.Print())
		}
	}
	fmt.Println(ans)
}
