/*
415 Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.

Example 1:

Input: num1 = "11", num2 = "123"
Output: "134"
Example 2:

Input: num1 = "456", num2 = "77"
Output: "533"
Example 3:

Input: num1 = "0", num2 = "0"
Output: "0"
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "456"
	num2 := "77"
	answer := addStrings(num1, num2)
	fmt.Println(answer)
}

func addStrings(num1, num2 string) string {
	if len(num1) < len(num2) {
		return addStrings(num2, num1)
	}
	result := ""
	carry := 0
	res := 0
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 {

		if j < 0 {
			res = int(num1[i]-'0') + carry
		} else {
			res = int(num1[i]-'0') + int(num2[j]-'0') + carry
		}
		remainder := res % 10
		carry = res / 10
		result = strconv.Itoa(remainder) + result
		i--
		j--
	}
	if carry > 0 {
		result = strconv.Itoa(res+carry) + result
	}
	return result
}
