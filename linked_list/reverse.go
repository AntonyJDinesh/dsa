package main

import (
    "fmt"
    "strings"
    "strconv"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func (listNode *ListNode) String() string {
    str := ""

    for node := listNode; node != nil; node = node.Next {
        str += strconv.Itoa(node.Val) + "->"
    }

    return strings.TrimSuffix(str, "->")
}

func main() {
    five := &ListNode{Val: 5, Next: nil}
    four := &ListNode{Val: 4, Next: five}
    three := &ListNode{Val: 3, Next: four}
    two := &ListNode{Val: 2, Next: three}
    one := &ListNode{Val: 1, Next: two}

    fmt.Printf("List: %s\n", one)
    fmt.Printf("Reverse List: %s\n", reverseList(one))
    
}

func reverseList(head *ListNode) *ListNode {

    	// 1 -> 2 -> 3 -> 4
	// r
	// 1 <- 2 -> 3 -> 4
	//      r
	// 1 <- 2 <- 3 -> 4
	//           r
	// 1 <- 2 <- 3 <- 4
	//				  r

	a := head
	if a == nil {
		return a
	}

	b := a.Next
	if b == nil {
		return a
	}

	c := b.Next
	a.Next = nil

	for c != nil {
		b.Next = a
		a = b
		b = c
		c = c.Next
	}

	b.Next = a

	return b
    // return recursiveReverse(head, nil)
}

func recursiveReverse(node, next *ListNode) *ListNode {
    if node == nil {
        return next
    }
    tmp := node.Next
    node.Next = next
    return recursiveReverse(tmp, node)
}  