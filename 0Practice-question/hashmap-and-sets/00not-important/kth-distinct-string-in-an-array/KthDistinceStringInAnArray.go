/*
2053 A distinct string is a string that is present only once in an array.

Given an array of strings arr, and an integer k, return the kth distinct string present in arr. If there are fewer than k distinct strings, return an empty string "".

Note that the strings are considered in the order in which they appear in the array.

Example 1:

Input: arr = ["d","b","c","b","c","a"], k = 2
Output: "a"
Explanation:
The only distinct strings in arr are "d" and "a".
"d" appears 1st, so it is the 1st distinct string.
"a" appears 2nd, so it is the 2nd distinct string.
Since k == 2, "a" is returned.
Example 2:

Input: arr = ["aaa","aa","a"], k = 1
Output: "aaa"
Explanation:
All strings in arr are distinct, so the 1st string "aaa" is returned.
Example 3:

Input: arr = ["a","b","a"], k = 3
Output: ""
Explanation:
The only distinct string is "b". Since there are fewer than 3 distinct strings, we return an empty string "".
*/
package main

import "fmt"

func main() {
	arr := []string{"d", "b", "c", "b", "c", "a"}
	answer := kthDistinceStringInAnArray(arr, 2)
	fmt.Println(answer)
}

func kthDistinceStringInAnArray(arr []string, k int) string {
	// Create a map to count the occurrences of each string
	arrMap := map[string]int{}
	for _, value := range arr {
		arrMap[value]++
	}

	// Create a slice to store distinct strings
	result := []string{}
	// Iterate over the array again to collect strings that occur exactly once
	for _, value := range arr {
		if arrMap[value] == 1 {
			result = append(result, value)
		}
	}

	// Check if there are fewer than k distinct strings
	if len(result) < k {
		return ""
	}
	// Return the k-th distinct string
	return result[k-1]
}
