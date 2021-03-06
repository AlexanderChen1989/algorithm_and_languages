package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := map[*ListNode]bool{}
    
    for headA != nil {
        m[headA] = true
        headA = headA.Next 
    }
    
    for headB != nil {
        if m[headB] {
            return headB 
        }
        headB = headB.Next
    }
    
    return nil
}