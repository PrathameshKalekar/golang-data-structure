/*
590 Given the root of an n-ary tree, return the postorder traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)

Example 1:

Input: root = [1,null,3,2,4,null,5,6]
Output: [5,6,3,2,4,1]
Example 2:

Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]
*/
package main

import "fmt"

// NaryTreeNode represents a node in an n-ary tree.
type NaryTreeNode struct {
	Val      int
	Children []*NaryTreeNode
}

// postorderTraversal performs a postorder traversal of an n-ary tree.
func postorderTraversal(root *NaryTreeNode) []int {
	var result []int
	postorderHelper(root, &result)
	return result
}

// Helper function to recursively traverse the tree.
func postorderHelper(node *NaryTreeNode, result *[]int) {
	if node == nil {
		return
	}
	for _, child := range node.Children {
		postorderHelper(child, result)
	}
	*result = append(*result, node.Val)
}

func main() {
	// Example 1
	root1 := &NaryTreeNode{Val: 1}
	child3 := &NaryTreeNode{Val: 3}
	child2 := &NaryTreeNode{Val: 2}
	child4 := &NaryTreeNode{Val: 4}
	root1.Children = []*NaryTreeNode{child3, child2, child4}
	child5 := &NaryTreeNode{Val: 5}
	child6 := &NaryTreeNode{Val: 6}
	child3.Children = []*NaryTreeNode{child5, child6}

	fmt.Println(postorderTraversal(root1)) // Output: [5, 6, 3, 2, 4, 1]

	// Example 2
	root2 := &NaryTreeNode{Val: 1}
	child22 := &NaryTreeNode{Val: 2}
	child23 := &NaryTreeNode{Val: 3}
	child24 := &NaryTreeNode{Val: 4}
	child25 := &NaryTreeNode{Val: 5}
	root2.Children = []*NaryTreeNode{child22, child23, child24, child25}
	child26 := &NaryTreeNode{Val: 6}
	child27 := &NaryTreeNode{Val: 7}
	child23.Children = []*NaryTreeNode{child26, child27}
	child211 := &NaryTreeNode{Val: 11}
	child27.Children = []*NaryTreeNode{child211}
	child28 := &NaryTreeNode{Val: 8}
	child24.Children = []*NaryTreeNode{child28}
	child212 := &NaryTreeNode{Val: 12}
	child28.Children = []*NaryTreeNode{child212}
	child29 := &NaryTreeNode{Val: 9}
	child210 := &NaryTreeNode{Val: 10}
	child25.Children = []*NaryTreeNode{child29, child210}
	child213 := &NaryTreeNode{Val: 13}
	child214 := &NaryTreeNode{Val: 14}
	child29.Children = []*NaryTreeNode{child213}
	child213.Children = []*NaryTreeNode{child214}

	fmt.Println(postorderTraversal(root2)) // Output: [2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1]
}
