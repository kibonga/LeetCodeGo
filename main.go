package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type TreeBF struct {
	isBalanced bool
	height     int
}

func newTreeNode(x int, l, r *TreeNode) *TreeNode {
	return &TreeNode{x, l, r}
}

func isBalanced(root *TreeNode) bool {
	return balanceFactor(root).isBalanced
}

func balanceFactor(root *TreeNode) TreeBF {
	if root == nil {
		return TreeBF{true, 0}
	}

	leftTree, rightTree := balanceFactor(root.left), balanceFactor(root.right)

	diff := int(math.Abs(float64(leftTree.height) - float64(rightTree.height)))
	isBalanced := leftTree.isBalanced &&
		rightTree.isBalanced &&
		diff < 2

	return TreeBF{isBalanced, 1 + max(leftTree.height, rightTree.height)}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// tn15 := newTreeNode(15, nil, nil)
	// tn7 := newTreeNode(7, nil, nil)

	// tn9 := newTreeNode(9, nil, nil)
	// tn20 := newTreeNode(20, tn15, tn7)

	// tn3 := newTreeNode(3, tn9, tn20)

	tn4 := newTreeNode(4, nil, nil)
	tn44 := newTreeNode(4, nil, nil)

	tn3 := newTreeNode(3, tn4, tn44)
	tn33 := newTreeNode(3, nil, nil)

	tn2 := newTreeNode(2, tn3, tn33)
	tn22 := newTreeNode(2, nil, nil)

	tn1 := newTreeNode(1, tn2, tn22)

	res := isBalanced(tn1)
	fmt.Println("Is balanaced= ", res)
}
