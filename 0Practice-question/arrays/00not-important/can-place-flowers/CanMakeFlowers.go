/*
605 You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: false
*/
package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	length := len(flowerbed)

	for i := 0; i < length; i++ {
		// Check if the current plot is empty and the left and right neighbors are also empty or out of bounds
		if flowerbed[i] == 0 {
			// Check left neighbor (i == 0 means it's the first plot, no left neighbor)
			left := (i == 0) || (flowerbed[i-1] == 0)
			// Check right neighbor (i == length-1 means it's the last plot, no right neighbor)
			right := (i == length-1) || (flowerbed[i+1] == 0)

			if left && right {
				// Plant a flower here
				flowerbed[i] = 1
				count++
				// If we've planted enough flowers, return true
				if count >= n {
					return true
				}
			}
		}
	}

	// If we planted less than n flowers, return false
	return count >= n
}

func main() {
	// Test example 1
	flowerbed1 := []int{1, 0, 0, 0, 1}
	n1 := 1
	fmt.Println(canPlaceFlowers(flowerbed1, n1)) // Output: true

	// Test example 2
	flowerbed2 := []int{1, 0, 0, 0, 1}
	n2 := 2
	fmt.Println(canPlaceFlowers(flowerbed2, n2)) // Output: false
}
