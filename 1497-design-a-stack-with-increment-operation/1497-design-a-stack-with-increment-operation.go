type CustomStack struct {
    data []int
    capacity int
    top int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{
        data: []int{},
        capacity: maxSize,
        top: -1,
    }
}


func (this *CustomStack) Push(x int)  {
    if this.top + 1 < this.capacity {
        this.data = append(this.data, x)
        this.top++
    }
}


func (this *CustomStack) Pop() int {
    if this.top == -1 {
        return -1
    }
    v := this.data[this.top]
    this.data = this.data[:this.top]
    this.top--
    return v
}


func (this *CustomStack) Increment(k int, val int)  {
    if k > this.top + 1 {
        k = this.top + 1
    }

    for i := 0; i < k; i++ {
        this.data[i] += val
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */