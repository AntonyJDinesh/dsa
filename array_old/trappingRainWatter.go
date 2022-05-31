package main

import "fmt"

func main() {
    // height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
    height := []int{4,2,3}
    fmt.Printf("Result: %d\n", trap(height))
}

func trap(height []int) int {
    
    trapped := 0
    if len(height) < 3 {
        return trapped
    }

    left, right, leftMax, rightMax := 0, len(height)-1, 0, 0

    for left <= right {
        if height[left] < height[right] {

            if leftMax < height[left] {
                leftMax = height[left]
            } else {
                trapped += leftMax - height[left]
            }
            left++
        } else {
            if rightMax < height[right] {
                rightMax = height[right]
            } else {
                trapped += rightMax - height[right]
            }
            right--
        }
    }

    return trapped
}