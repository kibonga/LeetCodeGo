package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type Queue struct {
	items []interface{}
}

func (q *Queue) enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) size() int {
	return len(q.items)
}

func newTreeNode(x int, l, r *TreeNode) *TreeNode {
	return &TreeNode{x, l, r}
}

func level_order(root *TreeNode) {
	if root == nil {
		fmt.Println("Tree is empty")
		return
	}

	queue := Queue{}
	queue.enqueue(root)

	fmt.Print("Level order = ")
	for {
		if queue.isEmpty() {
			break
		}
		item := queue.dequeue().(*TreeNode)
		fmt.Print(item.val, " ")
		if item.left != nil {
			queue.enqueue(item.left)
		}
		if item.right != nil {
			queue.enqueue(item.right)
		}
	}
	fmt.Println()
}

func main() {
	tn15 := newTreeNode(15, nil, nil)
	tn7 := newTreeNode(7, nil, nil)

	tn9 := newTreeNode(9, nil, nil)
	tn20 := newTreeNode(20, tn15, tn7)

	tn3 := newTreeNode(3, tn9, tn20)

	level_order(tn3)
}
