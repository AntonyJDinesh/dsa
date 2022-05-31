package heap

import (
	"fmt"
	"testing"
)

func TestMinHeap_Add(t *testing.T) {
	mh := NewHeap(MinHeap)
	mh.Add(8)
	if mh.Peek() != 8 {
		t.Errorf("Exp: %d, Got: %d", 4, mh.Peek())
	}
	mh.Add(5)
	if mh.Peek() != 5 {
		t.Errorf("Exp: %d, Got: %d", 5, mh.Peek())
	}
	mh.Add(6)
	if mh.Peek() != 5 {
		t.Errorf("Exp: %d, Got: %d", 4, mh.Peek())
	}
	mh.Add(7)
	mh.Print()
	mh.Add(1)
	mh.Print()
	mh.Add(2)
	mh.Add(3)
	mh.Add(4)

	if mh.Size() != 8 {
		t.Errorf("Exp: %d, Got: %d", 4, mh.Size())
	}

	mh.Print()

	fmt.Printf("Remove: %d\n", mh.Remove())
	if mh.Size() != 7 {
		t.Errorf("Exp: %d, Got: %d", 4, mh.Size())
	}
	mh.Print()
	fmt.Printf("Remove: %d\n", mh.Remove())
	mh.Print()
	fmt.Printf("Remove: %d\n", mh.Remove())
	mh.Print()
}
