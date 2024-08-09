/*
168 Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

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

Input: columnNumber = 1
Output: "A"
Example 2:

Input: columnNumber = 28
Output: "AB"
Example 3:

Input: columnNumber = 701
Output: "ZY"
*/
package main

import "fmt"

func main() {
	colNumber := 26
	answer := excelSheetColumnTitle(colNumber)
	fmt.Printf("The excel sheet column title of column %d is %s", colNumber, answer)
}

func excelSheetColumnTitle(colNumber int) string {
	result := ""
	for colNumber > 0 {
		colNumber--
		remainder := colNumber % 26
		result = string('A'+remainder) + result
		colNumber /= 26
	}
	return result
}
