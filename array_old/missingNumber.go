package main

import "fmt"

func main() {
    nums := []int{9,6,4,2,3,5,7,0,1}
    fmt.Printf("Missing: %d\n", missingNumber(nums))
    fmt.Printf("Missing: %d\n", missingNumber1(nums))
    fmt.Printf("Missing: %d\n", missingNumber2(nums))
}

func missingNumber1(nums []int) int {
    i := 0
    for i < len(nums) {
        if i != nums[i] && nums[i] < len(nums) {
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        } else {
            i++
        }
    }

    res := len(nums) 
    for pos, num := range nums {
        if pos != num {
            res = pos
            break
        }
    }


    return res
}


// 0, 1, 2 = 3
// 0, 2, 3 = 5
// n = 3


func missingNumber2(nums []int) int {
    res := len(nums)
    for pos, num := range nums {
        res += pos - num 
    }
    return res
}

func missingNumber(nums []int) int {
    res := len(nums)
    for pos, num := range nums {
        res = res ^ pos ^ num 
    }

    return res
}