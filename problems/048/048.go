/*
Self powers

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
*/

package main

import (
	"fmt"
	"log"
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
	if len(a.n) <= len(b.n) {
		for i, n := range a.n {
			b.n[i] += n
		}
	} else {
		for i := 0; i < len(b.n); i++ {
			b.n[i] += a.n[i]
		}
		for i := len(b.n); i < len(a.n); i++ {
			b.n = append(b.n, a.n[i])
		}
	}
	b.Regroup()
}

func (b *BigInt) String() string {
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
	}
	return true
}

func main() {
	log.Println("Self powers")

	var m *BigInt
	s := NewBigInt(0)

	for n := 1; n <= 1000; n++ {
		// Calculate the self power
		m = NewBigInt(n)
		for i := 1; i < n; i++ {
			m.Mul(n)
		}
		// Add it to the running sum
		s.AddBig(m)
	}
	log.Println(s)
}
