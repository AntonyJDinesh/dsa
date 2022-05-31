package string

func shortestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}

	i := len(s) - 1

	for i > 0 {
		palindrome := true
	inner:
		for j, k := 0, i; j < k; j, k = j+1, k-1 {
			if s[j] != s[k] {
				palindrome = false
				break inner
			}
		}

		if palindrome {
			break
		}
		i--
	}

	palBytes := make([]byte, len(s)+len(s)-i-1)

	inc := 0
	for j := len(s) - 1; j > i; j-- {
		palBytes[inc] = s[j]
		inc++
	}

	for j := 0; j < len(s); j++ {
		palBytes[inc] = s[j]
		inc++
	}

	return string(palBytes)
}

func shortestPalindrome(s string) string {
	revS := reverseStr(s)
	tmpStr := s + "&" + revS

	lps := genLps(tmpStr)
	return revS[0:len(revS)-lps[len(tmpStr)-1]] + s
}

func reverseStr(s string) string {

	revBytes := make([]byte, len(s))
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		revBytes[j] = s[i]
	}

	return string(revBytes)
}

func genLps(s string) []int {

	lps := make([]int, len(s))
	pre, suf := 0, 1
	lps[0] = 0

	for suf < len(s) {
		if s[pre] == s[suf] {
			pre++
			lps[suf] = pre
			suf++
		} else {
			if pre == 0 {
				lps[suf] = 0
				suf++
			} else {
				pre = lps[pre-1]
			}
		}
	}

	return lps
}
