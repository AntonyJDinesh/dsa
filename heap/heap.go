package heap

type HeapType int

const (
	MinHeap HeapType = 1
	MaxHeap HeapType = 1
)

type Heap interface {
	Add(val int)
	Remove() int
	Peek() int
	Size() int
	Print()
}

type heapNode struct {
	val         int
	right, left *heapNode
}

func NewHeap(heapType HeapType) (heap Heap) {
	if heapType == MinHeap {
		heap = &minHeap{bucket: make([]int, 1)}
	}

	return
}
