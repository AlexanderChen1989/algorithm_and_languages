package main


type MyCircularQueue struct {
    front *Node 
    back *Node 
}

type Node struct {
    Val int 
    Next *Node
}


func Constructor(k int) MyCircularQueue {
    
    var head Node 
    last := &head
    
    for i := 0; i < k; i++ {
        last.Next = &Node{}
        last = last.Next
    }
    
    last.Next = &head
    
    
    return MyCircularQueue {
        front: last,
        back: last,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }
    
    this.back = this.back.Next
    this.back.Val = value 
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    
    this.front = this.front.Next
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.front.Next.Val
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.back.Val
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.back == this.front
}


func (this *MyCircularQueue) IsFull() bool {
    return this.back.Next == this.front
}

