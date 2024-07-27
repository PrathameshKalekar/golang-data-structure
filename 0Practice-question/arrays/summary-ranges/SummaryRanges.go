/*
228 You are given a sorted unique integer array nums.

A range [a,b] is the set of all integers from a to b (inclusive).

Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b

Example 1:

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
Example 2:

Input: nums = [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{0, 1, 2, 4, 5, 8}
	answer := summaryRanges(arr)
	fmt.Println(answer)
}

func summaryRanges(arr []int) []string {
	//if there is no element in array return nil
	if len(arr) == 0 {
		return nil
	}
	var result []string
	//first element as start element
	start := arr[0]
	//iterating thorugh all element in array except first cause its the start element
	for i := 1; i < len(arr); i++ {
		//if current element and previous element not diffrence by 1 then the contigous is broken
		if arr[i] != arr[i-1]+1 {
			//if the previous element and start point is same form breaking point then no middle point present only add the point
			if start == arr[i-1] {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, arr[i-1]))
			}
			//start from new starting point after element added in result
			start = arr[i]
		}
	}

	//the last element is not iterated in upper for loop so do it personally
	if start == arr[len(arr)-1] {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, arr[len(arr)-1]))
	}
	//return the result
	return result
}
