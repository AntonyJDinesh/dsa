package string

import "strings"

func rotateString1(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == "" {
		return true
	}

	i, j := 0, 0
	for i < len(s) {
		if s[i] == goal[j] {
			j++
			i++
		} else if j != 0 {
			j = 0
		} else {
			i++
		}
	}

	i = 0
	for j < len(goal) {
		if s[i] == goal[j] {
			j++
			i++
		} else {
			return false
		}
	}

	return true
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	tempS := goal + "#" + s
	lps := genLps1(tempS)
	l := lps[len(tempS)-1]

	return strings.EqualFold(goal[l:], s[0:len(s)-l])
}

func genLps1(s string) []int {
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
