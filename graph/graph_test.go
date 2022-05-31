package graph

import "testing"

func TestStackImpl(t *testing.T) {

	stack := NewStack()
	stack.Push(1)
	if stack.Size() != 1 {
		t.Errorf("Push failed")
	}

	stack.Push(2)
	if stack.Size() != 2 {
		t.Errorf("Push failed")
	}

	if stack.Pop() != 2 && stack.Size() != 1 {
		t.Errorf("Pop failed")
	}
}
