package dp

import (
	"sort"
)

func deleteAndEarn(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	points := make(map[int]int, 0)
	uniqueNums := make([]int, 0)
	for _, num := range nums {
		if _, found := points[num]; found {
			points[num] += num
		} else {
			points[num] = num
			uniqueNums = append(uniqueNums, num)
		}
	}

	sort.Ints(uniqueNums)
	// fmt.Printf("points: %#v uniqueNums: %#v\n", points, uniqueNums)
	oneDown, twoDown := 0, 0
	for i := 0; i < len(uniqueNums); i++ {
		if i > 0 && uniqueNums[i-1] == uniqueNums[i]-1 {

			oneDown, twoDown = max(oneDown, twoDown+points[uniqueNums[i]]), oneDown
			// fmt.Printf("if: oneDown: %d, twoDown: %d\n", oneDown, twoDown)
		} else {
			oneDown, twoDown = oneDown+points[uniqueNums[i]], oneDown
		}
	}

	return oneDown
}

func max(val1, val2 int) int {
	// fmt.Printf("oneDown: %d, twoDown: %d\n", val1, val2)
	if val1 > val2 {
		return val1
	}

	return val2
}
