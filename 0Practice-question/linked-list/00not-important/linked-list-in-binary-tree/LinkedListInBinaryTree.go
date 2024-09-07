/*
1367 Given a binary tree root and a linked list with head as the first node.

Return True if all the elements in the linked list starting from the head correspond to some downward path connected in the binary tree otherwise return False.

In this context downward path means a path that starts at some node and goes downwards.

Example 1:

Input: head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
Output: true
Explanation: Nodes in blue form a subpath in the binary Tree.
Example 2:

Input: head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
Output: true
Example 3:

Input: head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
Output: false
Explanation: There is no path in the binary tree that contains all the elements of the linked list from head.
*/
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Helper function to check if the linked list can be matched from the current tree node.
func isSubPathFromNode(head *ListNode, root *TreeNode) bool {
	// If the linked list is empty, we've matched the whole list.
	if head == nil {
		return true
	}
	// If the tree node is null, we can't match the linked list.
	if root == nil {
		return false
	}
	// If the values match, recursively check left and right subtrees.
	if head.Val == root.Val {
		return isSubPathFromNode(head.Next, root.Left) || isSubPathFromNode(head.Next, root.Right)
	}
	return false
}

// Main function to check if there's any subpath in the binary tree matching the linked list.
func isSubPath(head *ListNode, root *TreeNode) bool {
	// If the binary tree is empty, return false.
	if root == nil {
		return false
	}
	// Check if the linked list matches starting from the current node or check recursively in left/right subtrees.
	return isSubPathFromNode(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func main() {
	// Example 1
	head := &ListNode{4, &ListNode{2, &ListNode{8, nil}}}
	root := &TreeNode{1, &TreeNode{4, nil, &TreeNode{2, nil, &TreeNode{6, &TreeNode{8, nil, nil}, nil}}}, &TreeNode{4, nil, nil}}
	fmt.Println(isSubPath(head, root)) // Output: true

	// Example 2
	head2 := &ListNode{1, &ListNode{4, &ListNode{2, &ListNode{6, nil}}}}
	root2 := &TreeNode{1, &TreeNode{4, nil, &TreeNode{2, nil, &TreeNode{6, &TreeNode{8, nil, nil}, nil}}}, &TreeNode{4, nil, nil}}
	fmt.Println(isSubPath(head2, root2)) // Output: true

	// Example 3
	head3 := &ListNode{1, &ListNode{4, &ListNode{2, &ListNode{6, &ListNode{8, nil}}}}}
	root3 := &TreeNode{1, &TreeNode{4, nil, &TreeNode{2, nil, &TreeNode{6, &TreeNode{8, nil, nil}, nil}}}, &TreeNode{4, nil, nil}}
	fmt.Println(isSubPath(head3, root3)) // Output: false
}
