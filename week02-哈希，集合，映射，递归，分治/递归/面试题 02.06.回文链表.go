/*
 * @lc app=leetcode.cn id=面试题 02.06 lang=golang
 * @lcpr version=20003
 *
 * [面试题 02.06] 回文链表
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	return r(head, &head)
}

func r(head *ListNode, old **ListNode) bool {
	if head == nil {
		return true
	}

	if !r(head.Next, old) {
		return false
	}

	if head.Val != (*old).Val {
		return false
	}

	*old = (*old).Next

	return true
}

// @lc code=end

/*
// @lcpr case=start
// 1->2\n
// @lcpr case=end

// @lcpr case=start
// 1->2->2->1\n
// @lcpr case=end

*/

