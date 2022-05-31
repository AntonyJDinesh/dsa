package bs

import (
	"fmt"
	"math"
)

func findClosestElements(arr []int, k int, x int) []int {

	left, right := 0, len(arr)-1
	for left <= right {
		p := left + int((right-left)/2)
		if arr[p] == x {
			left = p
			break
		} else if arr[p] > x {
			right = p - 1
		} else {
			left = p + 1
		}
	}

	right = left + 1
	fmt.Printf("left: %d, right: %d\n", left, right)
	for i := 1; i < k; i++ {

		if left == 0 {
			right += k - i
			break
		} else if right == len(arr) {
			left -= k - i
			break
		}

		ret := closerOne(x, arr[left], arr[right])
		if ret > 0 {
			left--
		} else {
			right++
		}
		fmt.Printf("left: %d, right: %d\n", left, right)
	}

	return arr[left:right]
}

func closerOne(x, left, right int) int {
	if math.Abs(float64(x)-float64(left)) <= math.Abs(float64(x)-float64(right)) {
		return 1
	}

	return -1
}
