/*
725 Given the head of a singly linked list and an integer k, split the linked list into k consecutive linked list parts.

The length of each part should be as equal as possible: no two parts should have a size differing by more than one. This may lead to some parts being null.

The parts should be in the order of occurrence in the input list, and parts occurring earlier should always have a size greater than or equal to parts occurring later.

Return an array of the k parts.

Example 1:

Input: head = [1,2,3], k = 5
Output: [[1],[2],[3],[],[]]
Explanation:
The first element output[0] has output[0].val = 1, output[0].next = null.
The last element output[4] is null, but its string representation as a ListNode is [].
Example 2:

Input: head = [1,2,3,4,5,6,7,8,9,10], k = 3
Output: [[1,2,3,4],[5,6,7],[8,9,10]]
Explanation:
The input has been split into consecutive parts with size difference at most 1, and earlier parts are a larger size than the later parts.
*/
package main

import "fmt"

// ListNode defines a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to split the linked list into k parts
func splitLinkedListInParts(head *ListNode, k int) []*ListNode {
	// Step 1: Calculate the total length of the linked list
	length := 0
	for current := head; current != nil; current = current.Next {
		length++
	}

	// Step 2: Determine the base size of each part and the extra nodes
	partSize := length / k
	extraNodes := length % k

	// Step 3: Create an array to store the resulting parts
	parts := make([]*ListNode, k)
	current := head

	for i := 0; i < k && current != nil; i++ {
		parts[i] = current // Assign the current node as the start of the i-th part
		size := partSize
		if i < extraNodes { // First `extraNodes` parts get an extra node
			size++
		}

		// Move current pointer `size` times to create a sublist
		for j := 1; j < size && current != nil; j++ {
			current = current.Next
		}

		// Disconnect this part from the rest of the list
		if current != nil {
			nextPart := current.Next
			current.Next = nil
			current = nextPart
		}
	}

	return parts
}

// Helper function to print the linked list parts
func printParts(parts []*ListNode) {
	for i, part := range parts {
		fmt.Printf("Part %d: ", i+1)
		for part != nil {
			fmt.Printf("%d ", part.Val)
			part = part.Next
		}
		fmt.Println()
	}
}

// Helper function to create a linked list from a slice of integers
func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func main() {
	// Example 1: Input: head = [1,2,3], k = 5
	head := createLinkedList([]int{1, 2, 3})
	k := 5
	parts := splitLinkedListInParts(head, k)
	printParts(parts)

	// Example 2: Input: head = [1,2,3,4,5,6,7,8,9,10], k = 3
	head2 := createLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	k2 := 3
	parts2 := splitLinkedListInParts(head2, k2)
	printParts(parts2)
}
