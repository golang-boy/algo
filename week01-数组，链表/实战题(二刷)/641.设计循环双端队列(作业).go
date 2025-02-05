/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start

/*
   双端循环队列
   1. 队列长度固定
   2. 能在对头入队和出队，也能在队尾入队和出队
   3. 队列满时，不能入队, 队列空时，不能出队

*/

type MyCircularDeque struct {
}

func Constructor(k int) MyCircularDeque {

}

func (this *MyCircularDeque) InsertFront(value int) bool {

}

func (this *MyCircularDeque) InsertLast(value int) bool {

}

func (this *MyCircularDeque) DeleteFront() bool {

}

func (this *MyCircularDeque) DeleteLast() bool {

}

func (this *MyCircularDeque) GetFront() int {

}

func (this *MyCircularDeque) GetRear() int {

}

func (this *MyCircularDeque) IsEmpty() bool {

}

func (this *MyCircularDeque) IsFull() bool {

}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end
