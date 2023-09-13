package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.left), maxDepth(root.right)) + 1
}

func newTreeNode(x int, l, r *TreeNode) *TreeNode {
	return &TreeNode{
		val:   x,
		left:  l,
		right: r,
	}
}

func main() {
	tn9 := newTreeNode(9, nil, nil)
	tn15 := newTreeNode(15, nil, nil)
	tn7 := newTreeNode(7, nil, nil)

	tn20 := newTreeNode(20, tn15, tn7)

	tn3 := newTreeNode(3, tn9, tn20)

	max := maxDepth(tn3)
	fmt.Println("Max depth= ", max)
}
