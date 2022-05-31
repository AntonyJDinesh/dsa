package main

import "fmt"

func main() {
    nums := []int{1,2,3,1}
    fmt.Printf("Peek Element Index: %d\n", findPeakElement(nums))
}

func findPeakElement(nums []int) int {
    
    peek := nums[0]
    peekIndex := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > peek {
            peek = nums[i]
            peekIndex = i
        }
    }

    return peekIndex
}