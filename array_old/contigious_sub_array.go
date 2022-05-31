package main

import (
    "fmt"
)

func main() {
    inputArray := []int{-2,1,-3,4,-1,2,1,-5,4}

    // Kadane's Algorithm

    var localMax, globalMax *int
    var localStartIndex, localEndIndex, globalStartIndex, globalEndIndex int
    for i, numberAtI := range inputArray {
         
        if localMax == nil {
            localMax = new(int)
            *localMax = numberAtI
            localStartIndex = i
        } else if *localMax + numberAtI < numberAtI {
            *localMax = numberAtI
            localStartIndex = i
            localEndIndex = i
        } else {
            localEndIndex = i
            *localMax = *localMax + numberAtI
        }
           
        if globalMax == nil || *globalMax < *localMax {
            if globalMax == nil {
                globalMax = new(int)
            }

            *globalMax = *localMax
            globalStartIndex = localStartIndex
            globalEndIndex = localEndIndex
        }

       //  fmt.Printf("globalMax: %d, localMax: %d\n", *globalMax, *localMax) 
    }

    fmt.Printf("globalMax: %d, globalStartIndex: %d, globalEndIndex: %d, subArray: %+v\n", *globalMax, globalStartIndex, globalEndIndex, inputArray[globalStartIndex:globalEndIndex+1])
}