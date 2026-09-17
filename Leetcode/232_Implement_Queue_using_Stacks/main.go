package main

import "fmt"

type MyQueue struct {
	input  []int
	output []int
}

func Constructor() MyQueue {
	return MyQueue{
		input:  make([]int, 0),
		output: make([]int, 0),
	}
}
func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}
func (this *MyQueue) Pop() int {
	this.move()
	last := len(this.output) - 1
	value := this.output[last]
	this.output = this.output[:last]

	return value
}

func (this *MyQueue) Peek() int {
	this.move()

	return this.output[len(this.output)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.input) == 0 && len(this.output) == 0
}

func (this *MyQueue) move() {
	// Old elements in output must be consumed first.
	if len(this.output) > 0 {
		return
	}
	// Move elements from input to output, reversing their order.
	for len(this.input) > 0 {
		last := len(this.input) - 1
		this.output = append(this.output, this.input[last])
		this.input = this.input[:last]
	}
}

func main() {
	queue := Constructor()
	queue.Push(10)
	queue.Push(20)
	queue.Push(30)
	fmt.Println(queue.Pop()) // 10
	queue.Push(40)
	queue.Push(50)
	fmt.Println(queue.Pop())   // 20
	fmt.Println(queue.Pop())   // 30
	fmt.Println(queue.Pop())   // 40
	fmt.Println(queue.Pop())   // 50
	fmt.Println(queue.Empty()) // true
}
