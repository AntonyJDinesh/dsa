package main

import "fmt"

func main() {
     // head = [3,2,0,-4], pos = 1    
    n4 := &ListNode{Val: -4, Next: nil}
    n3 := &ListNode{Val: 0, Next: n4}
    n2 := &ListNode{Val: 2, Next: n3}
    n1 := &ListNode{Val: 3, Next: n2}
    n4.Next = n2
    ent := detectCycle(n1)

    fmt.Printf("Ent Node: %s\n", ent)
}

type ListNode struct {
    Val int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    slow, fast := head, head
    var cycleFound bool
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            cycleFound = true
            break
        }
    }
    
    if cycleFound {
        slow = head
        for slow != fast {
            slow = slow.Next
            fast = fast.Next
        }
        
        
        return fast
    }
    return nil
}