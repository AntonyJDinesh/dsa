package main

import "fmt"

func main() {
    // nums := []int{100,4,200,1,3,2}
    nums := []int{0,3,7,2,5,8,4,6,0,1}
    fmt.Printf("Res: %d\n", longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
    
    m := make(map[int]struct{}, 0)
    for _, num := range nums {
        m[num] = struct{}{}
    }

    globalLC := 0
    for _, num := range nums {

        if _, exists := m[num-1]; !exists {
            
            delete(m, num)
            localLC := 1
            num++
            
            for _, nextNumFound := m[num]; nextNumFound; _, nextNumFound = m[num] {
                localLC++

                delete(m, num)
                num++
            }

            if localLC > globalLC {
                globalLC = localLC
            }
        }
    }

    return globalLC 
}