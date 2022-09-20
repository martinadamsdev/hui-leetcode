package main

import "fmt"

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{queue: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	pop := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return pop
}

func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {
	obj := Constructor()
	test := [6]any{nil, 1, 2, nil, nil, nil}
	var i int
	fmt.Printf("%s", "[")
	for i = 0; i < 6; i++ {
		if test[i] != nil {
			obj.Push(i)
			fmt.Printf("%s", "null,")
		}
	}
	fmt.Printf("%d%s%d%s%v%s", obj.Top(), ",", obj.Pop(), ",", obj.Empty(), "]")
}
