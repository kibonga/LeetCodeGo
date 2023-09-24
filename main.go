package main

import "fmt"

var GOOD_NODES int = 0

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(x int, l, r *TreeNode) *TreeNode {
	return &TreeNode{x, l, r}
}

func count_good_nodes(root *TreeNode, max_ancestor int) {
	if root == nil {
		return
	}

	if root.val >= max_ancestor {
		GOOD_NODES++
		max_ancestor = root.val
	}

	count_good_nodes(root.left, max_ancestor)
	count_good_nodes(root.right, max_ancestor)
}

func main() {
	tn4 := newTreeNode(4, nil, nil)
	tn2 := newTreeNode(2, nil, nil)
	tn3 := newTreeNode(3, tn4, tn2)
	tn33 := newTreeNode(3, tn3, nil)

	count_good_nodes(tn33, tn33.val)

	fmt.Println("Good nodes = ", GOOD_NODES)
}
