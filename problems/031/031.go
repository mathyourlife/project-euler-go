/*
Coin sums

In England the currency is made up of pound, £, and pence, p, and
there are eight coins in general circulation:

    1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).

It is possible to make £2 in the following way:

    1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

How many different ways can £2 be made using any number of coins?
*/

package main

import (
	"fmt"
)

func check(count int, a int,b int,c int,d int,e int,f int,g int,h int) (int, int) {
	total := a*1+b*2+c*5+d*10+e*20+f*50+g*100+h*200
	if total == 200 {
		count++
	}
	return total, count
}

func main() {

	var total, count int
	for a := 0; a <= 200; a++ {
		total, count = check(count,a,0,0,0,0,0,0,0)
		if total >= 200 { break }
	for b:= 0; b <= 200/2; b++ {
		total, count = check(count,a,b,0,0,0,0,0,0)
		if total >= 200 { break }
	for c:= 0; c <= 200/5; c++ {
		total, count = check(count,a,b,c,0,0,0,0,0)
		if total >= 200 { break }
	for d:= 0; d <= 200/10; d++ {
		total, count = check(count,a,b,c,d,0,0,0,0)
		if total >= 200 { break }
	for e:= 0; e <= 200/20; e++ {
		total, count = check(count,a,b,c,d,e,0,0,0)
		if total >= 200 { break }
	for f:= 0; f <= 200/50; f++ {
		total, count = check(count,a,b,c,d,e,f,0,0)
		if total >= 200 { break }
	for g:= 0; g <= 200/100; g++ {
		total, count = check(count,a,b,c,d,e,f,g,0)
		if total >= 200 { break }
	for h:= 0; h <= 200/200; h++ {
		total, count = check(count,a,b,c,d,e,f,g,h)
		if total >= 200 { break }
	}}}}}}}}
	fmt.Println(count)
}
