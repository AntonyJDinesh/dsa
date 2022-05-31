/*
   For a given array_old, identify the nearest smaller in the left of each element in the array_old. if no smaller element found then '-'
   Example:

   input: [1, 5, 3, 6, 7, 2]
   oputput: [-, 1, 1, 3, 6, 1]
*/

package stack

import "fmt"

func nearestSmaller() {
	input := []int{1, 5, 3, 6, 7, 2}
	output := make([]interface{}, len(input))

	tmpStack := NewIntStack()
	for i, number := range input {

		for tmpStack.Size() > 0 && tmpStack.Peek() > number {
			tmpStack.Pop()
		}

		if tmpStack.Size() == 0 {
			// ououtput: []interface {}{"-", 1, 1, 3, 6, 1}tput[i] = "-"
		} else {
			output[i] = tmpStack.Peek()
		}

		tmpStack.Push(number)
	}

	fmt.Printf("input: %#v\noutput: %#v\n", input, output)
}
