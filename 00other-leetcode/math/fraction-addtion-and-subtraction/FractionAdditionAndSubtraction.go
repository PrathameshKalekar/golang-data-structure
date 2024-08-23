/*
592 Given a string expression representing an expression of fraction addition and subtraction, return the calculation result in string format.

The final result should be an irreducible fraction. If your final result is an integer, change it to the format of a fraction that has a denominator 1. So in this case, 2 should be converted to 2/1.

Example 1:

Input: expression = "-1/2+1/2"
Output: "0/1"
Example 2:

Input: expression = "-1/2+1/2+1/3"
Output: "1/3"
Example 3:

Input: expression = "1/3-1/2"
Output: "-1/6"
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to calculate the greatest common divisor (GCD)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the least common multiple (LCM)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to parse a fraction string and return numerator and denominator
func parseFraction(fraction string) (int, int) {
	parts := strings.Split(fraction, "/")
	numerator, _ := strconv.Atoi(parts[0])
	denominator, _ := strconv.Atoi(parts[1])
	return numerator, denominator
}

// Main function to evaluate the expression
func fractionAddition(expression string) string {
	// Normalize the expression by inserting '+' before negative signs
	normalizedExpr := strings.Replace(expression, "-", "+-", -1)
	if normalizedExpr[0] == '+' {
		normalizedExpr = normalizedExpr[1:]
	}

	// Split the expression into individual fractions
	fractions := strings.Split(normalizedExpr, "+")
	numerators := []int{}
	denominators := []int{}

	// Parse the fractions and store numerators and denominators separately
	for _, frac := range fractions {
		num, denom := parseFraction(frac)
		numerators = append(numerators, num)
		denominators = append(denominators, denom)
	}

	// Find the least common multiple (LCM) of all denominators
	commonDenom := denominators[0]
	for _, denom := range denominators[1:] {
		commonDenom = lcm(commonDenom, denom)
	}

	// Calculate the total numerator with respect to the common denominator
	totalNumerator := 0
	for i, num := range numerators {
		totalNumerator += num * (commonDenom / denominators[i])
	}

	// Reduce the final fraction using the GCD
	finalGcd := gcd(abs(totalNumerator), commonDenom)
	finalNumerator := totalNumerator / finalGcd
	finalDenominator := commonDenom / finalGcd

	return fmt.Sprintf("%d/%d", finalNumerator, finalDenominator)
}

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	expression1 := "-1/2+1/2"
	expression2 := "-1/2+1/2+1/3"
	expression3 := "1/3-1/2"

	fmt.Println(fractionAddition(expression1)) // Output: "0/1"
	fmt.Println(fractionAddition(expression2)) // Output: "1/3"
	fmt.Println(fractionAddition(expression3)) // Output: "-1/6"
}
