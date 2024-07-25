package main

import (
	"fmt"
)

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)
	sortedArr := MergeSort(arr)
	fmt.Println("Sorted array:  ", sortedArr)
}
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	n := len(arr)
	median := int(n / 2)
	leftArray := MergeSort(arr[:median])
	rightArray := MergeSort(arr[median:])

	return merge(leftArray, rightArray)
}

func merge(leftArray, rightArray []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] < rightArray[j] {
			result = append(result, leftArray[i])
			i++
		} else {
			result = append(result, rightArray[j])
			j++
		}
	}
	result = append(result, leftArray[i:]...)
	result = append(result, rightArray[j:]...)
	return result
}
