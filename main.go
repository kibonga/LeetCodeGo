package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(x int, l, r *TreeNode) *TreeNode {
	return &TreeNode{x, l, r}
}

func right_side_view(root *TreeNode) {
	nodes := []*TreeNode{}
	nodes = append(nodes, root)
	result := []int{}

	for len(nodes) != 0 {
		size := len(nodes)

		for i := 0; i < size; i++ {
			temp := nodes[0]
			nodes = nodes[1:]
			if i == (size - 1) {
				result = append(result, temp.val)
			}

			if temp.left != nil {
				nodes = append(nodes, temp.left)
			}
			if temp.right != nil {
				nodes = append(nodes, temp.right)
			}
		}
	}

	fmt.Print("Right side view = ")
	for _, v := range result {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	tn7 := newTreeNode(7, nil, nil)
	tn5 := newTreeNode(5, tn7, nil)
	tn4 := newTreeNode(4, nil, nil)
	tn2 := newTreeNode(2, nil, tn5)
	tn3 := newTreeNode(3, nil, tn4)
	tn1 := newTreeNode(1, tn2, tn3)

	right_side_view(tn1)
}
