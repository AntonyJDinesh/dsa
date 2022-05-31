package main

import "fmt"
func main() {
	s := "loveleetcode"
	fmt.Printf("Res: %d\n", firstUniqChar(s))
}

func firstUniqChar(s string) int {
	nonDupMap := make(map[rune]int, 0)
	for idx, char := range s {
		if _, exists := nonDupMap[char]; exists {
			nonDupMap[char] = -1
		} else {
			nonDupMap[char] = idx
		}
	}

	// fmt.Printf("nonDupMap: %+v\n\n", nonDupMap)

	var smallestIdx *int
	for _, idx := range nonDupMap {
		if idx < 0 {
			continue
		}

		if smallestIdx == nil {
			smallestIdx = new(int)
			*smallestIdx = idx
		} else if *smallestIdx > idx {
			*smallestIdx = idx
		}
	}

	if smallestIdx != nil {
		return *smallestIdx
	}

	return -1
}
