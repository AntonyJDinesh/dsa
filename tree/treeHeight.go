package tree


func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    lHeight := maxDepth(root.Left)
    rHeight := maxDepth(root.Right)

    if rHeight > lHeight {
        return rHeight +1
    }

    return lHeight + 1
}