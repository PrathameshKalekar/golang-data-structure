package main

import "fmt"

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	InsertionSort(arr)
	fmt.Println(arr)
}

func InsertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
}
