package main 


func removeElements(head *ListNode, val int) *ListNode {
    var result ListNode 
    last := &result
    for head != nil {
        next := head.Next
        head.Next = nil
        if head.Val != val {
            last.Next = head
            last = last.Next
        }
        head = next
    }
    
    return result.Next
}