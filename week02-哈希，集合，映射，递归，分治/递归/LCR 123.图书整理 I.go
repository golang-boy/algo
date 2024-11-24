/*
 * @lc app=leetcode.cn id=LCR 123 lang=golang
 * @lcpr version=20003
 *
 * [LCR 123] 图书整理 I
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
func reverseBookList(head *ListNode) []int {

	/*
	  输入：链表头结点，节点值为书编号
	  输出：倒序的编号数组
	*/

	var res = []int{}
	r(head, &res)
	return res
}

func r(head *ListNode, res *[]int) {
	if head == nil {
		return
	}
	r(head.Next, res)
	*res = append(*res, head.Val)
}

/*
总结：
	耗时 0:3:52
*/

// @lc code=end

/*
// @lcpr case=start
// [3,6,4,1]\n
// @lcpr case=end

*/

