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

	// 1. 分组
	// 2. 反转
	// 3. 连接分组

	start := &ListNode{Next: head}
	cur := start
	for cur != nil {
		end := group(cur.Next, k)
		if end != nil {
			end = reverseGroup(cur.Next, end)
		}
		cur = end
	}

	return start.Next
}

func reverseGroup(start *ListNode, end *ListNode) *ListNode {

	cur := start
	var last *ListNode = nil
	next := end.Next
	for cur != end {
		temp := cur.Next
		cur.Next = last
		last = cur
		cur = temp
	}
	start.Next = next
	return start
}

func group(cur *ListNode, k int) *ListNode {

	for cur != nil && k > 0 {
		cur = cur.Next
		k--
	}
	if k != 0 {
		return nil
	}
	return cur
}

// @lc code=end

