package array

import (
	"reflect"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	testData := []struct {
		nums []int
		res  int
	}{
		/*	{nums: []int{1, 7, 3, 6, 5, 6}, res: 3},
			{nums: []int{1, 2, 3}, res: -1},
			{nums: []int{2, 1, -1}, res: 0},*/
		{nums: []int{-1, -1, -1, -1, -1, 0}, res: 2},
	}

	for _, td := range testData {
		res := pivotIndex(td.nums)
		if res != td.res {
			t.Errorf("%#v, Exp: %d, Got: %d", td.nums, td.res, res)
		}
	}
}

func Test_findKthLargest(t *testing.T) {

	testData := []struct {
		nums   []int
		k, res int
	}{
		{nil, 0, 0},
		{[]int{}, 0, 0},
		{[]int{3, 2, 5, 4, 1}, 10, 0},
		{[]int{3, 2, 5, 4, 1}, 4, 2},
		// 1, 4, 7, 10, 27, 48, 50, 100
		{[]int{10, 1, 7, 100, 4, 50, 48, 27}, 3, 48},
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 1, 2},
		{[]int{1, 1, 1}, 3, 1},
		{[]int{1, 1, 2}, 3, 1},
		{[]int{3, 2, 1}, 2, 2},
		{[]int{3, 2, 1, 5, 6, 6, 7}, 3, 6},
	}

	for _, td := range testData {
		res := findKthLargest(td.nums, td.k)
		if res != td.res {
			t.Errorf("%#v - K: %d, Exp: %d, Got: %d", td.nums, td.k, td.res, res)
		}
	}
}

func Test_quickWiggleSort(t *testing.T) {
	testData := []struct {
		nums, res []int
	}{
		{[]int{1, 3, 2}, []int{1, 3, 2}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
	}

	for _, td := range testData {
		wiggleSort(td.nums)
		if !reflect.DeepEqual(td.nums, td.res) {
			t.Errorf("Exp: %#v, Got: %#v", td.res, td.nums)
		}
	}
}

func Test_mod(t *testing.T) {
	t.Logf("%d", -1%10)
}
