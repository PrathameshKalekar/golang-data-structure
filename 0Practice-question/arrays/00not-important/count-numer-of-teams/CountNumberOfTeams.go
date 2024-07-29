/*
There are n soldiers standing in a line. Each soldier is assigned a unique rating value.

You have to form a team of 3 soldiers amongst them under the following rules:

Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k]).
A team is valid if: (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).
Return the number of teams you can form given the conditions. (soldiers can be part of multiple teams).

Example 1:

Input: rating = [2,5,3,4,1]
Output: 3
Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (5,3,1).
Example 2:

Input: rating = [2,1,3]
Output: 0
Explanation: We can't form any team given the conditions.
Example 3:

Input: rating = [1,2,3,4]
Output: 4
*/
package main

import "fmt"

func main() {
	rating := []int{
		1, 2, 3, 4,
	}
	answer := countNumberOfTeams(rating)
	fmt.Printf("Total unique rating count is : %d", answer)
}

func countNumberOfTeams(rating []int) int {
	n := len(rating)
	count := 0

	//looping j cause we have to find left high or low and righ thigh or low
	for j := 0; j < n; j++ {
		leftLow, leftHigh, rightLow, rightHigh := 0, 0, 0, 0
		//for finding any the left value is graeter or smaller
		for i := 0; i < j; i++ {
			if rating[i] < rating[j] {
				leftLow++
			} else if rating[i] > rating[j] {
				leftHigh++
			}
		}
		//finding any right value is greater or smaller
		for k := j + 1; k < n; k++ {
			if rating[j] < rating[k] {
				rightHigh++
			} else if rating[j] > rating[k] {
				rightLow++
			}
		}
		//formula for count if any of this is not match the it add 0 means the pair is not perfectly created
		count += leftHigh*rightLow + leftLow*rightHigh
	}
	//return count if no count is there it will return zero by default
	return count
}
