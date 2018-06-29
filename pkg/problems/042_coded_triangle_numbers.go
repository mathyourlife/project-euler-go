package problems

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type CodedTriangleNumbers struct {
	wordsFile string
}

func (p *CodedTriangleNumbers) ID() int {
	return 42
}

func (p *CodedTriangleNumbers) Text() string {
	return `The nth term of the sequence of triangle numbers
is given by, tn = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its
alphabetical position and adding these values we form a word value.
For example, the word value for SKY is 19 + 11 + 25 = 55 = t_10. If the
word value is a triangle number then we shall call the word a
triangle word.

Using words.txt, a 16K text file containing nearly two-thousand common
English words, how many are triangle words?
`
}

func (p *CodedTriangleNumbers) Solve() (string, error) {

	triangleNumbers := map[int]bool{}
	maxTriangle := 0

	// closure over the cache of triangle numbers map
	// should probably switch over to N(N+1)/2 formula (meh)
	isTriangle := func(n int) bool {
		for {
			if n <= maxTriangle {
				break
			}
			maxTriangle += len(triangleNumbers) + 1
			triangleNumbers[maxTriangle] = true
		}
		return triangleNumbers[n]
	}

	file, err := os.Open(p.wordsFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	content := string(b)
	quotedWords := strings.Split(content, ",")

	tally := 0
	for _, quotedWord := range quotedWords {
		sum := 0
		for i := 1; i < len(quotedWord)-1; i++ {
			sum += int(quotedWord[i]) - 64
		}
		if isTriangle(sum) {
			tally++
		}
	}

	return fmt.Sprintf("%d", tally), nil
}
