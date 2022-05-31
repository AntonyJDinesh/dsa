package string

import (
	"reflect"
	"testing"
)

func Test_patternMatching(t *testing.T) {
	testData := []struct {
		str, pat string
		res      []int
	}{
		{"abc", "abc", []int{0}},
		{"1abcdab", "abc", []int{1}},
	}

	for _, td := range testData {
		res := patternMatching(td.str, td.pat)
		if !reflect.DeepEqual(res, td.res) {
			t.Errorf("Str: %s, Pat: %s, Exp: %#v, Got: %#v", td.str, td.pat, td.res, res)
		}
	}
}

func Test_addBinary(t *testing.T) {
	testData := []struct {
		a, b, res string
	}{
		{"1", "1", "10"},
		{"0", "0", "0"},
		{"1111", "", "1111"},
		{"", "1010", "1010"},
		{"11", "1", "100"},
		{"11", "10", "101"},
		{"11", "11", "110"},
		{"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			"110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}

	for _, td := range testData {
		res := addBinary(td.a, td.b)
		if res != td.res {
			t.Errorf("a: %s, b: %s, Exp:%s, Got: %s", td.a, td.b, td.res, res)
		}
	}
}

func Test_rotateString(t *testing.T) {
	testData := []struct {
		s, goal string
		ret     bool
	}{
		{"", "", true},
		{"", "a", false},
		{"a", "", false},
		{"aa", "a", false},
		{"aa", "aaa", false},
		{"aaa", "aaa", true},
		{"abc", "abc", true},
		{"abc", "bca", true},
		{"abc", "cab", true},
		{"abc", "abd", false},
		{"abc", "bda", false},
		{"abc", "bcd", false},
		{"abc", "bcd", false},
	}

	for _, td := range testData {
		res := rotateString(td.s, td.goal)
		if res != td.ret {
			t.Errorf("s: %s, goal: %s, exp: %t, got: %t", td.s, td.goal, td.ret, res)
		}
	}
}

func Test_shortestPalindrome(t *testing.T) {
	testData := []struct {
		s, res string
	}{
		{"", ""},
		{"a", "a"},
		{"aa", "aa"},
		{"ab", "bab"},
		{"abc", "cbabc"},
	}

	for _, td := range testData {
		res := shortestPalindrome(td.s)
		if res != td.res {
			t.Errorf("s: %s, exp: %s, got: %s", td.s, td.res, res)
		}
	}
}
