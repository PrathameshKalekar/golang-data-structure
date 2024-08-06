/*
67 Given two binary strings a and b, return their sum as a binary string.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/
package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry != 0 {
		var bitA, bitB int

		if i >= 0 {
			bitA = int(a[i] - '0')
			i--
		}

		if j >= 0 {
			bitB = int(b[j] - '0')
			j--
		}

		sum := bitA + bitB + carry
		result = string(sum%2+'0') + result
		carry = sum / 2
	}

	return result
}

func main() {
	a := "11"
	b := "1"
	fmt.Println("Input: a =", a, ", b =", b)
	fmt.Println("Output:", addBinary(a, b))

	a = "1010"
	b = "1011"
	fmt.Println("Input: a =", a, ", b =", b)
	fmt.Println("Output:", addBinary(a, b))
}
