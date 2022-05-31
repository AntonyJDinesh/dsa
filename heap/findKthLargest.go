package heap

import (
	"container/heap"
)

type IntHeap []int

func (ih IntHeap) Len() int {
	return len(ih)
}

func (ih IntHeap) Less(i, j int) bool {
	return ih[i] > ih[j]
}

func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(ele interface{}) {
	*ih = append(*ih, ele.(int))
}

func (ih *IntHeap) Pop() (ele interface{}) {
	ele, *ih = (*ih)[len(*ih)-1], (*ih)[:len(*ih)-1]
	return
}

func findKthLargest(nums []int, k int) int {
	h := IntHeap{}
	h = append(h, nums...)

	heap.Init(&h) // n

	var res interface{}
	for i := 0; i < k; i++ {
		res = heap.Pop(&h)
	} // klog(n)

	// n + klog(n)
	return res.(int)
}
