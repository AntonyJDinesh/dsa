package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
     // head = [3,2,0,-4], pos = 1    
    n4 := &ListNode{Val: 3, Next: nil}
    n3 := &ListNode{Val: 3, Next: n4}
    n2 := &ListNode{Val: 3, Next: n3}
    n1 := &ListNode{Val: 1, Next: n2}
    deleteDuplicates(n1)

    fmt.Printf("List: %s\n", n1)
}

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

func deleteDuplicates(head *ListNode) *ListNode {
    node := head
    for node != nil {
        fmt.Printf("Node: %d\n", node.Val)
        if node.Next != nil && node.Val == node.Next.Val {
            node.Next = node.Next.Next
        } else {
            node = node.Next
        }
    }

    return head
}