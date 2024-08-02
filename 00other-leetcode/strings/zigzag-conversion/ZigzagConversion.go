/*
6 The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);


Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:

Input: s = "A", numRows = 1
Output: "A"
*/

package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	answer := zigzagConversion(s, 4)
	fmt.Println(answer)
}

func zigzagConversion(s string, numRow int) string {
	row := make([]string, numRow)
	currentRow, step := 0, 1
	for _, char := range s {
		row[currentRow] += string(char)
		if currentRow == 0 {
			step = 1
		} else if currentRow == numRow-1 {
			step = -1
		}
		currentRow += step
	}
	result := ""
	for _, char := range row {
		result += char
	}
	return result
}
