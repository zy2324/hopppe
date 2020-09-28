package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

type MinStack struct {
	Data   []int
	Length int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := new(Minstack)
	var stack Minstack
	return stack
}

func (this *MinStack) Push(x int) {
	this.Data[length] = x
	this.Length++
}

func (this *MinStack) Pop() {
	if this.Length == 0 {
		fmt.Println("stack is null")
		return
	}

	this.Data = this.Data[:this.Length-1]
	this.Length--
}

func (this *MinStack) Top() int {
	if this.Length == 0 {
		fmt.Println("stack is null")
		return
	}

	return this.Data[this.Length-1]
}

func (this *MinStack) GetMin() int {
	min := this.Data[0]

	for i := 0; i < this.Length; i++ {
		if min > this.Data[i] {
			min = this.Data[i]
		}
	}
	return min
}

/**
* Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
  * obj.Push(x);
   * obj.Pop();
    * param_3 := obj.Top();
     * param_4 := obj.GetMin();
     题解中都用了双切片
*/
