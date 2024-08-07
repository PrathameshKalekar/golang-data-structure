/*
273 Convert a non-negative integer num to its English words representation.

Example 1:

Input: num = 123
Output: "One Hundred Twenty Three"
Example 2:

Input: num = 12345
Output: "Twelve Thousand Three Hundred Forty Five"
Example 3:

Input: num = 1234567
Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
*/
package main

import (
	"fmt"
	"strings"
)

var belowTwenty = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var thousands = []string{"", "Thousand", "Million", "Billion"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	res := ""
	for i := 0; num > 0; i++ {
		if num%1000 != 0 {
			res = helper(num%1000) + thousands[i] + " " + res
		}
		num /= 1000
	}
	return strings.TrimSpace(res)
}

func helper(num int) string {
	if num == 0 {
		return ""
	} else if num < 20 {
		return belowTwenty[num] + " "
	} else if num < 100 {
		return tens[num/10] + " " + helper(num%10)
	} else {
		return belowTwenty[num/100] + " Hundred " + helper(num%100)
	}
}

func main() {
	fmt.Println(numberToWords(123))     // Output: "One Hundred Twenty Three"
	fmt.Println(numberToWords(12345))   // Output: "Twelve Thousand Three Hundred Forty Five"
	fmt.Println(numberToWords(1234567)) // Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
}
