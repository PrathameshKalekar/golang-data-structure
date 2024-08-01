/*
169 Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2
*/
package main

import "fmt"

func main() {
	arr := []int{2, 1, 1, 1, 2}
	answer := majorityElement(arr)
	fmt.Println(answer)
}

func majorityElement(arr []int) int {
	element := arr[0]
	elementMap := make(map[int]int)

	for _, value := range arr {
		elementMap[value]++
	}
	for ele, count := range elementMap {
		if elementMap[element] < count {
			element = ele
		}
	}
	return element
}
