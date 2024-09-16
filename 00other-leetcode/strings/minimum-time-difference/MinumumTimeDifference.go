/*
539 Given a list of 24-hour clock time points in "HH:MM" format, return the minimum minutes difference between any two time-points in the list.

Example 1:

Input: timePoints = ["23:59","00:00"]
Output: 1
Example 2:

Input: timePoints = ["00:00","23:59","00:00"]
Output: 0
*/
package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	// Helper function to convert "HH:MM" to minutes
	convertToMinutes := func(time string) int {
		parts := strings.Split(time, ":")
		hours, _ := strconv.Atoi(parts[0])
		minutes, _ := strconv.Atoi(parts[1])
		return hours*60 + minutes
	}

	// Step 1: Convert time points to minutes
	minutesList := make([]int, len(timePoints))
	for i, time := range timePoints {
		minutesList[i] = convertToMinutes(time)
	}

	// Step 2: Sort the minutes list
	sort.Ints(minutesList)

	// Step 3: Compute the minimum difference between consecutive times
	minDiff := math.MaxInt32
	for i := 1; i < len(minutesList); i++ {
		diff := minutesList[i] - minutesList[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	// Step 4: Compute the circular difference (between the last and first time)
	firstAndLastDiff := 1440 - minutesList[len(minutesList)-1] + minutesList[0]
	if firstAndLastDiff < minDiff {
		minDiff = firstAndLastDiff
	}

	return minDiff
}

func main() {
	// Example 1
	timePoints1 := []string{"23:59", "00:00"}
	fmt.Println(findMinDifference(timePoints1)) // Output: 1

	// Example 2
	timePoints2 := []string{"00:00", "23:59", "00:00"}
	fmt.Println(findMinDifference(timePoints2)) // Output: 0
}
