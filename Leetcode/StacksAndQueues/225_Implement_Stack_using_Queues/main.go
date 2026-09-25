package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println(stack.Pop())
	stack.Push(40)
	stack.Push(50)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	for i := 0; i < len(this.queue)-1; i++ {
		first := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, first)
	}
}

func (this *MyStack) Pop() int {
	first := this.queue[0]
	this.queue = this.queue[1:]

	return first
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
