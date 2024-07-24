package main

import "fmt"

func main() {

	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	SelectionSort(arr)
	fmt.Println(arr)
}

func SelectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		minIndex := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
