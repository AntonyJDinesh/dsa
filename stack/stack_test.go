package stack

import (
	"os"
	"testing"
)

func TestMinStack(t *testing.T) {
	stk := Constructor()
	stk.Push(2)
	stk.Push(0)
	stk.Push(3)
	stk.Push(0)

	stk.print()

	res := stk.GetMin()
	if res != 0 {
		t.Errorf("stk.GetMin(): Exp: %d, Got: %d", 0, res)
	}

	stk.Pop()

	stk.print()

	res = stk.GetMin()
	if res != 0 {
		t.Errorf("stk.GetMin(): Exp: %d, Got: %d", 0, res)
	}

	stk.Pop()
	stk.print()

	res = stk.GetMin()
	if res != 0 {
		t.Errorf("stk.GetMin(): Exp: %d, Got: %d", 0, res)
	}

	stk.Pop()
	stk.print()
	os.Exit(1)

	res = stk.GetMin()
	if res != 2 {
		stk.print()
		t.Errorf("stk.GetMin(): Exp: %d, Got: %d", 2, res)
	}
}
