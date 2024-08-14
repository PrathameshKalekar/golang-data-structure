/*
25 Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
Example 2:

Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
*/
package main

import "fmt"

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}}}}
	fmt.Print("Before swap : ")
	node.Display()
	answer := reverseNodesInKGroup(node, 2)
	fmt.Print("after swap : ")
	answer.Display()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseNodesInKGroup(head *ListNode, k int) *ListNode {
	result := &ListNode{}
	dummy := result
	temp := []int{}
	//check each node
	for head != nil {
		// add each node on temp
		temp = append(temp, head.Val)
		//if the len of temp reaches k means the k elements we need to reverse is obtained
		if len(temp) == k {
			//reverse the temp add add the value on listnode
			for i := k - 1; i >= 0; i-- {
				dummy.Next = &ListNode{Val: temp[i]}
				dummy = dummy.Next
			}
			//make tamp empty after reversing
			temp = []int{}
		}
		//shift head node
		head = head.Next
	}
	//if all temp not added add the reamining ass is is cause they not maked the k group
	if len(temp) > 0 {
		for _, value := range temp {
			dummy.Next = &ListNode{Val: value}
			dummy = dummy.Next
		}
	}
	//return the result.Next cause first node is empty
	return result.Next
}

//for displaying the node
func (list *ListNode) Display() {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
