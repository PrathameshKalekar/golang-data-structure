/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 4, 3, 543, 544, 545, 546, 8, 6, 9, 0}
	answer := longestConsecutiveNumber(arr)
	fmt.Println(answer)
}

func longestConsecutiveNumber(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	count, maxCount := 0, 0
	//sort the array
	sort.Ints(arr)
	//looping through each element in sorted array
	for i := 1; i < len(arr); i++ {
		//if current element is equal to previous + 1 the its a consecutive else both same the also its consecutive but no increament in count
		if arr[i] == arr[i-1]+1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else if arr[i] == arr[i-1] {
			continue
		} else {
			//both condition not satisfies means the consecutive sequence broken so set count to 0
			count = 0
		}

	}
	return maxCount + 1
}
