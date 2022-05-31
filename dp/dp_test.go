package dp

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	testData := []struct {
		cost   []int
		result int
	}{
		{
			cost:   nil,
			result: 0,
		},
		{
			cost:   []int{1},
			result: 1,
		},
		{
			cost:   []int{1, 1},
			result: 1,
		},
		{
			cost:   []int{2, 1},
			result: 1,
		},
		{
			cost:   []int{1, 2},
			result: 1,
		},
		{
			cost:   []int{1, 1, 1},
			result: 1,
		},
		{
			cost:   []int{1, 3, 1},
			result: 2,
		},
		{
			cost:   []int{1, 3, 1, 2},
			result: 2,
		},
		{
			cost:   []int{1, 3, 1, 2, 5},
			result: 4,
		},
	}

	for _, td := range testData {
		res := minCostClimbingStairs(td.cost)
		if res != td.result {
			t.Errorf("Failed: %#v, Expected: %d, Got: %d", td.cost, td.result, res)
		}
	}
}

func Test_tribonacci(t *testing.T) {
	testData := []struct {
		n   int
		res int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 7},
	}

	for _, td := range testData {
		res := tribonacci(td.n)
		if res != td.res {
			t.Errorf("n: %d, exp: %d, got: %d", td.n, td.res, res)
		}
	}
}

func Test_deleteAndEarn(t *testing.T) {
	testData := []struct {
		nums []int
		res  int
	}{
		{nil, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3}, 4},
		{[]int{1, 2, 3, 2, 2}, 6},
	}

	for _, td := range testData {
		res := deleteAndEarn(td.nums)
		if res != td.res {
			t.Errorf("%#v ==> Exp: %d, Got: %d", td.nums, td.res, res)
		}
	}
}
