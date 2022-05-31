package find_and_union

import (
	"fmt"
	"testing"
)

func TestQuickUnionImpl_Connected(t *testing.T) {
	qui := newQuickUnionImpl(10)

	// 1-2-5-6-7 3-8-9 4
	qui.Union(1, 2)
	qui.Union(2, 5)
	qui.Union(5, 6)
	qui.Union(3, 8)
	qui.Union(8, 9)
	qui.Union(6, 7)

	testData := []struct {
		v1, v2    int
		connected bool
	}{
		{1, 5, true},
		{5, 7, true},
		{4, 9, false},
	}

	for _, td := range testData {
		res := qui.Connected(td.v1, td.v2)
		if res != td.connected {
			t.Errorf("v1: %d, v2: %d. Exp: %t, Got: %t", td.v1, td.v2, td.connected, res)
		}
	}

	qui.Union(9, 4)
	if !qui.Connected(4, 9) {
		t.Errorf("v1: %d, v2: %d. Exp: %t, Got: %t", 4, 9, true, false)
	}

	fmt.Printf("qfi: %#v\n", qui)
}
