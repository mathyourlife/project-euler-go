/*
Double-base palindromes

The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in
base 10 and base 2.

(Please note that the palindromic number, in either base, may not include
leading zeros.)
*/

package main

import(
  "fmt"
  "strconv"
)

func is_palindrome(n string) bool {
	for i := 0; i < (len(n)+1)/2; i++ {
		if n[i:i+1] != n[len(n)-i-1:len(n)-i] {
			return false
		}
	}
	return true
}

func main() {
	pals_sum := int64(0)
	for n := int64(0); n < 1000000; n++ {
		if is_palindrome(fmt.Sprintf("%d", n)) &&
			is_palindrome(
				fmt.Sprintf("%s", strconv.FormatInt(n, 2))) {
			pals_sum += n
			// fmt.Println(n)
			fmt.Println(strconv.FormatInt(n, 2))
		}
	}
	fmt.Println(pals_sum)

}