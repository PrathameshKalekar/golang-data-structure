/*
2 You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/
package main

import "fmt"

type ListNode struct {
	Val      int
	NextNode *ListNode
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	printNode(result) // Output: 7 -> 0 -> 8

	// Example 2
	l1 = &ListNode{0, nil}
	l2 = &ListNode{0, nil}
	result = addTwoNumbers(l1, l2)
	printNode(result) // Output: 0

	// Example 3
	l1 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	result = addTwoNumbers(l1, l2)
	printNode(result)
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	current := temp
	carry := 0
	//check until oth node is nil means no more to add
	for l1 != nil || l1 != nil {
		sum := carry
		// if l1 is not empty add l1 node value in sum
		if l1 != nil {
			sum += l1.Val
			l1 = l1.NextNode
		}
		// if l2 is not empty add l2 node value in sum
		if l2 != nil {
			sum += l2.Val
			l2 = l2.NextNode
		}
		carry = sum / 10
		current.NextNode = &ListNode{Val: sum % 10}
		current = current.NextNode
	}
	//after all element carry is remaining then add it to next node
	if carry > 0 {
		current.NextNode = &ListNode{Val: carry}
	}
	//first node is empty therefore return temp.nextNode
	return temp.NextNode
}

func printNode(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.NextNode != nil {
			fmt.Print(" -> ")
		}
		node = node.NextNode
	}
	fmt.Println()
}
