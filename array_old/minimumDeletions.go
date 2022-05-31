package main

import "fmt"

func main() {
    nums := []int{2,10,7,5,4,1,8,6} // 5
    // nums := []int{0,-4,19,1,8,-2,-3,5} // 3
    fmt.Printf("minium deletions: %d\n", minimumDeletions(nums))
}

func minimumDeletions(nums []int) int {
    min, minIndex, max, maxIndex := nums[0], 0, nums[0], 0

    for i := 1; i < len(nums); i++ {
        if nums[i] < min {
            min, minIndex = nums[i], i
        }

        if nums[i] > max {
            max, maxIndex = nums[i], i
        }
    }

    start, end := 0, 0
    if minIndex < maxIndex {
        start, end = minIndex, maxIndex
    } else {
        start, end = maxIndex, minIndex
    }

    
    minDeletions := 0
    removeFromFront := end + 1

    if minDeletions < removeFromFront {
        minDeletions = removeFromFront
    }

    removeFromEnd := len(nums) - start
    if removeFromEnd < minDeletions {
        minDeletions = removeFromEnd
    }

    removeFromBothEnd := (start + 1) + (len(nums) - end)
    if removeFromBothEnd < minDeletions {
        minDeletions = removeFromBothEnd
    }


    return minDeletions
}