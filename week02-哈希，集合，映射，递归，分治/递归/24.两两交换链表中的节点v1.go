/*
 * @lc app=leetcode.cn id=24 lang=golang
 * @lcpr version=20003
 *
 * [24] 两两交换链表中的节点
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
func swapPairs(head *ListNode) *ListNode {

	/*
		   输入一个链表
		   从头结点开始，相连节点交换

		    pre = nil
		    cur = head

		    next = cur.Next

			next.Next = cur
			cur.Next = next

		   输出交换后的头结点
	*/

	swap(head, head.Next)
	return head.Next
}

func swap(head *ListNode, next *ListNode) {
	if head == nil {
		return
	}
	if next == nil {
		return
	}

	next = next.Next
	next.Next = head
	head.Next = next
	// next
	// 2 -> 1 -> 3
	//     head
	head = head.Next
	if head == nil {
		return
	}
	next = head.Next
	swap(head, next)
}

/*
 总结：
    思维混乱, 链表操作还是不熟悉, 道理知道，但手上来不了
*/

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

