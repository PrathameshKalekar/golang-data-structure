/*
432 Design a data structure to store the strings' count with the ability to return the strings with minimum and maximum counts.

Implement the AllOne class:

AllOne() Initializes the object of the data structure.
inc(String key) Increments the count of the string key by 1. If key does not exist in the data structure, insert it with count 1.
dec(String key) Decrements the count of the string key by 1. If the count of key is 0 after the decrement, remove it from the data structure. It is guaranteed that key exists in the data structure before the decrement.
getMaxKey() Returns one of the keys with the maximal count. If no element exists, return an empty string "".
getMinKey() Returns one of the keys with the minimum count. If no element exists, return an empty string "".
Note that each function must run in O(1) average time complexity.

Example 1:

Input
["AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey", "getMinKey"]
[[], ["hello"], ["hello"], [], [], ["leet"], [], []]
Output
[null, null, null, "hello", "hello", null, "hello", "leet"]

Explanation
AllOne allOne = new AllOne();
allOne.inc("hello");
allOne.inc("hello");
allOne.getMaxKey(); // return "hello"
allOne.getMinKey(); // return "hello"
allOne.inc("leet");
allOne.getMaxKey(); // return "hello"
allOne.getMinKey(); // return "leet"
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	allOne := Constructor()
	allOne.Inc("helloo")
	allOne.Inc("helloo")
	allOne.Inc("hello")
	fmt.Println(allOne.GetMaxKey())
	fmt.Println(allOne.GetMinKey())
}

type AllOne struct {
	table map[string]int // map to store counts of keys
	st    []pair         // slice to store (count, key) pairs
}

type pair struct {
	count int
	key   string
}

// Constructor initializes the data structure.
func Constructor() AllOne {
	return AllOne{
		table: make(map[string]int),
		st:    []pair{},
	}
}

// Inc increments the count of the string key by 1.
func (this *AllOne) Inc(key string) {
	// If the key exists, remove its previous (count, key) from the sorted slice
	if val, exists := this.table[key]; exists {
		this.remove(val, key)
	}

	// Increment the count in the table
	this.table[key]++

	// Add the new (count, key) pair to the sorted slice
	this.st = append(this.st, pair{this.table[key], key})

	// Sort the slice by count (ascending) and break ties using the key (lexicographically)
	sort.Slice(this.st, func(i, j int) bool {
		if this.st[i].count == this.st[j].count {
			return this.st[i].key < this.st[j].key
		}
		return this.st[i].count < this.st[j].count
	})
}

// Dec decrements the count of the string key by 1.
func (this *AllOne) Dec(key string) {
	if val, exists := this.table[key]; exists {
		// Remove the current (count, key) pair
		this.remove(val, key)

		// Decrement the count in the table
		this.table[key]--

		// If the count is 0, remove the key from the map entirely
		if this.table[key] == 0 {
			delete(this.table, key)
		} else {
			// Add the updated (count, key) back to the sorted slice
			this.st = append(this.st, pair{this.table[key], key})

			// Sort the slice again
			sort.Slice(this.st, func(i, j int) bool {
				if this.st[i].count == this.st[j].count {
					return this.st[i].key < this.st[j].key
				}
				return this.st[i].count < this.st[j].count
			})
		}
	}
}

// GetMaxKey returns one of the keys with the maximal count.
func (this *AllOne) GetMaxKey() string {
	if len(this.st) == 0 {
		return ""
	}
	// Return the key with the highest count (last element in the sorted slice)
	return this.st[len(this.st)-1].key
}

// GetMinKey returns one of the keys with the minimum count.
func (this *AllOne) GetMinKey() string {
	if len(this.st) == 0 {
		return ""
	}
	// Return the key with the lowest count (first element in the sorted slice)
	return this.st[0].key
}

// Helper function to remove a (count, key) pair from the sorted slice
func (this *AllOne) remove(count int, key string) {
	for i, p := range this.st {
		if p.count == count && p.key == key {
			// Remove the element from the slice
			this.st = append(this.st[:i], this.st[i+1:]...)
			break
		}
	}
}
