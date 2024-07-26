package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var key int
	fmt.Print("Enter value for search : ")
	fmt.Scan(&key)
	index, err := LinearSearch(arr, key)
	if err == nil {
		fmt.Println("Value found at index : ", index)
	} else {
		fmt.Println(err)
	}
}

func LinearSearch(arr []int, key int) (int, error) {
	for index, value := range arr {
		if value == key {
			return index, nil
		}
	}
	return 0, fmt.Errorf("Value not found in the array")
}
