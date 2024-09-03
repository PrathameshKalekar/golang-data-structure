/*
1071 For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""
*/
package main

import "fmt"

func main() {
	str1 := "ABABAB"
	str2 := "ABAB"
	answer := greatestCommonDivisorOfStrings(str1, str2)
	fmt.Println(answer)
}

func greatestCommonDivisorOfStrings(str1, str2 string) string {
	// Check if str1 + str2 is the same as str2 + str1
	if str1+str2 != str2+str1 {
		return ""
	}
	// Find the GCD of the lengths of str1 and str2
	lengthGCD := gcd(len(str1), len(str2))
	// Return the substring of str1 from the start to lengthGCD
	return str1[:lengthGCD]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
