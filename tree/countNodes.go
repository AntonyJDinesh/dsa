package tree

import (
	"fmt"
	"math"
)

func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh, rh := getTreeHeights(root)
	if lh == rh {
		return int(math.Pow(2, float64(lh+1))) - 1
	}

	completeTreeCount := int(math.Pow(2, float64(lh))) - 1
	lastRowCount := lastRowNodeCount(root, lh)
	fmt.Printf("completeTreeCount: %d, lastRowCount: %d", completeTreeCount, lastRowCount)
	return completeTreeCount + lastRowCount
}

func lastRowNodeCount(root *TreeNode, d int) int {

	start, end := 1, int(math.Pow(2, float64(d)))-1
	var node *TreeNode
	for start <= end {
		pivot := start + int((end-start)/2)
		node = getNodeByRowAndIdx(root, d, pivot)
		fmt.Printf("start: %d, end: %d, pivot: %d, node: %v\n", start, end, pivot, node)
		if node == nil {
			end = pivot - 1
		} else {
			start = pivot + 1
		}
	}
	fmt.Printf("start: %d, end: %d", start, end)

	return start
}

func getNodeByRowAndIdx(root *TreeNode, d, idx int) *TreeNode {
	left, right := 0, int(math.Pow(2, float64(d)))-1
	node := root
	for i := 0; i < d; i++ {
		pivot := left + int((right-left)/2)
		if idx <= pivot {
			node = node.Left
			right = pivot
		} else {
			node = node.Right
			left = pivot + 1
		}
	}

	return node
}

func getTreeHeights(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	left := 0
	node := root
	for node.Left != nil {
		node = node.Left
		left++
	}

	right := 0
	node = root
	for node.Right != nil {
		node = node.Right
		right++
	}

	return left, right
}
