package main

import "fmt"

func main() {
    nums := []int{1,3,4,2,2}
    fmt.Printf("Duplicate: %d\n", findDuplicate(nums))
}

func findDuplicate(nums []int) int {
    
    slow, fast := nums[0], nums[nums[0]]
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }

    slow = 0
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }

    return slow
}