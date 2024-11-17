/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	/* 关键信息:
	   1. push, pop, top 是栈的基本操作, 需要一个栈结构
	   2. getMin 是一个拓展操作，需要o(1)时间复杂度的实现

	   分析思路:

	   1. 入栈时，比较当前值和最小值, 变量保存, getMin 时直接返回, 满足o(1)
	   2. 出栈时，比较当前值是否是最小值，不是这没影响，是则需要更新值, 需要在栈中找到最小的
	      1. 因此出栈时，需要遍历栈，找最小值，出栈的时间复杂度是o(n)
		  2. getMin 的时间复杂度仍旧是o(1)

	   能否在入栈时，将最小值处理好，出栈时，降低时间复杂度呢？
	     最小值的变化，遵循先进后出的特点，因此，可以使用一个辅助栈，保存最小值

		   如果当前值小于最小值,这入栈,如果大于时，最小值继续入栈

		   两个栈，入栈和出栈是同步进行的, 每入一个元素，最小值栈就入一个元素，出栈时，两个栈同时出栈
	*/

	minStack []int
	stack    []int
}

func Constructor() MinStack {
	return MinStack{
		minStack: make([]int, 0),
		stack:    make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {

	this.stack = append(this.stack, val)

	who := func(v int) int {

		if len(this.minStack) == 0 {
			return v
		}

		if v < this.GetMin() {
			return v
		}

		if v > this.GetMin() {
			return this.GetMin()
		}
		return v
	}

	v := who(val)
	this.minStack = append(this.minStack, v)

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

