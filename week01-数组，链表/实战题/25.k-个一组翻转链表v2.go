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
		思路分析：
		  已知条件:
		    1. 单链表
			2. k个一组
			3. 组内反转
			4. 组间拼接

		1. 遍历链表，找到分组的起始节点和结束节点
		2. 反转起始节点和结束节点之间的链表
		3. 将反转后的链表进行拼接



		返回什么？
		    反转后的链表的头结点  ==> 分组反转后，头结点会变，需要有一个变量来保存头结点dummy
	*/

	dummy := &ListNode{Next: head}
	last := dummy

	for {
		if head == nil {
			break
		}

		//	1. 遍历链表，找到分组的起始节点和结束节点
		tail := getTail(head, k)
		if tail == nil {
			break
		}
		// 保存下一组的起始节点
		nextGroup := tail.Next

		//	2. 反转起始节点和结束节点之间的链表
		reverse(head, nextGroup)

		//	3. 将反转后的链表进行拼接
		// 此时，head变为尾，tail变为头, 要进行拼接，需要知道上一个组的尾
		// 连接前面的
		last.Next = tail
		// 连接后面的
		head.Next = nextGroup

		// 4. 继续循环
		last = head
		head = nextGroup
	}

	return dummy.Next

}

func getTail(head *ListNode, k int) *ListNode {
	for i := 0; i < k-1; i++ { // k-1次, head为第一个节点, 需要注意这里的边界条件
		if head == nil {
			return head
		}
		head = head.Next
	}
	return head
}

func reverse(head, tail *ListNode) {
	var pre *ListNode

	for head != tail {

		next := head.Next
		head.Next = pre

		pre = head
		head = next
	}
}

// @lc code=end

