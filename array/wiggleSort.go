package array

func wiggleSort(nums []int) {
	quickWiggleSort(nums, 0, len(nums)-1)
}

func quickWiggleSort(nums []int, start, end int) {

	if start >= end || end-start < 2 {
		return
	}

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

	nums[i-1], nums[end] = nums[end], nums[i-1]

	quickWiggleSort(nums, start, i-1)
	quickWiggleSort(nums, i+1, end)
}
