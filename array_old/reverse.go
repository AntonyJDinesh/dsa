package main

import "fmt"

func main() {
    arr := []string{"h","e","l","l","o"}
    reverseString(arr)
    fmt.Printf("%#v", arr)
}

func reverseString(s []string) {
    for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}