/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains
10 terms. Although it has not been proved yet (Collatz Problem), it is
thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import (
	"fmt"
)

func next_collatz(n int64) int64 {
	if n%2 == 0 {
		// n → n/2 (n is even)
		return n / 2
	} else {
		// n → 3n + 1 (n is odd)
		return 3*n + 1
	}
}

func collatz_sequence(n int64, lengths map[int64]int64) {
	seq := []int64{}
	for {
		seq = append(seq, n)
		// fmt.Println(n, lengths[n], seq)
		if n == 1 {
			break
		}
		if lengths[n] != 0 {
			// fmt.Println("already know chain", n, "is", lengths[n])
			break
		}
		n = next_collatz(n)
	}
	record_lengths(seq, lengths)
}

func record_lengths(seq []int64, lengths map[int64]int64) {

	l := int64(len(seq))
	tail_len := int64(1)
	if seq[l-1] != 1 {
		tail_len = lengths[seq[l-int64(1)]]
	}
	// fmt.Println("Using a tail length of", tail_len)
	for i := l - 1; i >= 0; i-- {
		lengths[seq[i]] = l - i + tail_len - 1
	}
}

func main() {
	lengths := map[int64]int64{}

	for i := int64(1); i <= 1000000; i++ {
		collatz_sequence(i, lengths)
	}

	max_n := int64(0)
	max_count := int64(0)
	for n, c := range lengths {
		if c > max_count {
			max_count = c
			max_n = n
		}
	}
	fmt.Printf("%d takes %d steps\n", max_n, max_count)

}
