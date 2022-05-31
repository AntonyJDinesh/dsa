package find_and_union

import "testing"

func Test_findCircleNum(t *testing.T) {

	testData := []struct {
		connected [][]int
		res       int
	}{
		{
			[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2,
		},
	}

	for _, td := range testData {
		res := findCircleNum(td.connected)
		if res != td.res {
			t.Errorf("%#v Exp: %d, Got: %d", td.connected, td.res, res)
		}
	}
}
