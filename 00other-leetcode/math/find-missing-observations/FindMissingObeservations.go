/*
2028 You have observations of n + m 6-sided dice rolls with each face numbered from 1 to 6. n of the observations went missing, and you only have the observations of m rolls. Fortunately, you have also calculated the average value of the n + m rolls.

You are given an integer array rolls of length m where rolls[i] is the value of the ith observation. You are also given the two integers mean and n.

Return an array of length n containing the missing observations such that the average value of the n + m rolls is exactly mean. If there are multiple valid answers, return any of them. If no such array exists, return an empty array.

The average value of a set of k numbers is the sum of the numbers divided by k.

Note that mean is an integer, so the sum of the n + m rolls should be divisible by n + m.

Example 1:

Input: rolls = [3,2,4,3], mean = 4, n = 2
Output: [6,6]
Explanation: The mean of all n + m rolls is (3 + 2 + 4 + 3 + 6 + 6) / 6 = 4.
Example 2:

Input: rolls = [1,5,6], mean = 3, n = 4
Output: [2,3,2,2]
Explanation: The mean of all n + m rolls is (1 + 5 + 6 + 2 + 3 + 2 + 2) / 7 = 3.
Example 3:

Input: rolls = [1,2,3,4], mean = 6, n = 4
Output: []
Explanation: It is impossible for the mean to be 6 no matter what the 4 missing rolls are.
*/
package main

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	// Calculate the total sum needed for all n + m rolls
	totalSum := mean * (n + len(rolls))

	// Calculate the sum of the known rolls
	knownSum := 0
	for _, roll := range rolls {
		knownSum += roll
	}

	// Calculate the missing sum
	missingSum := totalSum - knownSum

	// Check if it's possible to find the missing rolls
	if missingSum < n || missingSum > 6*n {
		return []int{}
	}

	// Create the result array and distribute the missing sum
	result := make([]int, n)
	avg := missingSum / n
	remainder := missingSum % n

	// Distribute the average and the remainder
	for i := 0; i < n; i++ {
		result[i] = avg
		if i < remainder {
			result[i]++
		}
	}

	return result
}

func main() {
	fmt.Println(missingRolls([]int{3, 2, 4, 3}, 4, 2)) // Output: [6, 6]
	fmt.Println(missingRolls([]int{1, 5, 6}, 3, 4))    // Output: [2, 3, 2, 2]
	fmt.Println(missingRolls([]int{1, 2, 3, 4}, 6, 4)) // Output: []
}
