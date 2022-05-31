package main

import "fmt"

func main() {

    // head = [3,2,0,-4], pos = 1    
    // n4 := &ListNode{Val: -4, Next: nil}
    // n3 := &ListNode{Val: 0, Next: n4}
    n2 := &ListNode{Val: 2, Next: nil}
    n1 := &ListNode{Val: 3, Next: n2}

    // n2.Next = n1
    fmt.Printf("hasCyslse: %t\n", hasCycle(n1))
}

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    slow, fast := head, head.Next
    var cycleFound bool
    for fast != nil {
        if slow == fast {
            cycleFound = true
            break
        }

        slow = slow.Next
        if fast.Next == nil {
            fast = nil
        } else {
            fast = fast.Next.Next
        }
    }

    return cycleFound
}