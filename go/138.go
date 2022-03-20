package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := map[*Node]int{}
	n := 0
	p := head
	for p != nil {
		m[p] = n
		n++
		p = p.Next
	}
	p = head
	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i].Val = p.Val
		if i < n-1 {
			nodes[i].Next = &nodes[i+1]
		}
		if idx, ok := m[p.Random]; ok {
			nodes[i].Random = &nodes[idx]
		}
		p = p.Next
	}

	return &nodes[0]
}
