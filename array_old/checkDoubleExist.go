package main

import (
    "fmt"
)

func main() {
    // inp := []int{2,4} // true
    // inp := []int{4,2} // true
    inp := []int{3,1,7,11} // false
    fmt.Printf("%t", checkIfExist(inp))
}

func checkIfExist(arr []int) bool {
    m := make(map[int]struct{}, 0)
    for _, val := range arr {
        if _, exists := m[val*2]; exists {
            return true
        } else if _, exists := m[val/2]; val%2 ==0 && exists {
            return true
        } else {
            m[val] = struct{}{}
        }
    }

    return false
}