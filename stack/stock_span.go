/*
The stock span problem is a financial problem where we have a series of n daily price quotes for a stock and we need to calculate span of stock’s price for all n days.
The span Si of the stock’s price on a given day i is defined as the maximum number of consecutive days just before the given day, for which the price of the stock on the current day is less than its price on the given day.
For example, if an array_old of 7 days prices is given as {100, 80, 60, 70, 60, 75, 85}, then the span values for corresponding 7 days are {1, 1, 1, 2, 1, 4, 6}
*/

package stack

import (
	"fmt"
)

func stockSpan() {

	prices := []int{100, 80, 60, 70, 60, 75, 85} // ans: {1, 1, 1, 2, 1, 4, 6}
	spans := make([]int, len(prices))

	stk := NewIntStack()
	i := len(prices) - 1
	for i > -1 {

		if stk.Size() == 0 {
			stk.Push(i)
		} else if prices[stk.Peek()] > prices[i] {
			stk.Push(i)
		} else {
			tpDay := i
			for stk.Size() > 0 && prices[stk.Peek()] < prices[i] {
				day := stk.Pop()
				spans[day] = day - tpDay
			}
			stk.Push(i)
		}

		i--
	}

	tpDay := i
	if stk.Size() > 0 {
		for stk.Size() > 0 {
			day := stk.Pop()
			spans[day] = day - tpDay
		}
	}

	fmt.Printf("prices: %#v\nspans: %#v", prices, spans)
}
