/*
11 You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.



Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1
*/

package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	answer := containerWithMostWater(height)
	fmt.Printf("Max area filled with water is : %d", answer)
}

func containerWithMostWater(height []int) int {
	//making two point that handle both side value
	left, right, answer := 0, len(height)-1, 0
	//loop till all element are scanned
	for left < right {
		//water can only hold till low height so find min
		h := min(height[left], height[right])
		//width is distanse between right and left
		width := right - left
		//formula for area
		currentAnswer := h * width
		//if area is grater tha answer change the max area to answer
		if currentAnswer > answer {
			answer = currentAnswer
		}
		//if left height is small shift left side else shift right side
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	//return answer
	return answer
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
