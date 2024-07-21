package minstack

type MinStack struct {
	stack []int
	head  int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		head:  0,
		min:   0,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.head++

	if this.min > val {
		this.min = val
	}
}

func (this *MinStack) findMin() int {
	min := this.stack[0]

    for _, v := range this.stack {
		if v < min {
			min = v
		}
	}
	return min
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.head--
}

func (this *MinStack) Top() int {
	return this.stack[this.head]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
