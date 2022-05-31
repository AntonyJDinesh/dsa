package main

import (
    "fmt"
)

func main() {
    prices := []int{7,1,5,3,6,4}
    fmt.Printf("Max profit: %d\n", maxProfit(prices))
}

func maxProfit(prices []int) int {
    
    var globalMaxProfit, localMaxProfit int
    var localLowPrice *int
    
    for _, price := range prices {

        if localLowPrice == nil {
            localLowPrice = new(int)
            *localLowPrice = price
        }

        profit := price - *localLowPrice
        if profit > localMaxProfit {
            localMaxProfit = profit
        }

        if *localLowPrice > price {
            *localLowPrice = price
        }

        if globalMaxProfit < localMaxProfit {
            globalMaxProfit = localMaxProfit
        }
    }


    return globalMaxProfit
}