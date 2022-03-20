package main 


func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil {
        return nil
    }
	var nodes []*ListNode 
    for head != nil {
        nodes = append(nodes, head)
        head = head.Next
    }
    
    if n == len(nodes) {
        return nodes[1]
    }
    
    nodes[len(nodes) - n - 1].Next = nodes[len(nodes) - n].Next
    
    return nodes[0]
}