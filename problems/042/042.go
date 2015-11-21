/*
Coded triangle numbers

The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1);
so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its
alphabetical position and adding these values we form a word value.
For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the
word value is a triangle number then we shall call the word a triangle word.

Using words.txt, a 16K text file containing nearly two-thousand common
English words, how many are triangle words?
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Coded triangle numbers")
	fi, err := os.Open("p042_words.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	triangles := map[int]bool{}

	var t int
	for n := 1; n <= 20; n++ {
		t = n * (n + 1) / 2
		triangles[t] = true
	}

	// make a buffer to keep chunks that are read
	buf := make([]byte, 128)
	word_sum := 0
	count := 0
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		for _, c := range buf[:n] {
			// Words are quote wrapped and comma separated
			// If we hit a comma, reset the word sum
			if c == 34 {
				// Double quote, either the start or end of a word
				if triangles[word_sum] {
					count++
				}
				if word_sum > t {
					panic(fmt.Sprintf("Precomputed triangle numbers do not cover up to %d", word_sum))
				}
				word_sum = 0
			} else if c == 44 {
				// Do nothing with commas
				continue
			} else {
				word_sum += int(c - 64)
			}
		}
	}
	fmt.Println(count)
}
