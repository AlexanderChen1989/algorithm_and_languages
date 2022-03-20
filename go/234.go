package main 


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    n := 0 
    p := head 
    for p != nil {
        n++ 
        p = p.Next
    }
    var reversed *ListNode 
    p = head
    for i := 0; i < n / 2; i++ {
        next := p.Next 
        p.Next = reversed 
        reversed = p
        p = next
    }
    if n % 2 != 0 {
        p = p.Next 
    }
    
    for p != nil && reversed != nil {
        if p.Val != reversed.Val {
            return false
        }
        p = p.Next
        reversed = reversed.Next
    }
    
    return p == reversed
}