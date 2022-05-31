package find_and_union

import (
	"fmt"
	"testing"
)

func TestQuickFindImpl_Connected(t *testing.T) {
	qfi := newQuickFindImpl(10)

	// 1-2-5-6-7 3-8-9 4
	qfi.Union(1, 2)
	qfi.Union(2, 5)
	qfi.Union(5, 6)
	qfi.Union(3, 8)
	qfi.Union(8, 9)
	qfi.Union(6, 7)

	testData := []struct {
		v1, v2    int
		connected bool
	}{
		{1, 5, true},
		{5, 7, true},
		{4, 9, false},
	}

	for _, td := range testData {
		res := qfi.Connected(td.v1, td.v2)
		if res != td.connected {
			t.Errorf("v1: %d, v2: %d. Exp: %t, Got: %t", td.v1, td.v2, td.connected, res)
		}
	}

	qfi.Union(9, 4)
	if !qfi.Connected(4, 9) {
		t.Errorf("v1: %d, v2: %d. Exp: %t, Got: %t", 4, 9, true, false)
	}

	fmt.Printf("qfi: %#v\n", qfi)
}
