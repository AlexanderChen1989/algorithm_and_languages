package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result ListNode 
	p := &result;

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1 
			list1 = list1.Next
		} else {
			p.Next = list2 
			list2 = list2.Next
		}
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1 
	}
	if list2 != nil {
		p.Next = list2
	}

	return result.Next
}
