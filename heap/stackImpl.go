package heap

type Stack struct {
	container []*heapNode
}

func NewStack() *Stack {
	return &Stack{container: make([]*heapNode, 0)}
}

func (s *Stack) Push(element *heapNode) {
	s.container = append(s.container, element)
}

func (s *Stack) Pop() (element *heapNode) {
	size := len(s.container)
	element, s.container = s.container[size-1], s.container[:size-1]
	return
}

func (s *Stack) Peek() *heapNode {
	return s.container[len(s.container)-1]
}

func (s *Stack) Size() int {
	return len(s.container)
}

func (s *Stack) HasNext() bool {
	return len(s.container) > 0
}
