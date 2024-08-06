/*
43 Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "498828660196"
	s2 := "840477629533"
	answer := multiplyStrings(s1, s2)
	fmt.Println(answer)
}
func multiplyStrings(num1, num2 string) string {
	if num1 == "0" || num2 == "0" { // Edge case: if either number is "0"
		return "0"
	}
	m, n := len(num1), len(num2) // Lengths of the input strings
	result := make([]int, m+n)   // Result array to hold the intermediate results

	// Multiply each digit of num1 with each digit of num2
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0') // Multiply digits
			p1, p2 := i+j, i+j+1                       // Positions in the result array
			sum := mul + result[p2]                    // Add the current multiplication to the existing value in result

			result[p1] += sum / 10 // Carry over to the next position
			result[p2] = sum % 10  // Store the current digit
		}
	}

	// Convert the result array to a string, skipping leading zeros
	sb := strings.Builder{}
	for i := 0; i < len(result); i++ {
		if !(sb.Len() == 0 && result[i] == 0) {
			sb.WriteByte(byte(result[i]) + '0')
		}
	}

	return sb.String()
}
