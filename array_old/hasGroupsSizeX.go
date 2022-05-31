package main

import "fmt"

func main() {
    // inp := []int{1,2,3,4,4,3,2,1}
    // inp := []int{1,1,1,2,2,2,3,3}
    // inp := []int{1}
    inp := []int{1,1,2,2,2,2}

    fmt.Printf("Res: %t", hasGroupsSizeX(inp))
}

func hasGroupsSizeX(deck []int) bool {
    
    groupMap := make(map[int]int, 0)
    groupSize := 1 
    groupSizeCard := deck[0]

    for _, card := range deck {
        if _, exists := groupMap[card]; exists {
            groupMap[card]++
        } else {
            groupMap[card] = 1
        }

        if card == groupSizeCard {
            groupSize = groupMap[card]
        } else if card != groupSizeCard && groupSize > groupMap[card] {
            card = groupSizeCard
            groupSize = groupMap[card]
        }
    }
    
    for i :=  groupSize; i >= 2; i-- {
        ret := true
        for _, val := range groupMap {
            if val%i != 0 {
                ret = false
                break
            }
        }

        if ret {
            return true
        }
    }

    return false
}