/*
 * @lc app=leetcode.cn id=LCR 141 lang=golang
 * @lcpr version=20003
 *
 * [LCR 141] 训练计划 III
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

func trainningPlan(head *ListNode) *ListNode {
	var newHead *ListNode

	if head == nil {
		return nil
	}

	head = reverse(head, &newHead)
	head.Next = nil
	return newHead
}

func reverse(head *ListNode, newHead **ListNode) *ListNode {

	/*
	  链表反转
	     是否需要定义前一个，后一个？

	  递归
	*/

	if head.Next == nil {
		*newHead = head
		return head
	}
	n := reverse(head.Next, newHead)
	n.Next = head
	return head
}

/*
总结:

   耗时 0:13:39
   重点是返回结果的处理，时间花在思考这里
   分析阶段不好判断是否可以源函数上进行递归，需要评估是否有全局变量需要传递
*/

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

