/*
 * @lc app=leetcode.cn id=234 lang=golang
 * @lcpr version=20003
 *
 * [234] 回文链表
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

	/*
		   问题：判断是否是回文链表
		     1. 链表反转，判断每个节点是否相等
			 2. 递归

		   条件：
		     1. head.Val != head1.Val  return false
			 2. 任意一个等于nil时  return false
	*/

	head1 := reverse(head)
	return r(head, head1)
}

func r(head, head1 *ListNode) bool {

	if head == nil && head1 != nil {
		return false
	}

	if head != nil && head1 == nil {
		return false
	}

	if head.Val != head1.Val {
		return false
	}

	if head.Next == nil && head1.Next == nil {
		return true
	}

	return r(head.Next, head1.Next)
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode = head
	for {
		if cur == nil {
			return pre
		}
		next := cur.Next // 保存下一个节点
		cur.Next = pre   // 当前节点下一个执行前一个
		pre = cur        // 前一个移到当前
		cur = next       // 当前移动到下一个
	}
}

/*
总结：
    思路没问题，问题出在整个链表反转时，都是针对同一个链表进行操作的

	head为反转后的最末尾节点，head为反转后的头节点
	这种情况下朝着一个方向递归，必定不等，会出问题

	因此：
	  1.要么反转时构建一个新的链表，需要额外的n的空间,空间复杂度为O(n)
	  2.要么找到中点，把后半段反转，然后俩个半段进行比较

*/

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/

