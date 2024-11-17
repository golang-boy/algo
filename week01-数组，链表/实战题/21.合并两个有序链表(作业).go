/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	/*

	   关键信息：
	   1. 两个链表都是有序的
	   2. 合并后也是有序的

	   约束：
	   1.左右俩链表，长度不一

	   思路：
	   1. 从俩链表头部各取一个，比较大小，小的节点放到新链表中
	   2. 然后小的节点向后移动

	    取谁谁移动, 移动到末尾后，另一个链表直接接上

	*/

	var head = &ListNode{}
	var cur = head

	for list1 != nil && list2 != nil {
		var temp *ListNode

		if list1.Val <= list2.Val {
			temp = list1
			list1 = list1.Next
		} else {
			temp = list2
			list2 = list2.Next
		}

		cur.Next = temp
		cur = temp
	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return head.Next
}

// @lc code=end

