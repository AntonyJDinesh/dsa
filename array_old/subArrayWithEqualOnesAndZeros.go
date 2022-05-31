package main

import "fmt"

func main() {
    nums := []int{0,1,0}
    fmt.Printf("Res: %d\n", findMaxLength(nums))
}

func findMaxLength1(nums []int) int {
    
    maxLen := 0

    for start := 0; start < len(nums); start++ {
        for end := start+1; end <= len(nums); end++ {
            subArray := nums[start : end]
            
            zero, one := 0, 0 
            for _, num := range subArray {
                if num == 0 {
                    zero++
                } else {
                    one++
                }
            }
            if zero == one && ((zero*2) > maxLen) {
                maxLen = 2*zero
            }
        }
    }

    return maxLen
}

func findMaxLength(nums []int) int {
    
    maxLen, count := 0, 0
    countMap := make(map[int]int, 0)
    countMap[0] = -1

    for idx, ele := range nums {
        if ele == 0 {
            count += -1
        } else {
            count += 1
        }

        if lstIdx, exists := countMap[count]; exists {
            len := idx - lstIdx
            if len > maxLen {
                maxLen = len
            }
        } else {
            countMap[count] = idx
        }
    }

    return maxLen
}