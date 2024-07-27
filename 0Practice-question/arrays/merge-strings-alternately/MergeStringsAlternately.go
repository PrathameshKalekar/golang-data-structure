/*
1768. You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

Example 1:

Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
Example 2:

Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b
word2:    p   q   r   s
merged: a p b q   r   s
Example 3:

Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q
merged: a p b q c   d
*/
package main

import "fmt"

func main() {
	firstString, secondString := "pahmesh", "rta"
	answer := mergeStringsAlternately(firstString, secondString)
	fmt.Printf("The mergered string is %s", answer)
}

func mergeStringsAlternately(firstString, secondString string) string {
	//for tracking each rune in the string
	firstLen, secondLen := len(firstString), len(secondString)
	i, j := 0, 0
	result := ""
	//run till all rune are added
	for i < firstLen || j < secondLen {
		//if first strings rune is there to add the add first strings i th run in string
		if i < firstLen {
			result += string(firstString[i])
			i++
		}
		//if second strings rune is there to add the add second strings i th run in string
		if j < secondLen {
			result += string(secondString[j])
			j++
		}
	}
	//return merged result
	return result
}
