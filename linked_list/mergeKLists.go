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

    // [[1,4,5],[1,3,4],[2,6]]
    l1n3 := &ListNode{Val: 5, Next: nil}
    l1n2 := &ListNode{Val: 4, Next: l1n3}
    l1n1 := &ListNode{Val: 1, Next: l1n2}

    l2n3 := &ListNode{Val: 4, Next: nil}
    l2n2 := &ListNode{Val: 3, Next: l2n3}
    l2n1 := &ListNode{Val: 1, Next: l2n2}

    
    l3n2 := &ListNode{Val: 6, Next: nil}
    l3n1 := &ListNode{Val: 2, Next: l3n2}

    fmt.Printf("Reverse List: %s\n", mergeKLists([]*ListNode{l1n1, l2n1, l3n1}))
    
}


func mergeKLists(lists []*ListNode) *ListNode {
    var head, curr *ListNode

    for true {
        smaller := lists[0]
        smallerList := 0
        for i := 1; i < len(lists); i++ {
            if smaller == nil || (lists[i] != nil && smaller.Val > lists[i].Val) {
                smaller = lists[i]
                smallerList = i
            }
        }

        if smaller == nil {
            break
        } else {
            lists[smallerList] = lists[smallerList].Next
            newNode := &ListNode{Val: smaller.Val}

            if head == nil {
                curr = newNode
                head = curr
            } else {
                curr.Next = newNode
                curr = newNode
            }
        }
    }

    return head
}