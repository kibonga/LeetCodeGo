package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func swap(a, b interface{}) {
	pA, okA := a.(*TreeNode)
	pB, okB := b.(*TreeNode)

	if !okA || !okB {
		fmt.Println("Failed to swap tree nodes, both A and B must be of type TreeNode")
		return
	}

	if pA == nil || pB == nil {
		return
	}

	*pA, *pB = *pB, *pA
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	swap(root.left, root.right)
	invertTree(root.left)
	invertTree(root.right)
	return root
}

func display(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.val)
	display(root.left)
	display(root.right)
}

func newTreeNode(x int, l, r *TreeNode) *TreeNode {
	return &TreeNode{
		left:  l,
		right: r,
		val:   x,
	}
}

func main() {
	tn1 := newTreeNode(1, nil, nil)
	tn3 := newTreeNode(3, nil, nil)
	tn6 := newTreeNode(6, nil, nil)
	tn9 := newTreeNode(9, nil, nil)

	tn2 := newTreeNode(2, tn1, tn3)
	tn7 := newTreeNode(7, tn6, tn9)

	tn4 := newTreeNode(4, tn2, tn7)

	fmt.Print("Tree= ")
	display(tn4)
	res := invertTree(tn4)
	fmt.Print("\n")
	fmt.Print("Tree= ")
	display(res)
}
