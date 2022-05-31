package array

func pivotIndex(nums []int) int {
	var sum, leftSum int
	for _, num := range nums {
		sum += num
	}

	for i := 0; i < len(nums); i++ {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
