package main

import "fmt"

func main() {
	//creating a hash table
	hashMap := make(map[string]int)

	hashMap["first"] = 1
	hashMap["second"] = 2
	hashMap["third"] = 3
	fmt.Println(hashMap)
	value, isExist := hashMap["second"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("No map found")
	}
}
