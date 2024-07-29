/*
18 Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.



Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 0, -1, 0, -2, 2}
	target := 0
	answer := fourSum(arr, target)
	fmt.Println(answer)
}

func fourSum(arr []int, target int) [][]int {
	sort.Ints(arr)
	var result [][]int
	//lopping and skipping the last two cause this place will take the pointer i and j
	for i := 0; i < len(arr)-3; i++ {
		// if same then continue means skip the line
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		// second loop that will handle the second array pointer j
		for j := i + 1; j < len(arr)-2; j++ {
			//skip if current and previous is same
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			//making two more popinter to handle the rest two points
			left, right := j+1, len(arr)-1
			for left < right {

				sum := arr[i] + arr[j] + arr[left] + arr[right]
				//if sum == 0 append the points in the array else handle the left right pointer
				if sum == target {
					result = append(result, []int{arr[i], arr[j], arr[left], arr[right]})
					for left < right && arr[left] == arr[left+1] {
						left++
					}
					for left < right && arr[right] == arr[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	//return result
	return result
}
