package main

import "fmt"

func main() {
	arrs := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	var lists []*ListNode

	for _, arr := range arrs {
		lists = append(lists, arrToList(arr))
	}

	list := mergeKLists(lists)
	printList(list)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func arrToList(arr []int) *ListNode {
	var head ListNode
	p := &head

	for _, v := range arr {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}

	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) >= 2 {
		var lsts []*ListNode
		for i := 0; i < len(lists)-1; i += 2 {
			lst := mergeTwoLists(lists[i], lists[i+1])
			lsts = append(lsts, lst)
		}
		if len(lists) % 2 != 0 {
			lsts = append(lsts, lists[len(lists)-1])
		}
		lists = lsts
	}

	return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result ListNode
	p := &result

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
