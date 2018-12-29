package problems

import (
	"fmt"
)

type PowerfulDigitCounts struct{}

func (p *PowerfulDigitCounts) ID() int {
	return 63
}

func (p *PowerfulDigitCounts) Text() string {
	return `The 5-digit number 16807=7^5, is also a fifth power. Similarly, the
9-digit number, 134217728=8^9, is a ninth power.

How many n-digit positive integers exist which are also an nth power?`
}

func (p *PowerfulDigitCounts) Solve() (string, error) {
	var i uint64
	tally := 0
	for nth := 1; nth <= 100; nth++ {
		found := 0
		i = uint64(1)
		for {
			pow := NewBigInt(1)
			for j := 0; j < nth; j++ {
				pow.Mul(int(i))
				pow.Regroup()
			}
			digits := len(pow.n)
			if digits > nth {
				break
			}
			if digits == nth {
				found++
			}
			i++
		}
		tally += found
		if found == 0 {
			break
		}
	}
	return fmt.Sprintf("%d", tally), nil
}
