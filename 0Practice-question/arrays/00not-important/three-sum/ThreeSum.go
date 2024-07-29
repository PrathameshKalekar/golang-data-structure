/*
15 Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	answer := threeSum(arr)
	fmt.Println(answer)
}

func threeSum(arr []int) [][]int {
	result := [][]int{}
	sort.Ints(arr)
	n := len(arr)
	//checking by each node
	for i := 0; i < n-2; i++ {
		//skip the point if its a duplicate
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		//left and right pointer
		left, right := i+1, n-1
		for left < right {
			//sum conitions
			sum := arr[i] + arr[left] + arr[right]
			if sum == 0 {
				result = append(result, []int{arr[i], arr[left], arr[right]})
				//if there is duplicate value after left then skip
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				//if there is a duplicate value after right then skip
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				//after done with the current left right modify them
				left++
				right--
				//else if the sum is less than 0 means the - value is more so shift left else right
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
