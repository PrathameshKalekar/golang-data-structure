/*
4  Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).



Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1, 3}
	arr2 := []int{2}
	answer := medianOfTwoSortedArray(arr1, arr2)
	fmt.Printf("Both arrays median is : %f", answer)
}

func medianOfTwoSortedArray(arr1, arr2 []int) float64 {
	if len(arr1) == 0 && len(arr2) == 0 {
		return 0
	}
	var temp []int
	//append both array in temparary array
	temp = append(arr1, arr2...)
	//sort the array
	sort.Ints(temp)
	//formula if  the left of array is even or orr on that basis it will calculate the median and return it
	if len(temp)%2 == 0 {
		answer := float64(temp[(len(temp)/2)-1]+temp[len(temp)/2]) / 2
		return answer
	} else {
		answer := float64(temp[len(temp)/2])
		return answer
	}
}
