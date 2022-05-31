/*
   evelavuate the infix bodmas expression
*/

package stack

import (
	"fmt"
	"os"
	"strconv"
)

func bodmas() {

	if len(os.Args) < 2 {
		fmt.Println("pls pass expression as command line args")
		os.Exit(0)
	}

	expStr := os.Args[1]
	fmt.Printf("Expression: %s\n", expStr[0:len(expStr)])

	tokens := make([]string, 0)
	i, j := 0, 0

	for i < len(expStr) && j < len(expStr) {
		if IsOperator(expStr[j : j+1]) {
			tokens = append(tokens, expStr[i:j])
			tokens = append(tokens, expStr[j:j+1])
			i, j = j+1, j+1
		} else {
			j++
		}
	}

	tokens = append(tokens, expStr[i:len(expStr)])
	fmt.Printf("tokens: %+v\n", tokens)

	numberStack := NewStack()
	operatorStack := NewStack()
	for _, token := range tokens {
		if !IsOperator(token) {
			numberStack.Push(token)
		} else if token == "(" {
			operatorStack.Push(token)
		} else if token == ")" {
		BRACKET:
			for operatorStack.HasNext() {
				operator := operatorStack.Pop()
				if operator == "(" {
					break BRACKET
				}

				rightNuber := numberStack.Pop()
				leftNuber := numberStack.Pop()
				result := Calculate(leftNuber, operator, rightNuber)
				numberStack.Push(result)
			}
		} else if operatorStack.Size() > 0 && operatorStack.Peek() != "(" && FindPrecedence(operatorStack.Peek()) > FindPrecedence(token) {
		BRACKET_2:
			for operatorStack.HasNext() {

				if operatorStack.Peek() == "(" {
					break BRACKET_2
				}
				operator := operatorStack.Pop()
				rightNuber := numberStack.Pop()
				leftNuber := numberStack.Pop()
				result := Calculate(leftNuber, operator, rightNuber)
				numberStack.Push(result)
			}
		} else {
			switch token {
			case "+", "-", "*", "/":
				operatorStack.Push(token)
			}
		}
	}

	for operatorStack.HasNext() {
		operator := operatorStack.Pop()
		rightNuber := numberStack.Pop()
		leftNuber := numberStack.Pop()
		result := Calculate(leftNuber, operator, rightNuber)
		numberStack.Push(result)
	}

	fmt.Println("Result: " + numberStack.Pop())
}

func FindPrecedence(token string) (precedence int) {
	switch token {
	case "+", "-":
		precedence = 1
	case "/", "*":
		precedence = 2
	case "^":
		precedence = 3
	case "(", ")":
		precedence = 4
	default:
		precedence = -1
	}
	return
}

func IsOperator(char string) bool {
	_, err := strconv.Atoi(char)
	return err != nil
}

func Calculate(leftNuber, operator, rightNuber string) (result string) {

	left, _ := strconv.Atoi(leftNuber)
	right, _ := strconv.Atoi(rightNuber)
	switch operator {
	case "+":
		result = strconv.Itoa(left + right)
	case "-":
		result = strconv.Itoa(left - right)
	case "*":
		result = strconv.Itoa(left * right)
	case "/":
		result = strconv.Itoa(left / right)
	}

	return
}
