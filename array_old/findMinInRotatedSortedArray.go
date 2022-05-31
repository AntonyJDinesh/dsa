package main

import "fmt"

func main() {
	nums := []int{2, 1}
	fmt.Printf("Min: %d\n", findMin(nums))
}

func findMin(nums []int) int {

	if nums == nil {
		return -1
	}

	if len(nums) == 1 || nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	start, end := 0, len(nums)-1

	for start < end {

		p := start + int((end-start)/2)

		if p > 0 && nums[p] < nums[p-1] {
			return nums[p]
		}
		if p < len(nums)-1 && nums[p+1] < nums[p] {
			return nums[p+1]
		}

		if nums[0] < nums[p] {
			start = p + 1
		} else {
			end = p - 1
		}
	}

	return nums[start]
}
