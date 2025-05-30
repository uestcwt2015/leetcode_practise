type Node struct {
    Val int
    Next *Node 
}

type MyLinkedList struct {
    head *Node
    tail *Node
    Len int
}


func Constructor() MyLinkedList {
    return MyLinkedList{
        Len: 0,
    }
}


func (this *MyLinkedList) Get(index int) int {
    if index >= this.Len {
        return -1
    }
    node := this.head
    for i := 0; i < index; i++ {
        node = node.Next
    }
    return node.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &Node{Val: val}
    if this.head == nil {
        this.head = newNode
        this.tail = newNode
    } else {
        newNode.Next = this.head
        this.head = newNode
    }
    this.Len++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    newNode := &Node{Val: val}
    if this.head == nil {
        this.head = newNode
        this.tail = newNode
    } else {
        this.tail.Next = newNode
        this.tail = newNode
    }
    this.Len++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.Len {
        return
    }
    if index == this.Len {
        this.AddAtTail(val)
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }

    node := this.head
    for i := 0; i < index-1; i++ {
        node = node.Next
    }
    newNode := &Node{Val: val, Next: node.Next}
    node.Next = newNode
    this.Len++
    return
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index >= this.Len {
        return
    }

    if index == 0 && this.Len != 0 {
        this.head = this.head.Next
    } else {
        node := this.head
        for i := 0; i < index-1; i++ {
            node = node.Next
        }
        if node.Next != nil {
            node.Next = node.Next.Next
        }
        if index == this.Len - 1 {
            this.tail = node
        }
    }

    this.Len--
    if this.Len == 0 {
        this.head = nil
        this.tail = nil
    }
    return
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */