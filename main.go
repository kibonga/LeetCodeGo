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

func swap(a, b *int) {
	c := a
	a = b
	b = c
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func max_height_dfs(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	left := max_height_dfs(root.left, diameter)
	right := max_height_dfs(root.right, diameter)

	if (left + right) > *diameter {
		*diameter = left + right
	}

	return max(left, right) + 1
}

func tree_diameter(root *TreeNode) int {
	diameter := 0
	max_height_dfs(root, &diameter)
	return diameter
}

func main() {
	tn6 := newTreNode(6, nil, nil)
	tn8 := newTreNode(8, nil, nil)

	tn4 := newTreNode(4, tn6, nil)
	tn5 := newTreNode(5, tn8, nil)

	// tn4 := newTreNode(4, nil, nil)
	// tn5 := newTreNode(5, nil, nil)

	// tn7 := newTreNode(7, nil, nil)

	tn2 := newTreNode(2, tn4, tn5)
	tn3 := newTreNode(3, nil, nil)

	tn1 := newTreNode(1, tn2, tn3)

	res := tree_diameter(tn1)
	fmt.Println("Tree diameter = ", res)
}
