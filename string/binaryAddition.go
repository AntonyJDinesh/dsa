package string

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}

	var c int
	var aArr, bArr, result []int

	/*ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)

	fmt.Printf("ai: %d, bi: %d\n", ai, bi)

	if ai == 0 {
		aArr = append(aArr, 0)
	} else {
		for ai > 0 {
			aArr = append([]int{ai % 10}, aArr...)
			ai = ai / 10
		}
	}

	if bi == 0 {
		bArr = append(bArr, 0)
	} else {
		for bi > 0 {
			bArr = append([]int{bi % 10}, bArr...)
			bi = bi / 10
		}
	}*/

	for _, r := range a {
		aArr = append(aArr, int(r)-'0')
	}

	for _, r := range b {
		bArr = append(bArr, int(r)-'0')
	}

	fmt.Printf("aArr: %#v, bArr: %#v\n", aArr, bArr)

	i, j := len(aArr)-1, len(bArr)-1
	for i >= 0 || j >= 0 {

		var t int
		if i >= 0 && j >= 0 {
			t = aArr[i] + bArr[j] + c
		} else if i >= 0 {
			t = aArr[i] + c
		} else if j >= 0 {
			t = bArr[j] + c
		}

		switch t {
		case 0:
			result = append([]int{0}, result...)
			c = 0
		case 1:
			result = append([]int{1}, result...)
			c = 0
		case 2:
			result = append([]int{0}, result...)
			c = 1
		case 3:
			result = append([]int{1}, result...)
			c = 1
		}

		i, j = i-1, j-1
	}

	if c > 0 {
		result = append([]int{1}, result...)
	}

	res := ""
	for _, b := range result {
		res += strconv.Itoa(b)
	}

	return res
}
