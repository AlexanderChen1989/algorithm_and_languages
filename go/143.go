package main

func reorderList(head *ListNode)  {
    if head.Next == nil {
        return
    }
    p := head
    n := 0 
    for p != nil {
        n += 1
        p = p.Next
    }
    p = head
    for i := 0; i < n / 2; i++ {
        node := p
        p = p.Next 
        if i == n / 2 - 1 {
            node.Next = nil
        }
    }
    var reversed *ListNode 
    for p != nil {
        next := p.Next 
        p.Next = reversed 
        reversed = p
        p = next
    }
    
    
    p = head
    var last *ListNode 
    for p != nil && reversed != nil {
        next := p.Next 
        p.Next = reversed
        if last != nil {
            last.Next = p
        }
        last = reversed 
        reversed = reversed.Next
        p = next
    }
}