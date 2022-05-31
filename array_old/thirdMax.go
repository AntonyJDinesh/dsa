package main

import "fmt"

func main() {
    input := []int{2,2,3,1}
    // input := []int{3,2,1}
    fmt.Printf("Result: %d\n\n", thirdMax(input))
}

func thirdMax(nums []int) int {
    
    var tMax, sMax, max *int
     dupMap := make(map[int]struct{}, 0)

    for i:=0; i<len(nums); i++ {

        if _, exists :=  dupMap[nums[i]]; exists {
            continue
        } else {
            dupMap[nums[i]] = struct{}{}
        }

        if tMax == nil {
            tMax = new(int)
            *tMax = nums[i]
        } else if *tMax < nums[i] {
            *tMax = nums[i]
        }

        if sMax == nil || *tMax > *sMax {
            sMax, tMax = tMax, sMax
        }

        if max == nil || *sMax > *max {
            sMax, max = max, sMax
        }

        fmt.Printf("tMax: %v, sMax: %v, max: %v\n", tMax, sMax, max)
    }

    var ret *int
    if  tMax != nil {
        ret = tMax
    } else {
        ret = max
    }

    return *ret
}