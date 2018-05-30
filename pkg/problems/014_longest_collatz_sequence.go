package problems

import (
	"fmt"
)

type LongestCollatzSequence struct {
	links    map[int64]int64
	chainLen map[int64]int64
}

func (p *LongestCollatzSequence) ID() int {
	return 14
}

func (p *LongestCollatzSequence) Text() string {
	return `The following iterative sequence is defined for the
set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the
following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains
10 terms. Although it has not been proved yet (Collatz Problem), it is
thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
`
}

func (p *LongestCollatzSequence) Solve() (string, error) {
	limit := int64(1000000)

	p.links = map[int64]int64{}
	p.chainLen = map[int64]int64{}

	for i := int64(2); i <= limit; i++ {
		t := i
		l := int64(1)
		for {
			l++
			p.links[t] = p.nextLink(t)
			t = p.links[t]
			if t == 1 {
				break
			}
			if p.chainLen[t] > 0 {
				l += p.chainLen[t] - 1
				break
			}
		}
		p.chainLen[i] = l
	}

	maxLen := int64(0)
	maxVal := int64(0)
	for v, l := range p.chainLen {
		if l > maxLen {
			maxLen = l
			maxVal = v
		}
	}
	return fmt.Sprintf("%d", maxVal), nil
}

func (p *LongestCollatzSequence) nextLink(n int64) int64 {
	if n%2 == 0 {
		// n → n/2 (n is even)
		return n / 2
	}

	// n → 3n + 1 (n is odd)
	return 3*n + 1
}
