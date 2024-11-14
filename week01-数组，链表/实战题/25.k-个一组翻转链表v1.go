/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {

	/*
		   关键信息：
		    1. 一个单链表
			2. 分组反转
			3. k不是整数倍时，最后剩余的节点保持顺序
			4. 返回修改后的链表 ==> 需要返回头结点

		   分析：
			 2. 分组内的元素反转
			 3. 分组内反转后, 需要知道头节点和尾节点，方便与前后分组连接
			 4. 循环遍历, 每次遍历k个, 然后下一个循环, 起始点为当前节点
			 5. 什么时候结束循环?  剩余节点数不足k时, 结束循环, end为nil时
			 6. 返回start节点

		   老师的思路：
		     1. 分组， 往后走k-1步，开头和结尾节点
			 2. 操作什么? 组内反转
			 3. 更新每组更前一组和后一组的关系

	*/

	reverse := func(start *ListNode, k int) (*ListNode, *ListNode) {
		var pre *ListNode = nil
		var dummy = start
		var cur = start
		for {

			// 能遍历到k个节点的情况
			if k == 0 {
				break
			}
			next := cur.Next // 保存下一个节点
			cur.Next = pre   // 反转

			// 移动
			pre = cur  // 保存当前节点
			cur = next // 移动到下一个节点
			k--
		}

		dummy.Next = cur // cur为下一个分组头节点

		return pre, dummy
	}

	start := head

	pre := nil
	for {
		// 得获取反转后的分组的头节点, 和尾节点
		start, end = reverse(start, k)

		if end.Next == nil {
			break
		}

		pre.Next = start // 连接前后分组
		start = end.Next // 移动到下一个分组
		pre = end
	}
}

// @lc code=end

