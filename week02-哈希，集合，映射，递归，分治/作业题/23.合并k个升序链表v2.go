/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}

	left := lists[:n/2]
	right := lists[n/2:]

	l := mergeKLists(left)
	r := mergeKLists(right)

	return mergeTwo(l, r)
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwo(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwo(l1, l2.Next)
	return l2
}

// @lc code=end

