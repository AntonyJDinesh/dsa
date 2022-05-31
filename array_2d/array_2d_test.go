package array_2d

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {

	testData := []struct {
		matrix [][]int
		res    []int
	}{
		{
			/*
				1  2  3  4
				5  6  7  8
				9 10 11 12
			*/
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			res:    []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, td := range testData {
		res := spiralOrder(td.matrix)
		if !reflect.DeepEqual(res, td.res) {
			t.Errorf("Matrix: %#v, Exp: %#v, Got: %#v", td.matrix, td.res, res)
		}
	}
}
