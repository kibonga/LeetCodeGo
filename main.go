package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func newTreNode(x int, l, r *TreeNode) *TreeNode {
	return &TreeNode{x, l, r}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func max_height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(max_height(root.left), max_height(root.right)) + 1
}

func tree_diameter(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max_height(root.left) + max_height(root.right)
}

func main() {
	tn4 := newTreNode(4, nil, nil)
	tn5 := newTreNode(5, nil, nil)

	tn2 := newTreNode(2, tn4, tn5)
	tn3 := newTreNode(3, nil, nil)

	tn1 := newTreNode(1, tn2, tn3)

	res := tree_diameter(tn1)
	fmt.Println("Tree diameter = ", res)
}
