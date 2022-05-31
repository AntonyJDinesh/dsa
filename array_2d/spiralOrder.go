package array_2d

import "fmt"

func spiralOrder(matrix [][]int) []int {

	type Direction int
	const (
		LeftToRight Direction = iota
		TopToBottom
		RightToLeft
		BottomToTop
	)

	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	res := make([]int, 0)
	i, j := 0, 0
	left, right := 0, len(matrix[0])-1
	top, bottom := 0, len(matrix)-1
	direction := LeftToRight

	/*
		1 2 3
		4 5 6
		7 8 9

		top: 1, right: 1, bottom: 1, left: 1
	*/
	for left <= right && top <= bottom {

		fmt.Printf("d: %d, i: %d, j: %d\n", direction, i, j)
		res = append(res, matrix[i][j])

		switch direction {
		case LeftToRight:
			if j < right {
				j++
			} else {
				direction = TopToBottom
				top++
				i++
			}
		case TopToBottom:
			if i < bottom {
				i++
			} else {
				direction = RightToLeft
				right--
				j--
			}
		case RightToLeft:
			if j > left {
				j--
			} else {
				direction = BottomToTop
				bottom--
				i--
			}
		case BottomToTop:
			if i > top {
				i--
			} else {
				direction = LeftToRight
				left++
				j++
			}
		}
	}

	return res
}
