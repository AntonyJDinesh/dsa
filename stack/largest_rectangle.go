/*
   Find the largest rectangle that can be formed within the bounds of consecutive buildings.
*/

package stack

/*
 * Complete the 'largestRectangle' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY h as parameter.
 */

func largestRectangle(h []int32) int64 {
	// Write your code here
	size := int32(len(h))
	var max_area *int64 = new(int64)
	for i := int32(0); i < size; i++ {
		j := i
	inner1:
		for j+1 < size {
			if h[j+1] < h[i] {
				break inner1
			} else {
				j++
			}
		}

		k := i
	inner2:
		for k-1 > -1 {
			if h[k-1] < h[i] {
				break inner2
			} else {
				k--
			}
		}

		area := int64(h[i] * (j - k + 1))
		if max_area == nil || *max_area < area {
			*max_area = area
		}
	}

	return *max_area
}

func largestRectangle2(h []int32) int64 {
	var max_area int64
	stk := NewIntStack()

	for position, element := range h {
		if stk.Size() == 0 {
			stk.Push(position)
		} else if h[stk.Peek()] <= element {
			stk.Push(position)
		} else {
			for {
				topPosition := stk.Pop()
				if h[topPosition] <= element {
					break
				}

				area := int64(int32(position-topPosition) * h[topPosition])
				if area > max_area {
					max_area = area
				}
			}
		}
	}

	position := len(h) - 1
	for {
		topPosition := stk.Pop()
		area := int64(int32(position-topPosition) * h[topPosition])
		if area > max_area {
			max_area = area
		}
	}

	return max_area
}
