/*
138 A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

val: an integer representing Node.val
random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
Your code will only be given the head of the original linked list.

Example 1:

Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
Example 2:

Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]
Example 3:

Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]
*/
package main

import "fmt"

func main() {
	// Creating a sample list with random pointers.
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	node1.Random = node3 // Random pointer from node1 to node3
	node2.Random = node1 // Random pointer from node2 to node1
	node3.Random = node4 // Random pointer from node3 to node4
	node4.Random = node2 // Random pointer from node4 to node2

	fmt.Println("Original list:")
	printList(node1)

	// Call the function to copy the list.
	copiedList := copyListWithRandomPointer(node1)

	fmt.Println("\nCopied list:")
	printList(copiedList)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Helper function to print the list for verification.
func printList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Printf("Node: %d", curr.Val)
		if curr.Random != nil {
			fmt.Printf(", Random: %d\n", curr.Random.Val)
		} else {
			fmt.Println(", Random: nil")
		}
		curr = curr.Next
	}
}

func copyListWithRandomPointer(head *Node) *Node {
	if head == nil {
		return head
	}
	curr := head
	//add dummy node with same value to each node
	for curr != nil {
		next := &Node{Val: curr.Val}
		next.Next = curr.Next
		curr.Next = next
		curr = next.Next
	}
	curr = head
	//for that dummy node add the random node as per the original
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}
	curr = head
	copyHead := head.Next
	copyCurr := copyHead
	//saperate the original and the copy node
	for curr != nil {
		curr.Next = curr.Next.Next
		if copyCurr.Next != nil {
			copyCurr.Next = copyCurr.Next.Next
		}
		curr = curr.Next
		copyCurr = copyCurr.Next
	}
	//return the copy node
	return copyHead
}
