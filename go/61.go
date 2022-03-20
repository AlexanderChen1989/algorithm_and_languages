package main 


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }
    var nodes []*ListNode 
    p := head 
    for p != nil {
        nodes = append(nodes, p)
        p = p.Next
    }
    k %= len(nodes)
    if k == 0 {
        return head
    }
    nodes[len(nodes)-1].Next = nodes[0]
    idxHead := len(nodes) - k
    idxLast := (idxHead - 1 + len(nodes)) % len(nodes)
    
    head = nodes[idxHead]
    nodes[idxLast].Next = nil
    
    return head
}