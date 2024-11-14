/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	/*

			关键点：
			    1. 一个单链表
				2. 反转
				3. 返回反转后的单链表的头节点

			分析：
		      1. 需要遍历链表
			  2. 变化的信息有哪些？
			     * 当前节点, 遍历时，当前节点不停的变
				 * 下一个节点, 随着当前节点的移动，下一个节点也在移动
				 * 前一个节点, 随着当前节点的移动，前一个节点也在移动
			  3. 当当前节点的next执行前一个节点时，当前节点就反转了
			     反转操作为：  cur.next = pre
			     但是，需要再处理下一个节点时，cur.next的值被覆盖，所以需要临时保存下一个节点
				 pre怎么来？ pre初始为nil,每轮循环时，pre = cur



	*/

	var pre *ListNode = nil
	var next *ListNode = nil
	cur := head
	for {
		if cur == nil { // 循环退出
			break
		}

		// 反转
		// 1. 当前节点的下一个节点指针，要指向前一个节点
		// 2. 临时保存下一个节点，将当前节点执行前一个

		next = cur.Next
		cur.Next = pre // 当前节点指向下一个节点

		pre = cur  //  获取前一个节点
		cur = next // 指针后移
	}
	return pre

}

// @lc code=end

