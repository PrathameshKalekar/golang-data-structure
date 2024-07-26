package main

import (
	"fmt"
)

type Node struct {
	RightNode *Node
	LeftNode  *Node
	Value     int
}

type BinaryTree struct {
	RootNode *Node
}

func main() {
	var tree BinaryTree
	tree.Add(2)
	tree.Add(1)
	tree.Add(3)
	tree.Display(tree.RootNode)

}

// Add value in the binary tree
func (tree *BinaryTree) Add(value int) {
	newNode := &Node{
		Value: value,
	}

	if tree.RootNode == nil {
		tree.RootNode = newNode
		return
	}

	addNodeInTree(tree.RootNode, newNode)
}

func addNodeInTree(rootNode, newNode *Node) {

	if newNode.Value < rootNode.Value {
		if rootNode.LeftNode == nil {
			rootNode.LeftNode = newNode
		} else {
			addNodeInTree(rootNode.LeftNode, newNode)
		}
	} else {
		if rootNode.RightNode == nil {
			rootNode.RightNode = newNode
		} else {
			addNodeInTree(rootNode.RightNode, newNode)
		}
	}

}

// Display value  in the binary tree
func (tree *BinaryTree) Display(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Value)
	tree.Display(node.LeftNode)
	tree.Display(node.RightNode)
}
