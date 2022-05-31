package bs

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	testData := []struct {
		arr []int
		x   int
		k   int
		res []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, -1, 4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 2, 1, []int{2}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, []int{3, 4, 5, 6}},
		{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 5, 3, []int{3, 3, 4}},
		{[]int{2, -1, 1, 2, 3, 4, 5}, 3, 7, []int{2, -1, 1, 2, 3, 4, 5}},
		{[]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8}, 2, 2, []int{1, 3}},
	}
	for _, td := range testData {
		res := findClosestElements(td.arr, td.k, td.x)
		if !reflect.DeepEqual(res, td.res) {
			t.Errorf("arr: %#v, x: %d, k: %d,  => Exected: %#v, Got: %#v\n", td.arr, td.x, td.k, td.res, res)
		}
	}
}

func Test_Slice(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Logf("s: %d, e: %d, arr: %#v\n", 0, 1, arr[0:len(arr)])
	// t.Logf("s: %d, e: %d, arr: %#v\n", 4, 1, arr[4:1])
}
