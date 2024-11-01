package main

import "fmt"

func main() {
	s := "aaabaaa"
	answer := deleteCharactersToMakeFancyStrings(s)
	fmt.Println(answer)
}

func deleteCharactersToMakeFancyStrings(s string) string {
	if len(s) < 3 {
		return s
	}
	ans, tmp := []byte{}, []byte(s)
	for i := 2; i < len(s); i++ {
		b := false
		c := tmp[i-2]
		temp := tmp[i-2 : i+1]
		for j := 1; j < 3; j++ {
			if temp[j] != c {
				b = true
				break
			}
		}
		if b {
			ans = append(ans, c)
		}
	}
	return string(append(ans, tmp[len(tmp)-2:]...))

}
