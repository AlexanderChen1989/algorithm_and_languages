package main 

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    started := false
    s, f := head, head 
    
    for  !started || s != f {
        if !started {
            started = true
        }
        if s == nil {
            return false 
        }
        if f == nil || f.Next == nil {
            return false
        }
        s = s.Next 
        f = f.Next.Next
    }
    
    return true
}