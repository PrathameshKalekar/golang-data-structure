/*
2134 A swap is defined as taking two distinct positions in an array and swapping the values in them.

A circular array is defined as an array where we consider the first element and the last element to be adjacent.

Given a binary circular array nums, return the minimum number of swaps required to group all 1's present in the array together at any location.

Example 1:

Input: nums = [0,1,0,1,1,0,0]
Output: 1
Explanation: Here are a few of the ways to group all the 1's together:
[0,0,1,1,1,0,0] using 1 swap.
[0,1,1,1,0,0,0] using 1 swap.
[1,1,0,0,0,0,1] using 2 swaps (using the circular property of the array).
There is no way to group all 1's together with 0 swaps.
Thus, the minimum number of swaps required is 1.
Example 2:

Input: nums = [0,1,1,1,0,0,1,1,0]
Output: 2
Explanation: Here are a few of the ways to group all the 1's together:
[1,1,1,0,0,0,0,1,1] using 2 swaps (using the circular property of the array).
[1,1,1,1,1,0,0,0,0] using 2 swaps.
There is no way to group all 1's together with 0 or 1 swaps.
Thus, the minimum number of swaps required is 2.
Example 3:

Input: nums = [1,1,0,0,1]
Output: 0
Explanation: All the 1's are already grouped together due to the circular property of the array.
Thus, the minimum number of swaps required is 0.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{0, 1, 0, 1, 1, 0, 0}
	answer := munimumSwapToGroupAll1sToghether2(arr)
	fmt.Println(answer)
}
func munimumSwapToGroupAll1sToghether2(arr []int) int {
	n := len(arr)
	totalOnes := 0
	//counting all 1s
	for _, value := range arr {
		if value == 1 {
			totalOnes++
		}
	}
	//no one or all are one then return 0
	if totalOnes == 0 || totalOnes == n {
		return 0
	}
	//creating a array to handle circular nature
	extended := append(arr, arr...)
	minZeros := math.MaxInt32
	currentZero := 0
	for i := 0; i < totalOnes; i++ {
		if extended[i] == 0 {
			currentZero++
		}
	}
	minZeros = currentZero

	//slide window over the extended array
	for i := totalOnes; i < len(extended); i++ {
		if extended[i-totalOnes] == 0 {
			currentZero--
		}
		if extended[i] == 0 {
			currentZero++
		}
		if currentZero < minZeros {
			minZeros = currentZero
		}
	}
	return minZeros
}
