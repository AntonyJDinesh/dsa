package heap

import "errors"

type Queue interface {
	Enqueue(ele any) error
	Dequeue() (any, error)
	Size() int
}

type queueImpl struct {
	maxSize int
	bucket  []any
}

func NewQueue(size int) Queue {
	return &queueImpl{maxSize: size}
}

func (qi *queueImpl) Enqueue(ele any) (err error) {
	if qi.Size() >= qi.maxSize {
		return errors.New("queue full")
	}

	qi.bucket = append(qi.bucket, ele)
	return
}

func (qi *queueImpl) Dequeue() (ele any, err error) {
	if qi.Size() == 0 {
		return nil, errors.New("queue is empty")
	}

	ele, qi.bucket = qi.bucket[0], qi.bucket[1:]

	return
}

func (qi *queueImpl) Size() int {
	return len(qi.bucket)
}
