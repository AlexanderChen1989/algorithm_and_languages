package main 


func swapPairs(head *ListNode) *ListNode {
	var result ListNode
    last := &result 
    
    for head != nil && head.Next != nil {
        next := head.Next.Next
        a, b := head, head.Next 
        a.Next, b.Next = nil, nil 
        last.Next = b 
        last.Next.Next = a 
        last = a
        
        head = next
    }
    last.Next = head
    
    return result.Next
}