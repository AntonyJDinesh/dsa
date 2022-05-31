package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
     // head = [3,2,0,-4], pos = 1    
    n4 := &ListNode{Val: -4, Next: nil}
    n3 := &ListNode{Val: 0, Next: n4}
    n2 := &ListNode{Val: 2, Next: n3}
    n1 := &ListNode{Val: 3, Next: n2}
    n4.Next = n2
    removeCycle(n1)

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

func removeCycle(head *ListNode) {
    if head == nil {
        return
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
        
        slow = head
        entry := true
        for slow.Next != fast || entry {
            if slow.Next == fast {
                entry = false
            }

            slow = slow.Next
        }

        slow.Next = nil
    }
}