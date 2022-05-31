package main

import "fmt"

func main() {
    // inp := []int{0,3,2,1} // true
    // inp := []int{2,1} // false
    // inp := []int{3,5,5} // false
    // inp := []int{0,2,3,3,5,2,1,0} // false
    inp := []int{1,3,2} // true
    // inp := []int{0,1,2,3,4,5,6,7,8,9} // false
    // inp := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0} // false
    fmt.Printf("Res: %t\n", validMountainArray(inp))
}

func validMountainArray(arr []int) bool {
    
    if len(arr) < 3 {
        return false
    }

    up := arr[0] < arr[1]
    down := false
    top := false
    
    for i:=1; i<len(arr); i++ {
        if arr[i-1] == arr[i] {
            return false
        }

        if !top {
            if arr[i] < arr[i-1] {
                top = true
                down = true
            } else {
                up = true   
            }
        } else {
            if arr[i] > arr[i-1] {
                return false
            }
        }
    }

    return up && down
}