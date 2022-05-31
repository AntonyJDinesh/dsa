package main

import "fmt"

func main() {
    height := []int{1,8,6,2,5,4,8,3,7}
    // height := []int{1,2,4,3}
    // height := []int{1,0,0,0,0,0,0,2,2}
    fmt.Printf("Result: %d\n", maxArea(height))
}

func maxArea(height []int) int {
    
    if len(height) < 2 {
        return 0
    }

    start, end := 0, len(height)-1
    globalMax := 0 
    for start <= end {
        if height[start] == 0 {
            start++
        } 

        if height[end] == 0 {
            end--
        }

        capacity := 0
        if height[start] < height[end] {
            capacity = (end-start)*height[start]
            start++
        } else {
            capacity = (end-start)*height[end]
            end--
        }

        if capacity > globalMax {
            globalMax = capacity
        }
    }

    return globalMax
}


func maxArea1(height []int) int {
    
    if len(height) < 2 {
        return 0
    }

    globalMax, localMax := 0, 0
    globalStart, globalEnd, localStart := 0, 0, 0
    for i := 1; i < len(height); i++ {

        if height[i] == 0 {
            continue
        }

        capacity := height[i-1]
        if height[i] < height[i-1] {
            capacity = height[i]
        }

        if height[localStart] < height[i] {
            localMax = height[localStart] * (i-localStart)
        } else {
            localMax = height[i] * (i-localStart)
        }

        if capacity >= localMax {
            localMax = capacity
            localStart = i-1
        }
        

        fmt.Printf("localStart: %d, localMax: %d, globalMax: %d", localStart, localMax, globalMax)
        if localMax > globalMax {
            globalMax = localMax
            globalStart = localStart
            globalEnd = i
        }
        fmt.Printf(", globalMax: %d\n", globalMax)
    }

    fmt.Printf("globalStart: %d, globalEnd: %d\n", globalStart, globalEnd)

    return globalMax
}