/*
349 Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.
*/
package main

import "fmt"

func main() {
	arr := []int{4, 9, 5}
	arr2 := []int{9, 4, 9, 8, 4}
	answer := intersectionOfTwoArrays(arr, arr2)
	fmt.Println(answer)
}

func intersectionOfTwoArrays(arr, arr2 []int) []int {
	//num map to keep track of al element
	numMap := map[int]bool{}
	//adding all element of first array in hash table to checck the intersection in other array
	for _, value := range arr {
		numMap[value] = true
	}
	//result of intersect numbers
	result := []int{}
	//keep track of result that no reapeated number added
	resultMap := map[int]bool{}
	//checking the intersected element in second array
	for i := 0; i < len(arr2); i++ {
		//if the num is present in nummap means intersected and not in result then add in result
		if numMap[arr2[i]] && !resultMap[arr2[i]] {
			result = append(result, arr2[i])
			resultMap[arr2[i]] = true
		}
	}
	//return the result
	return result
}
