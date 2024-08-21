/*
664 There is a strange printer with the following two special properties:

The printer can only print a sequence of the same character each time.
At each turn, the printer can print new characters starting from and ending at any place and will cover the original existing characters.
Given a string s, return the minimum number of turns the printer needed to print it.

Example 1:

Input: s = "aaabbb"
Output: 2
Explanation: Print "aaa" first and then print "bbb".
Example 2:

Input: s = "aba"
Output: 2
Explanation: Print "aaa" first and then print "b" from the second place of the string, which will cover the existing character 'a'.
*/
package main

import (
	"fmt"
)

func strangePrinter(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Initialize dp array
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Base case: single character substrings
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	// Fill the dp array
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			dp[i][j] = dp[i][j-1] + 1
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j-1])
				} else {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}

	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(strangePrinter("aaabbb")) // Output: 2
	fmt.Println(strangePrinter("aba"))    // Output: 2
}
