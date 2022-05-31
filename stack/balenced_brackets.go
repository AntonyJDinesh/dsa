package stack

import (
	"fmt"
)

func isBalanced(s string) (res string) {
	res = "YES"
	operatorStack := NewStack()
	for _, t := range s {
		token := fmt.Sprintf("%c", t)
		switch token {
		case "(", "[", "{":
			operatorStack.Push(token)
		case ")":
			if operatorStack.Size() > 0 && operatorStack.Peek() == "(" {
				operatorStack.Pop()
			} else {
				operatorStack.Push(token)
			}
		case "}":
			if operatorStack.Size() > 0 && operatorStack.Peek() == "{" {
				operatorStack.Pop()
			} else {
				operatorStack.Push(token)
			}
		case "]":
			if operatorStack.Size() > 0 && operatorStack.Peek() == "[" {
				operatorStack.Pop()
			} else {
				operatorStack.Push(token)
			}
		}
	}

	if operatorStack.Size() > 0 {
		res = "NO"
	}

	return
}
