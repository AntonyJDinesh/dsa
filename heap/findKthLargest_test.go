package heap

import "testing"

func Test_findKthLargest(t *testing.T) {
	testData := []struct {
		nums   []int
		k, res int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for _, td := range testData {
		res := findKthLargest(td.nums, td.k)
		if res != td.res {
			t.Errorf("%#v, K: %d; Exp: %d, Got: %d", td.nums, td.k, td.res, res)
		}
	}
}
