/*
 * @lc app=leetcode.cn id=LCR 142 lang=golang
 * @lcpr version=20003
 *
 * [LCR 142] 训练计划 IV
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
func trainningPlan(l1 *ListNode, l2 *ListNode) *ListNode {

	/*
	   合并两个升序链表, 返回头节点
	*/

	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil && l2 != nil {
		return l2
	}

	if l1 != nil && l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l2.Next = trainningPlan(l1, l2.Next)
		return l2
	}

	l1.Next = trainningPlan(l1.Next, l2)
	return l1
}

/*
总结： 耗时 0:13:30
    1. 先写核心逻辑
	2. 最后写返
	3. 再写条件
*/

// @lc code=end

/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

*/

