package stack

type Stack struct {
	container []string
}

func NewStack() *Stack {
	return &Stack{container: make([]string, 0)}
}

func (s *Stack) Push(element string) {
	s.container = append(s.container, element)
}

func (s *Stack) Pop() (element string) {
	size := len(s.container)
	element, s.container = s.container[size-1], s.container[:size-1]
	return
}

func (s *Stack) Peek() string {
	return s.container[len(s.container)-1]
}

func (s *Stack) Size() int {
	return len(s.container)
}

func (s *Stack) HasNext() bool {
	return len(s.container) > 0
}

//=========================================

type IntStack struct {
	container []int
}

func NewIntStack() *IntStack {
	return &IntStack{container: make([]int, 0)}
}

func (s *IntStack) Push(element int) {
	s.container = append(s.container, element)
}

func (s *IntStack) Pop() (element int) {
	size := len(s.container)
	element, s.container = s.container[size-1], s.container[:size-1]
	return
}

func (s *IntStack) Peek() int {
	return s.container[len(s.container)-1]
}

func (s *IntStack) Size() int {
	return len(s.container)
}

func (s *IntStack) HasNext() bool {
	return len(s.container) > 0
}
