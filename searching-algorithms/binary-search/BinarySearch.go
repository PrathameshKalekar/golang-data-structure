package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var key int
	fmt.Print("Enter value for search : ")
	fmt.Scan(&key)
	index, err := BinarySearch(arr, key)
	if err == nil {
		fmt.Println("Value found at index : ", index)
	} else {
		fmt.Println(err)
	}
}

func BinarySearch(arr []int, key int) (int, error) {
	sort.Ints(arr)

	low, high := 0, len(arr)-1

	for low <= high {
		median := (low + high) / 2

		if arr[median] == key {
			return median, nil
		}

		if arr[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	return 0, fmt.Errorf("Value not found in array")
}
