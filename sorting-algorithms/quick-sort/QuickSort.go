package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)
	sortedArr := QuickSort(arr)
	fmt.Println("Sorted array:  ", sortedArr)
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]
	for index, _ := range arr {
		if arr[index] < arr[right] {
			arr[index], arr[left] = arr[left], arr[index]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	return arr
}
