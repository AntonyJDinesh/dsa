package array

func findKthLargest(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || len(nums) < k {
		return 0
	}

	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, start, end, k int) int {

	i, j := start, start
	for j < end {
		if nums[j] < nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]

	var kth int
	if (len(nums)-1)-(k-1) == i {
		kth = nums[i]
	} else if (len(nums)-1)-(k-1) > i {
		kth = quickSelect(nums, i+1, end, k)
	} else {
		kth = quickSelect(nums, start, i-1, k)
	}

	return kth
}
