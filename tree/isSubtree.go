package tree

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return identicalTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func identicalTree(node1, node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		return node1 == node2
	}

	left := identicalTree(node1.Left, node2.Left)
	right := identicalTree(node1.Right, node2.Right)
	return left && right && (node1.Val == node2.Val)
}
