package main


func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil && head.Next == nil {
        return head 
    }
    
    p := head 
    
    for p.Next != nil {
        if p.Val == p.Next.Val {
            p.Next = p.Next.Next
        } else {
            p = p.Next
        }
    }
    
    return head
}