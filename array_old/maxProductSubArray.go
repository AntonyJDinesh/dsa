package main

import "fmt"

func main() {
    nums := []int{7, -2, -4}
    fmt.Printf("Res: %d\n", maxProduct(nums))
}


func maxProduct(nums []int) int {
    
    globalMax := nums[0]
    localMax := nums[0]
    var lastNegative *int

    for idx := 1; idx < len(nums); idx++ {
        localMax *= nums[idx]
        
        if localMax < 0  {
            if lastNegative != nil {
                localMax *= *lastNegative
            } else {
                lastNegative = new(int)
                *lastNegative = localMax/nums[idx]
            }
        } 

        if nums[idx] > localMax {
            localMax = nums[idx]
        }

        if localMax > globalMax {
            globalMax = localMax
        }
    }

    return globalMax
}