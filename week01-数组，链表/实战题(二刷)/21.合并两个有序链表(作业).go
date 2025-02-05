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

	head := &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			cur = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = list2
			list2 = list2.Next
		}
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

