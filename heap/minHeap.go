package heap

import "fmt"

type minHeap struct {
	bucket []int
}

func (mh *minHeap) Add(val int) {
	mh.bucket = append(mh.bucket, val)
	index := len(mh.bucket) - 1
	parent := index / 2

	for index > 1 && mh.bucket[index] < mh.bucket[parent] {
		mh.bucket[index], mh.bucket[parent] = mh.bucket[parent], mh.bucket[index]
		index, parent = parent, parent/2
	}
}

func (mh *minHeap) Remove() (val int) {
	val = mh.bucket[1]
	mh.bucket[1] = mh.bucket[len(mh.bucket)-1]
	mh.bucket = mh.bucket[0 : len(mh.bucket)-1]

	parent := 1
	left, right := parent*2, parent*2+1
	for parent < int(len(mh.bucket)/2) {
		if mh.bucket[left] < mh.bucket[parent] || mh.bucket[right] < mh.bucket[parent] {
			if mh.bucket[left] < mh.bucket[right] {
				mh.bucket[left], mh.bucket[parent] = mh.bucket[parent], mh.bucket[left]
				parent = left
			} else {
				mh.bucket[right], mh.bucket[parent] = mh.bucket[parent], mh.bucket[right]
				parent = right
			}
		} else {
			break
		}
		left, right = parent*2, parent*2+1
	}

	return
}

func (mh *minHeap) Size() int {
	return len(mh.bucket) - 1
}

func (mh *minHeap) Peek() int {
	return mh.bucket[1]
}

func (mh *minHeap) Print() {
	fmt.Printf("%#v\n", mh.bucket[1:])
}
