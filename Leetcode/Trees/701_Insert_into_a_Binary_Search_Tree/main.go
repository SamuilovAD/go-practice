package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}

	root = insertIntoBST(root, 5)

	printInOrder(root)
	fmt.Println() // 1 2 3 4 5 7
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	var insertToBSTrecursive func(currentNode *TreeNode) *TreeNode
	insertToBSTrecursive = func(currentNode *TreeNode) *TreeNode {
		if currentNode == nil {
			return &TreeNode{Val: val}
		}

		if val > currentNode.Val {
			currentNode.Right = insertToBSTrecursive(currentNode.Right)
		} else {
			currentNode.Left = insertToBSTrecursive(currentNode.Left)
		}
		return currentNode
	}

	return insertToBSTrecursive(root)
}

func printInOrder(node *TreeNode) {
	if node == nil {
		return
	}

	printInOrder(node.Left)
	fmt.Print(node.Val, " ")
	printInOrder(node.Right)
}
