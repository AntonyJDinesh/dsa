package main

import "fmt"

func main() {
    // nums := []int{1,2,10,5,7}
    // nums := []int{2,3,1,2}
    // nums := []int{1, 1, 1}
    // nums := []int{105,924,32,968}
    nums := []int{1,4,1,2,3}
    fmt.Println(canBeIncreasing(nums))
}

func canBeIncreasing(nums []int) bool {
    
    var flag bool 

    for i := 1; i < len(nums); i++ {
        if nums[i] <= nums[i-1] {
            if flag {
                return false
            }
            flag = true

            if i > 1 && nums[i-2] >= nums[i] {
                nums[i] = nums[i-1]
            }
        }
    }

    return true
}