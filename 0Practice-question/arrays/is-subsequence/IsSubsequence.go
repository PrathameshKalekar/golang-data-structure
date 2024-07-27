/*
392 Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).



Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false
*/

package main

import "fmt"

func main() {
	firstString, secondString := "abc", "ahbgc"
	answer := isSubsequence(firstString, secondString)
	fmt.Printf("Is string one and two are subsequence : %t", answer)

}
func isSubsequence(firstString, secondString string) bool {
	//keep tracking of each value in both string
	firstLen, secondLen := len(firstString), len(secondString)
	i, j := 0, 0
	//both string at i th and j th index matched add i++ to scan next value
	for i < firstLen && j < secondLen {
		if firstString[i] == secondString[j] {
			i++
		}
		j++
	}
	//if the i == firstlen means all rune in string ar present in second so it will return true if not the false
	return i == firstLen
}
