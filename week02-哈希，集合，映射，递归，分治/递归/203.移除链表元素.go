/*
 * @lc app=leetcode.cn id=203 lang=golang
 * @lcpr version=20003
 *
 * [203] 移除链表元素
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
func removeElements(head *ListNode, val int) *ListNode {

	/*
	   输入： 链表，值
	   输出： 链表

	   逻辑处理：删除链表中值为val的所有节点

	   思路：递归
	*/
	if head == nil {
		return head
	}

	if head.Val == val {
		head = removeElements(head.Next, val)
	} else {
		head.Next = removeElements(head.Next, val)
	}
	return head
}

/*
总结：
    这是一道简单题，知道是递归，但是没有分析指针的变化,即返回节点指针与head的关系

	链表递归题目需要分析清楚递归返回，与当前节点的关系
*/

// @lc code=end

/*
// @lcpr case=start
// [1,2,6,3,4,5,6]\n6\n
// @lcpr case=end

// @lcpr case=start
// []\n1\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7]\n7\n
// @lcpr case=end

*/

