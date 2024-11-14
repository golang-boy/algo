/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	/*

		   1. 先找到相遇点
		   2. 分别从相遇点和起始点出发，相等时即为入口点

		     为什么这样呢？
			   是个数学计算的结果
			     设环为r个节点，到达环的入口走了l个节点，此后相遇时，走了p个节点，快指针总共走了k圈
				 则  2(l+p) = l + p + k*r, 左侧为2倍速度，右侧为1倍速度
				 则 l = (k-1)*r + (r-p), 同余关系,
				 r-p为从相遇点到入口的距离，(k-1)*r为从起始点到入口的距离
	*/

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			for head != slow {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}

// @lc code=end

