package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	/*lca := &LCA{}
	lca.doRecursiveSearch(root, p, q)
	return lca.lca*/
	lca := new(*TreeNode)
	found := new(int)
	*found = 0
	doRecursiveSearch1(root, p, q, lca, found)
	return *lca
}


func doRecursiveSearch1(currentNode, p, q *TreeNode, lca **TreeNode, found *int) bool {

	if *found == 2 || *lca != nil {
		return false
	}

	if currentNode == nil {
		return false
	}

	var left, right, curr int
	if currentNode == p || currentNode == q {
		curr = 1
		*found++
	}

	if doRecursiveSearch1(currentNode.Left, p, q, lca, found) {
		left = 1
	}
	if doRecursiveSearch1(currentNode.Right, p, q, lca, found) {
		right = 1
	}

	if curr+left+right >= 2 {
		*lca = currentNode
	}
	
	if curr+left+right == 1 {
		return true
	}

	return false
}

type LCA struct {
	lca *TreeNode
}

func (lca *LCA) doRecursiveSearch(currentNode, p, q *TreeNode) bool {

	if currentNode == nil {
		return false
	}

	var left, right, curr int
	if currentNode == p || currentNode == q {
		curr = 1
	}
	if lca.doRecursiveSearch(currentNode.Left, p, q) {
		left = 1
	}
	if lca.doRecursiveSearch(currentNode.Right, p, q) {
		right = 1
	}

	if left+right+curr >= 2 {
		lca.lca = currentNode
	}

	if left+right+curr == 1 {
		return true
	}

	return false
}