type MinStack struct {
	stack []int
	idx int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		idx: -1,
	}
}

func (this *MinStack) Push(val int) {
	this.idx += 1
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.idx]
	this.idx -= 1
}

func (this *MinStack) Top() int {
	return this.stack[this.idx]
}

func (this *MinStack) GetMin() int {
	var minSoFar int
	for i, element := range this.stack[:this.idx+1] {
		if i==0 {
			minSoFar = element
		} else {
			minSoFar = min(minSoFar, element)
		}
	}
	return minSoFar
}
