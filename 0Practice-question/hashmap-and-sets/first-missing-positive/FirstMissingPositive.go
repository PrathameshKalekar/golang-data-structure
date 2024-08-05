/*
41 Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.

You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.



Example 1:

Input: nums = [1,2,0]
Output: 3
Explanation: The numbers in the range [1,2] are all in the array.
Example 2:

Input: nums = [3,4,-1,1]
Output: 2
Explanation: 1 is in the array but 2 is missing.
Example 3:

Input: nums = [7,8,9,11,12]
Output: 1
Explanation: The smallest positive integer 1 is missing.
*/

package main

import "fmt"

func main() {
	arr := []int{3, 4, -1, 1}
	answer := firstMissingPositive(arr)
	fmt.Println("First Missing positive number in array is : ", answer)
}

func firstMissingPositive(arr []int) int {
	// for mapping each integer in arr
	numMap := map[int]bool{}
	//add all positive number in numMap and make then true
	for _, value := range arr {
		if value > 0 {
			//if the num is greater that 0 then only add it in the map
			numMap[value] = true
		}
	}
	//default pointer that check the missing value
	missingPositve := 1
	//infinite loop till the missing positve is not got
	for {
		//if missingpositive number in not in mapp then its the missing positive and return it
		if !numMap[missingPositve] {
			return missingPositve
		} else {
			//else increment the pointer
			missingPositve++
		}
	}

}
