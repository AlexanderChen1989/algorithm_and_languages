
package main 

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    var result ListNode
    lastResult := &result
    
    for head != nil {
        var reversed *ListNode 
        counter := 0 
        last := head
        for head != nil && counter < k {
            t := head.Next 
            head.Next = reversed 
            reversed = head
            counter++      
            head = t
        }
        
        if counter < k {
            var p *ListNode
            
            for reversed != nil {
                t := reversed.Next 
                reversed.Next = p 
                p = reversed 
                reversed = t
            }
            
            lastResult.Next = p
            break
        }
        
        lastResult.Next = reversed
        lastResult = last
    }
    
    return result.Next
}