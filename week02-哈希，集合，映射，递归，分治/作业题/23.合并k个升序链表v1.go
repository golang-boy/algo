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

	// var cur = dummy
	// var min *ListNode = lists[0]
	// if min == nil {
	// 	return nil
	// }
	// for i := 1; i < len(lists); i++ {
	// 	if lists[i] != nil && lists[i].Val <= min.Val {
	// 		min = lists[i]
	// 	}
	// }

	// if min == nil {
	// 	return nil
	// }

	// cur.Next = min
	// cur = cur.Next
	// cur.Next = mergeKLists(lists)
	// return dummy.Next
	return merge(nil, lists)
}

func merge(l *ListNode, lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return l
	}
	l = mergeTwo(l, lists[0])
	return merge(l, lists[1:])
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

