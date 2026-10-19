package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}

		root.Val = successor.Val
		root.Right = deleteNode(root.Right, successor.Val)
	}

	return root
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

func main() {
	//         8
	//        / \
	//       5   10
	//      / \
	//     3   7
	//        /
	//       6

	root := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{
				Val:  7,
				Left: &TreeNode{Val: 6},
			},
		},
		Right: &TreeNode{Val: 10},
	}

	fmt.Print("Before: ")
	inorder(root)
	fmt.Println()

	root = deleteNode(root, 5)

	fmt.Print("After:  ")
	inorder(root)
	fmt.Println()
}