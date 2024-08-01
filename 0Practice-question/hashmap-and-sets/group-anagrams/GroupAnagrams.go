/*
49 nGiven an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	answer := groupAnagrams(arr)
	fmt.Println(answer)
}

func groupAnagrams(arr []string) [][]string {
	arrMap := make(map[string][]string)
	result := [][]string{}
	for _, value := range arr {
		sortedString := sortedString(value)
		arrMap[sortedString] = append(arrMap[sortedString], value)
	}
	for _, group := range arrMap {
		result = append(result, group)
	}
	return result
}

func sortedString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
