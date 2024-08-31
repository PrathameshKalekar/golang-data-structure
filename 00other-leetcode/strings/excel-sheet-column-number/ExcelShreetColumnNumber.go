/*
171 Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

Example 1:

Input: columnTitle = "A"
Output: 1
Example 2:

Input: columnTitle = "AB"
Output: 28
Example 3:

Input: columnTitle = "ZY"
Output: 701
*/
package main

import (
	"fmt"
)

// titleToNumber function converts a column title to its corresponding column number.
func titleToNumber(columnTitle string) int {
	result := 0
	for i := 0; i < len(columnTitle); i++ {
		// Convert the character to a number: 'A' -> 1, 'B' -> 2, ..., 'Z' -> 26
		value := int(columnTitle[i] - 'A' + 1)
		result = result*26 + value
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(titleToNumber("A"))       // Output: 1
	fmt.Println(titleToNumber("AB"))      // Output: 28
	fmt.Println(titleToNumber("ZY"))      // Output: 701
	fmt.Println(titleToNumber("FXSHRXW")) // Output: 2147483647 (test upper bound)
}
