package 栈
// link https://leetcode-cn.com/problems/implement-queue-using-stacks/
/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	//in 栈如果不为空 in栈内的元素入out栈
	this.inToOut()
	res := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.inToOut()
	return this.out[len(this.out)-1]
}
func (this *MyQueue) inToOut() {
	if len(this.in) > 0 {
		this.out = append(this.out, this.in...)
		this.in = this.in[:0]
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
