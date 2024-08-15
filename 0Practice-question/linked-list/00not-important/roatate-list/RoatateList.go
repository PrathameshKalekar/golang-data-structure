/*
61 Given the head of a linked list, rotate the list to the right by k places.

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]
Example 2:

Input: head = [0,1,2], k = 4
Output: [2,0,1]
*/
package main

import "fmt"

func main() {
	node := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	fmt.Print("Before swap : ")
	node.Display()
	answer := rotateList(node, 2)
	fmt.Print("after swap : ")
	answer.Display()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//for displaying the node
func (list *ListNode) Display() {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}

func rotateList(head *ListNode, k int) *ListNode {
	//if head is nil or single node or k = 0 then return the head as it is cause no roatation possible
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	//manage the count and to change the node
	count, change := head, head
	countValue := 0
	for count != nil {
		countValue++
		count = count.Next
	}
	//accurate the swap value
	k = k % countValue
	if k == 0 {
		return head
	}
	dummy := change
	//reach at the swap possition in the node
	for i := 0; i < countValue-k-1; i++ {
		dummy = dummy.Next
	}
	//the next nodes make is saperate and make the next node of if as nil
	next := dummy.Next
	result := next
	dummy.Next = nil
	//at the end of the node add the remaining part
	for next != nil && next.Next != nil {
		next = next.Next
	}
	next.Next = change
	//return the result
	return result
}
