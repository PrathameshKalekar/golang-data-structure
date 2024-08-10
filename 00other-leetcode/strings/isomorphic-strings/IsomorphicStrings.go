/*
205 Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
Example 3:

Input: s = "paper", t = "title"
Output: true
*/
package main

import "fmt"

func main() {
	s := "foo"
	t := "baa"
	answer := isomorphicStirngs(s, t)
	fmt.Println(answer)
}

func isomorphicStirngs(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	//make two map for s and t
	sMap := map[byte]byte{}
	tMap := map[byte]byte{}

	//for each element in smap and tmap
	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		//if mapped char exist check if it matches for tmap and smap both and if not add them in map
		if mappedChar, exist := tMap[tChar]; exist {
			if mappedChar != sChar {
				//if char not matched then return flase
				return false
			}
		} else {
			tMap[tChar] = sChar
		}
		//same as tmap
		if mappedChar, exist := sMap[sChar]; exist {
			if mappedChar != tChar {
				return false
			}
		} else {
			sMap[sChar] = tChar
		}
	}
	//if all condition satisfies then its a isomorphic strings then return true
	return true
}
