package string

func patternMatching(str, pat string) []int {
	res := make([]int, 0)
	strLen := len(str)
	patLen := len(pat)

	i, j := 0, 0
	for i < strLen {
		if str[i] == pat[j] {
			if j+1 == patLen {
				res = append(res, i-j)
				i++
				j = 0
			} else {
				i++
				j++
			}
		} else {
			i = i - j + 1
			j = 0
		}
	}

	// abcabd
	// abd

	// 0, 0; 1; 1; 1,0; 2,0, 3,0, 4,1, 5,2 , i-j

	return res
}
