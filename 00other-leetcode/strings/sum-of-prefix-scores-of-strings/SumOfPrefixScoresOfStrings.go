/*
2416 You are given an array words of size n consisting of non-empty strings.

We define the score of a string word as the number of strings words[i] such that word is a prefix of words[i].

For example, if words = ["a", "ab", "abc", "cab"], then the score of "ab" is 2, since "ab" is a prefix of both "ab" and "abc".
Return an array answer of size n where answer[i] is the sum of scores of every non-empty prefix of words[i].

Note that a string is considered as a prefix of itself.

Example 1:

Input: words = ["abc","ab","bc","b"]
Output: [5,4,3,2]
Explanation: The answer for each string is the following:
- "abc" has 3 prefixes: "a", "ab", and "abc".
- There are 2 strings with the prefix "a", 2 strings with the prefix "ab", and 1 string with the prefix "abc".
The total is answer[0] = 2 + 2 + 1 = 5.
- "ab" has 2 prefixes: "a" and "ab".
- There are 2 strings with the prefix "a", and 2 strings with the prefix "ab".
The total is answer[1] = 2 + 2 = 4.
- "bc" has 2 prefixes: "b" and "bc".
- There are 2 strings with the prefix "b", and 1 string with the prefix "bc".
The total is answer[2] = 2 + 1 = 3.
- "b" has 1 prefix: "b".
- There are 2 strings with the prefix "b".
The total is answer[3] = 2.
Example 2:

Input: words = ["abcd"]
Output: [4]
Explanation:
"abcd" has 4 prefixes: "a", "ab", "abc", and "abcd".
Each prefix has a score of one, so the total is answer[0] = 1 + 1 + 1 + 1 = 4.
*/
package main

import "fmt"

// TrieNode represents each node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	count    int
}

// Trie represents the Trie structure
type Trie struct {
	root *TrieNode
}

// NewTrie creates and returns a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

// Insert inserts a word into the Trie and updates prefix counts
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
		node.count++ // increment count at each prefix
	}
}

// GetPrefixScore calculates the score of all prefixes of the given word
func (t *Trie) GetPrefixScore(word string) int {
	node := t.root
	score := 0
	for _, ch := range word {
		if node.children[ch] == nil {
			return score
		}
		node = node.children[ch]
		score += node.count // add the count of this prefix
	}
	return score
}

// sumPrefixScores calculates the sum of prefix scores for each word in the list
func sumPrefixScores(words []string) []int {
	trie := NewTrie()

	// Insert all words into the Trie
	for _, word := range words {
		trie.Insert(word)
	}

	// Calculate prefix scores for each word
	result := make([]int, len(words))
	for i, word := range words {
		result[i] = trie.GetPrefixScore(word)
	}

	return result
}

func main() {
	// Example 1
	words := []string{"abc", "ab", "bc", "b"}
	fmt.Println(sumPrefixScores(words)) // Output: [5, 4, 3, 2]

	// Example 2
	words2 := []string{"abcd"}
	fmt.Println(sumPrefixScores(words2)) // Output: [4]
}
