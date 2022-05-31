package main

import "fmt"

func main() {
    // nums := []int{4,2,3} // true
    // nums := []int{4,2,1} // false
    // nums := []int{3,4,2,3} // false
    // nums := []int{1,1,1} // true
    nums := []int{5,7,1,8} // true

    fmt.Printf("Res: %t\n", checkPossibility(nums))
}

func checkPossibility(nums []int) bool {
    
    var check int = 0
    
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            check++

            if i > 1 && nums[i-2] > nums[i] {
                nums[i] = nums[i-1]
            }
        }

        if check > 1 {
            break
        }
    }

    return check < 2
}