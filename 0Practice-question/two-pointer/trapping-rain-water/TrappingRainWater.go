/*
42 Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

Example 1:

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
Example 2:

Input: height = [4,2,0,3,2,5]
Output: 9
*/
package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trappingRainWater(height))
}
func trappingRainWater(heights []int) int {
	left, right := 0, len(heights)-1
	leftMax, rightMax, totalWater := 0, 0, 0
	for left <= right {
		//if left is less than right the change leftmax and totalwater
		if heights[left] <= heights[right] {
			if heights[left] >= leftMax {
				leftMax = heights[left]
			} else {
				totalWater += leftMax - heights[left]
			}
			left++
		} else {
			//else change rightMax and totalwater
			if heights[right] >= rightMax {
				rightMax = heights[right]
			} else {
				totalWater += rightMax - heights[right]
			}
			right--
		}
	}
	//return totalwater
	return totalWater
}
