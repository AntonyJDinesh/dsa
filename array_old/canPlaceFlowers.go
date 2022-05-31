package main

import "fmt"

func main() {
    
    // flowerbed := []int{1,0,1,0,0,0,1}
    // n := 1


    flowerbed := []int{0,0,1,0,1}
    n := 1

    fmt.Printf("Res: %t\n", canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {

    found := 0
    continus := 1
    for _, bed := range flowerbed {

        if bed == 0 {
            continus++
        } else {
            continus = 0
        }

        if continus == 3 {
            found++
            continus = 1
        }

        if found == n {
            break
        }
    }

    if continus == 2 {
        found++
    }

    var ret bool
    if found >= n {
        ret = true
    }

    return ret
}