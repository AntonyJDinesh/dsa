package main

import "fmt"

func main() {

    nums := []int{10, 5, 3, 4, 3, 5, 6}
    fmt.Printf("Res: %d\n", firstRepeating(nums)
}

func firstRepeating(nums []int) int {
    dupMapCheck := make(map[int]int, 0)
    firstCheckMap := make(map[int]int, 0)

    for pos, num := range nums {
        if idx, exists := dupMapCheck[num]; exists {
            firstCheckMap[num] = idx
        } else {
            dupMapCheck[num] = pos
        }
    }

    var smallestIdx, number int
    for num, idx := range firstCheckMap {
        if smallestIdx > idx {
            number = num
            smallestIdx = idx
        }
    }

    return number
}