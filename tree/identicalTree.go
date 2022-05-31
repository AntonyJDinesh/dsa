package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p == nil || q == nil {
        return false
    }

    left := isSameTree(p.Left, q.Left)
    right := isSameTree(p.Right, q.Right)

    return left && right && (p.Val == q.Val)
}