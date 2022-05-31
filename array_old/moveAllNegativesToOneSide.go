package main

import "fmt"

func main() {
    arr := []int{-2,3,-1,4,5,-1,4, -1}
    moveAllNegativesToOneSide(arr)
    fmt.Printf("%#v\n", arr)
}

func moveAllNegativesToOneSide(arr []int) {
    start, end := 0, len(arr)-1
    for start < end {
        if arr[start] >= 0 { // find negative
            start++
        } else if arr[end] < 0 {
            end--
        } else {
            arr[start], arr[end] = arr[end], arr[start]
        }
    }
}