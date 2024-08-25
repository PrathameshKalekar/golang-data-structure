/*
145 Given the root of a binary tree, return the postorder traversal of its nodes' values.

Example 1:

Input: root = [1,null,2,3]
Output: [3,2,1]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]
*/
package main

import "fmt"

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// postorderTraversal performs a postorder traversal of a binary tree.
func postorderTraversal(root *TreeNode) []int {
	var result []int
	traverse(root, &result)
	return result
}

// Helper function to recursively traverse the tree.
func traverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, result)         // Visit left subtree
	traverse(node.Right, result)        // Visit right subtree
	*result = append(*result, node.Val) // Visit node itself
}

func main() {
	// Example 1
	root1 := &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 3}
	fmt.Println(postorderTraversal(root1)) // Output: [3, 2, 1]

	// Example 2
	var root2 *TreeNode
	fmt.Println(postorderTraversal(root2)) // Output: []

	// Example 3
	root3 := &TreeNode{Val: 1}
	fmt.Println(postorderTraversal(root3)) // Output: [1]
}
