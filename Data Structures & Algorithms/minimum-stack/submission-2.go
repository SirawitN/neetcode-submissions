type MinStack struct {
	heap []int
	size int
}

func Constructor() MinStack {
	return MinStack{
		heap: []int{},
		size: 0,
	}
}

func (this *MinStack) Push(val int) {
	this.heap = append(this.heap, val)
	this.size += 1
}

func (this *MinStack) Pop() {
	this.size -= 1
	this.heap = this.heap[:this.size]
}

func (this *MinStack) Top() int {
	return this.heap[this.size-1]
}

func (this *MinStack) GetMin() int {
	var minSoFar int
	for i, element := range this.heap[:this.size] {
		if i==0 {
			minSoFar = element
		} else {
			minSoFar = min(minSoFar, element)
		}
	}
	return minSoFar
}
