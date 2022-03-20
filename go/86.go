package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func partition(head *ListNode, x int) *ListNode {
    var l1, l2 ListNode 
    l1Last, l2Last := &l1, &l2 
    
    for head != nil {
        next := head.Next 
        head.Next = nil
        if head.Val < x {
            l1Last.Next = head
            l1Last = l1Last.Next 
        } else {
            l2Last.Next = head
            l2Last = l2Last.Next
        }
        head = next
    }
    
    l1Last.Next = l2.Next
    
    return l1.Next
}