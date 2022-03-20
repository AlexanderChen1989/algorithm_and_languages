package main

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func sortList(head *ListNode) *ListNode {
    var vals []int
    
    p := head

    for p != nil {
        vals = append(vals, p.Val)
        p = p.Next
    }
    
    sort.Ints(vals)
    p = head
    
    for i := 0; i < len(vals); i++ {
        p.Val = vals[i]
        p = p.Next
    }
    
    return head
}