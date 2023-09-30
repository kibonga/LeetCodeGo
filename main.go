package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(x int) *TreeNode {
	return &TreeNode{x, nil, nil}
}

func indexOf(slice []int, x int) int {
	for i, val := range slice {
		if val == x {
			return i
		}
	}
	return -1
}

func build(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := newTreeNode(preorder[0])

	mid := indexOf(inorder, root.val)

	root.left = build(preorder[1:], inorder[:mid])
	root.right = build(preorder[mid+1:], inorder[mid+1:])

	return root
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	print(root.val, " ")
	preorder(root.left)
	preorder(root.right)
}

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}

	root := build(pre, in)
	print("Tree = ")
	preorder(root)
}
