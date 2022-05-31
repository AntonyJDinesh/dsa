package main

import "fmt"

func main() {
    nums := []int{1,2,0}
    // nums := []int{3,4,-1,1}
    // nums := []int{7,8,9,11,12}
    // nums := []int{1,1}
    
    fmt.Printf("Res: %d\n", firstMissingPositive(nums))
}


func firstMissingPositive(nums []int) int {
    
    n := len(nums)

    for i := 0; i < n; i++ {

        for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }

    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }

    return n+1
}

func firstMissingPositive1(nums []int) int {
    
    ret := 1
    m := make(map[int]struct{}, 0)

    for _, num := range nums {

        if num < 1 {
            continue
        }
        
        m[num] = struct{}{}

        _, exists := m[ret]
        for exists {
            ret++
            _, exists = m[ret]
        }
    }

    return ret
}