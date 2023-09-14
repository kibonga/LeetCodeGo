package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(x int, l, r *TreeNode) *TreeNode {
	return &TreeNode{
		val:   x,
		left:  l,
		right: r,
	}
}

func isSameTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.val != root2.val {
		return false
	}
	return isSameTree(root1.left, root2.left) && isSameTree(root1.right, root2.right)
}

func main() {
	tn2 := newTreeNode(2, nil, nil)
	tn3 := newTreeNode(3, nil, nil)
	tn1 := newTreeNode(1, tn2, tn3)

	tn22 := newTreeNode(2, nil, nil)
	tn33 := newTreeNode(3, nil, nil)
	tn11 := newTreeNode(1, tn22, tn33)

	res := isSameTree(tn1, tn11)
	fmt.Println("Result= ", res)
}
