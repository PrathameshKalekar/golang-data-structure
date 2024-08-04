/*
739 Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]
*/
package main

import "fmt"

func main() {
	arr := []int{71, 71, 76, 71, 76, 71, 71}
	answer := dailyTemperatures(arr)
	fmt.Println(answer)
}

func dailyTemperatures(temp []int) []int {
	n := len(temp)
	result := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temp[i] > temp[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

	return result
}

/*
func dailyTemperatures(temp []int) []int {
    // Initialize result array with the same length as temp, filled with zeros
    result := make([]int, len(temp))

    // Iterate through each temperature in the temp array
    for i := 0; i < len(temp); i++ {
        added := false // Flag to indicate if a warmer day was found
        index := 1 // Tracks the number of days until a warmer temperature

        // Iterate through the subsequent days to find a warmer temperature
        for j := i + 1; j < len(temp); j++ {
            if temp[i] < temp[j] { // Check if the current day is colder than the future day
                if result[i] == 0 { // If no warmer day has been recorded yet
                    added = true // Set the flag indicating a warmer day was found
                    result[i] = index // Record the number of days until a warmer day
                }
            }
            index++ // Increment the index to track the number of days passed
        }

        // If no warmer day was found, set the result for that day to 0
        if !added {
            result[i] = 0
        }
    }

    // Return the result array containing the number of days until a warmer temperature for each day
    return result
}
*/
