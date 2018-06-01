package problems

import (
	"fmt"
	"strings"
)

type NumberLetterCounts struct {
	digit     map[int]string
	tens      map[int]string
	groupings []string
}

func (p *NumberLetterCounts) ID() int {
	return 17
}

func (p *NumberLetterCounts) Text() string {
	return `If the numbers 1 to 5 are written out in words: one, two,
three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters
used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written
out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
forty-two) contains 23 letters and 115 (one hundred and fifteen) contains
20 letters. The use of "and" when writing out numbers is in compliance
with British usage.
`
}

func (p *NumberLetterCounts) Solve() (string, error) {
	p.load()
	total := 0
	for i := 1; i <= 1000; i++ {
		s := p.numberToName(i)
		s = strings.Replace(s, " ", "", -1)
		s = strings.Replace(s, "-", "", -1)
		total += len(s)
	}
	return fmt.Sprintf("%d", total), nil
}

func (p *NumberLetterCounts) load() {
	p.digit = map[int]string{
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
	p.tens = map[int]string{
		2: "twenty",
		3: "thirty",
		4: "forty",
		5: "fifty",
		6: "sixty",
		7: "seventy",
		8: "eighty",
		9: "ninety",
	}
	p.groupings = []string{
		"",
		"thousand",
		"million",
		"billion",
		"trillion",
	}
}

func (p *NumberLetterCounts) numberToName(n int) string {
	if n == 0 {
		return "zero"
	}
	s := ""
	groupNum := 0
	for {
		group := n % 1000
		n = n / 1000
		groupText := p.numberGrouping(group)
		groupName := p.groupings[groupNum]
		if groupNum > 0 {
			groupText += " " + groupName
		}
		if len(groupText) > 0 && len(s) > 0 {
			s = groupText + " " + s
		} else {
			s = groupText
		}
		if n == 0 {
			break
		}
		groupNum++
	}

	return s
}

func (p *NumberLetterCounts) numberGrouping(n int) string {
	if n >= 1000 {
		panic(fmt.Errorf("number grouping was longer than 3 digits"))
	}

	var h string
	hundreds := (n / 100) % 10
	if hundreds > 0 {
		h = p.digit[hundreds] + " hundred"
	}

	var l string
	last2 := n % 100
	if last2 < 20 {
		l = p.digit[last2]
	} else {
		l = p.tens[last2/10]
		if last2%10 != 0 {
			l += "-" + p.digit[last2%10]
		}
	}

	var s string
	if len(h) > 0 && len(l) > 0 {
		s = h + " and " + l
	} else {
		s = h + l
	}
	return s
}
