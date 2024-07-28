/*
56  Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.



Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	answer := mergeIntervals(intervals)
	fmt.Println(answer)
}

func mergeIntervals(intervals [][]int) [][]int {
	//if the intervals are empty return empty
	if len(intervals) < 2 {
		return intervals
	}
	//arranging all slices in sorted order
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	current := intervals[0]

	//lopping with each element in the intervals
	for i := 1; i < len(intervals); i++ {
		//if the i th intervasl 0 th value is less that currents 1st value then the  intervals is ovelapped
		if intervals[i][0] <= current[1] {
			//and if currents 2 nd value is smaller than ith intervals second value teh interchange the value
			if intervals[i][1] > current[1] {
				current[1] = intervals[i][1]
			}
		} else {
			//after chnagin the current interval append it to result
			result = append(result, current)
			current = intervals[i]
		}
	}
	//the for loop not handling the last element so add it personally
	result = append(result, current)
	//return result
	return result
}
