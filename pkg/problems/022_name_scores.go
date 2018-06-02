package problems

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type NameScores struct {
	namesFile string
}

func (p *NameScores) ID() int {
	return 22
}

func (p *NameScores) Text() string {
	return `Using names.txt, a 46K text file containing over
five-thousand first names, begin by sorting it into alphabetical order.
Then working out the alphabetical value for each name, multiply this
value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN,
which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the
list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
`
}

func (p *NameScores) Solve() (string, error) {
	file, err := os.Open(p.namesFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	content := string(b)
	quotedNames := strings.Split(content, ",")

	names := make([]string, 0, len(quotedNames))
	for _, quotedName := range quotedNames {
		names = append(names, quotedName[1:len(quotedName)-1])
	}
	sort.Strings(names)

	total := 0
	for i, name := range names {
		sum := 0
		for _, n := range []byte(name) {
			sum += int(n) - 64
		}
		total += (i + 1) * sum
	}

	return fmt.Sprintf("%d", total), nil
}
