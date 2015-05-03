/*
You are given the following information, but you may prefer to do
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

*/

package main

import (
	"fmt"
	"strings"
)

type Date struct {
	Year    int
	Month   int
	Day     int
	Weekday int
}

func NewDate(year int, month int, day int) *Date {
	return &Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func (d *Date) Print() string {
	weekdays := map[int]string{
		0: "Monday",
		1: "Tuesday",
		2: "Wednesday",
		3: "Thursday",
		4: "Friday",
		5: "Saturday",
		6: "Sunday",
	}
	ds := fmt.Sprintf("%4d-%2d-%2d", d.Year, d.Month, d.Day)
	ds = strings.Replace(ds, " ", "0", -1)
	return fmt.Sprintf("%s %s", ds, weekdays[d.Weekday])
}

func (d *Date) NextMonth() {

	length := 31
	for _, m := range []int{4, 6, 9, 11} {
		if m == d.Month {
			length = 30
			break
		}
	}
	if d.Month == 2 && (d.Year%4 == 0) {
		length = 29
	} else if d.Month == 2 {
		length = 28
	}

	d.Weekday = (d.Weekday + length) % 7
	d.Month++
	if d.Month > 12 {
		d.Year++
		d.Month = d.Month % 12
	}
}

func main() {

	// d := &Date{
	// 	Year: 1901,
	// 	Month: 1,
	// 	Day: 1,
	// 	Weekday: "T",
	// }
	d := &Date{
		Year:    1901,
		Month:   1,
		Day:     1,
		Weekday: 1,
	}

	sundays := 0
	for {
		if d.Weekday == 6 {
			sundays++
		}
		d.NextMonth()
		if d.Year == 2001 && d.Month == 1 {
			break
		}
	}
	fmt.Println(sundays)
}
