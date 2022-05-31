package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Printf("Min: %d\n", searchRange(nums, 8))
}

func searchRange(nums []int, target int) []int {
	if nums == nil {
		return []int{-1, -1}
	}

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	start, end := 0, len(nums)-1
	var pivot int
	for start <= end {
		pivot = start + int((end-start)/2)

		fmt.Printf("start: %d, end: %d, p: %d\n", start, end, pivot)
		if nums[pivot] == target {
			break
		} else if nums[pivot] > target {
			end = pivot - 1
		} else {
			start = pivot + 1
		}
	}

	if nums[pivot] == target {
		start, end = pivot, pivot
		for true {
			if start > 0 && nums[start-1] == target {
				start = start - 1
			} else if end < len(nums)-1 && nums[end+1] == target {
				end = end + 1
			} else {
				return []int{start, end}
			}
		}
	}

	return []int{-1, -1}
}
