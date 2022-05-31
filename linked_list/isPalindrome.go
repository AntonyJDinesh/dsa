package main

import (
    "fmt"
    "strconv"
    "strings"
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

    // [1,0,3,4,0,1]
    n6 := &ListNode{Val: 1, Next: nil}
    n5 := &ListNode{Val: 0, Next: n6}
    n4 := &ListNode{Val: 3, Next: n5}
    n3 := &ListNode{Val: 3, Next: n4}
    n2 := &ListNode{Val: 0, Next: n3}
    n1 := &ListNode{Val: 1, Next: n2}


    // fmt.Printf("List %s :: Res: %t\n", one, isPalindrome(one))
    fmt.Printf("List %s :: Res: %t\n", n1, isPalindrome1(n1))
}

func isPalindrome1(head *ListNode) bool {

    fmt.Printf("List: %s\n", head)

    if head.Next == nil {
        return true
    }
    
    if head.Next.Next == nil {
        if head.Val == head.Next.Val {
            return true    
        } else {
            return false
        }
        
    }
    
    if head.Next.Next.Next == nil {
        if head.Val == head.Next.Next.Val {
            return true    
        } else {
            return false
        }
        
    }

    // step 1: reverse the second Half
    slow, fast := head, head.Next.Next
    for fast != nil {
        if fast.Next == nil {
            fast = nil
        } else {
            fast = fast.Next.Next
        }
        slow = slow.Next
    }

    mid := slow
    var p, q, r *ListNode
    
    fmt.Printf("Mid: %s\n", mid)

    p = mid.Next
    if p != nil {
        q = p.Next
    }

    if q != nil {
        r = q.Next
    }
    p.Next = nil
    
    for r != nil {
        q.Next = p
        p = q
        q = r
        r = r.Next

    }
    
    q.Next = p
    mid.Next = q

    fmt.Printf("Reversed: %s\n", head)

    // step 2: check first half and second half are equal.
    f, s := head, mid.Next
    for s != nil {
        if f.Val != s.Val {
            return false
        }

        f = f.Next
        s= s.Next
    }

    return true
}


// func isPalindrome(head *ListNode) bool {
//     /* Method 1
//     slow, fast := head, head
//     firstHalf := make([]*ListNode, 0)

//     for fast != nil {
//         if fast.Next == nil{
//             fast = nil
//         } else {
//             firstHalf = append(firstHalf, slow)
//             fast = fast.Next.Next
//         }
//          slow = slow.Next
//     }

//     for i:=len(firstHalf)-1; i>=0; i-- {
//         if firstHalf[i].Val != slow.Val {
//             return false
//         }
//         slow = slow.Next
//     }

//     return true
//     */

//     // Method 2
//     // pptr := new(*ListNode)
// 	// return isPalindromeRecursive(head, head, pptr)

//     // Method 3
//     pptr := new(*ListNode1)
// 	*pptr = head
// 	return isPalindromeRecursive1(head, pptr)
// }


// func isPalindromeRecursive(slow, fast *ListNode, pptr **ListNode) bool {
// 	if fast == nil {
// 		*pptr = slow
// 		return true
// 	} else if fast.Next == nil {
// 		*pptr = slow.Next
// 		return true
// 	}

// 	res1 := isPalindromeRecursive(slow.Next, fast.Next.Next, pptr)
// 	res2 := slow.Val == (*pptr).Val
// 	fmt.Printf("slow: %d, pptr: %d\n", slow.Val, (*pptr).Val)
// 	*pptr = (*pptr).Next
// 	return res1 && res2
// }

// func isPalindromeRecursive1(next *ListNode1, pptr **ListNode1) bool {
// 	if next == nil {
// 		return true
// 	}
// 	res1 := isPalindromeRecursive1(next.Next, pptr)
// 	res2 := next.Val == (*pptr).Val
// 	fmt.Printf("next: %d, pptr: %d\n", next.Val, (*pptr).Val)
// 	*pptr = (*pptr).Next
// 	return res1 && res2
// }
