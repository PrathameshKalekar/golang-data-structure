/*
14 Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/

package main

import "fmt"

func main() {
	strings := []string{"flower", "flow", "flight"}
	answer := longestCommonPrefix(strings)
	fmt.Printf("the longest common prefix in array is : %s", answer)
}

func longestCommonPrefix(strings []string) string {
	//if no string are present in the array return empty string
	if len(strings) < 1 {
		return ""
	}
	//set prefix as first string
	prefix := strings[0]
	//iterating each string except first cause its the prefix
	for _, str := range strings[1:] {
		i := 0
		//increment i when letter are same in both strings
		for i < len(prefix) && i < len(str) && prefix[i] == str[i] {
			i++
		}
		//the prefix till the i is added cause its the condition
		prefix = prefix[:i]
	}
	//return prefix
	return prefix
}
