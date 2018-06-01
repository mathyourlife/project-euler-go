package problems

import (
	"fmt"
	"time"
)

type CountingSundays struct{}

func (p *CountingSundays) ID() int {
	return 19
}

func (p *CountingSundays) Text() string {
	return `You are given the following information, but you may prefer to do
some research for yourself.

* 1 Jan 1900 was a Monday.
* Thirty days has September,
	April, June and November.
	All the rest have thirty-one,
	Saving February alone,
	Which has twenty-eight, rain or shine.
	And on leap years, twenty-nine.
 * A leap year occurs on any year evenly divisible by 4, but not on
	a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth
century (1 Jan 1901 to 31 Dec 2000)?
`
}

func (p *CountingSundays) Solve() (string, error) {
	// taking the easy way out and just using built-in
	d := time.Date(1901, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC)

	sundays := 0
	for {
		if d.Weekday() == time.Sunday {
			sundays++
		}
		d = d.AddDate(0, 1, 0)
		if d.After(end) {
			break
		}
	}
	return fmt.Sprintf("%d", sundays), nil
}
