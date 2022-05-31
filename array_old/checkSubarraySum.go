package main

import "fmt"

func main() {
    // nums := []int{23,2,4,6,7} 
    // k := 6

    // nums := []int{23,2,6,4,7} 
    // k := 6

    // nums := []int{23,2,6,4,7} 
    // k := 13

    nums := []int{23,2,6,4,7} 
    k := 7
    
    
    fmt.Printf("Res: %t\n", checkSubarraySum(nums, k))
}

func checkSubarraySum(nums []int, k int) bool {
    m := make(map[int]int, 0)
    m[0] = -1
    mod := 0
    for pos, num := range nums {
        mod = (mod + num)%k
        if idx, exists := m[mod]; exists {
            if pos - idx > 1 {
                return true
            }
            
        } else {
            m[mod] = pos
        }
    }

    return false
}