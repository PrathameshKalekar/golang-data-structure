/*
2239. Given an integer array nums of size n, return the number with the value closest to 0 in nums. If there are multiple answers, return the number with the largest value.

Example 1:

Input: nums = [-4,-2,1,4,8]
Output: 1
Explanation:
The distance from -4 to 0 is |-4| = 4.
The distance from -2 to 0 is |-2| = 2.
The distance from 1 to 0 is |1| = 1.
The distance from 4 to 0 is |4| = 4.
The distance from 8 to 0 is |8| = 8.
Thus, the closest number to 0 in the array is 1.
Example 2:

Input: nums = [2,-1,1]
Output: 1
Explanation: 1 and -1 are both the closest numbers to 0, so 1 being larger is returned.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-4, -2, 1, 4, 8}
	answer := findClosestNumberToZero(arr)
	fmt.Printf("Closest value to zero is : %d", answer)
}

func findClosestNumberToZero(arr []int) int {
	//Storing first variable in array
	temp := arr[0]

	//lopping each element
	for _, value := range arr {
		//if absolute value is less that temp variable then temp is that less value
		if math.Abs(float64(value)) < math.Abs(float64(temp)) {
			temp = value
		} else if math.Abs(float64(value)) == math.Abs(float64(temp)) && value > temp {
			//in case of -1 and 1 the absolute is same then it goes for this case that vlaue > temp means the actual value -1 and 1 the 1 is the closest
			temp = value
		}
	}
	//returning the closest value
	return temp
}
