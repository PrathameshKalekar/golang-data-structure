/*
50 Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

Example 1:

Input: x = 2.00000, n = 10
Output: 1024.00000
Example 2:

Input: x = 2.10000, n = 3
Output: 9.26100
Example 3:

Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
*/
package main

import "fmt"

// StoneGameII computes the maximum number of stones Alice can get
func StoneGameII(piles []int) int {
	n := len(piles)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	suffixSum := make([]int, n)
	suffixSum[n-1] = piles[n-1]

	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + piles[i]
	}

	for i := n - 1; i >= 0; i-- {
		for m := 1; m <= n; m++ {
			if i+2*m >= n {
				dp[i][m] = suffixSum[i]
			} else {
				for x := 1; x <= 2*m; x++ {
					dp[i][m] = max(dp[i][m], suffixSum[i]-dp[i+x][max(m, x)])
				}
			}
		}
	}

	return dp[0][1]
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	piles := []int{2, 7, 9, 4, 4}
	fmt.Println(StoneGameII(piles)) // Output: 10
}
