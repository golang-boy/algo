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
	   输出交换后的头结点
	*/

	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next
	others := head.Next.Next

	second.Next = first
	first.Next = swapPairs(others)

	return second
}

/*
 总结：
    时间复杂度：O(n)
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

