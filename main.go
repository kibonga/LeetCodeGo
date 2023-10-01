package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type Result struct {
	max int
}

func newResult() *Result {
	return &Result{0}
}

func newTreeNode(x int, l, r *TreeNode) *TreeNode {
	return &TreeNode{x, l, r}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func dfs(root *TreeNode, res *Result) int {

	if root == nil {
		return 0
	}

	left_subtree := dfs(root.left, res)
	right_subtree := dfs(root.right, res)

	sum := left_subtree + right_subtree + root.val

	if sum > res.max {
		res.max = sum
	}

	if (left_subtree <= 0) && (right_subtree <= 0) {
		return root.val
	}

	return max(left_subtree, right_subtree) + root.val
}

func max_path_sum(root *TreeNode) int {
	res := newResult()
	dfs(root, res)
	return res.max
}

func main() {
	tn15 := newTreeNode(15, nil, nil)
	tn7 := newTreeNode(7, nil, nil)
	tn9 := newTreeNode(9, nil, nil)
	tn20 := newTreeNode(20, tn15, tn7)
	tn_10 := newTreeNode(-10, tn9, tn20)

	max := max_path_sum(tn_10)
	print("Max path = ", max, "\n")
}
