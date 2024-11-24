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

	return r(head, &head)
}

func r(head *ListNode, head1 **ListNode) bool {
	// 不停的next，直到最下面
	//  递归就是一个栈，将前面入栈，直到末尾
	//  到达末尾后，与带入的head值进行比较
	//  比较完毕，更新head1,执行下一个

	if head != nil {

		if !r(head.Next, head1) {
			return false
		}

		if head.Val != (*head1).Val {
			return false
		}
		*head1 = (*head1).Next
	}
	return true
}

/*
总结：
   递归函数参数中携带全局信息, 指向头

   进入递归后，先执行递归函数，到达尾部
   将尾部数据和全局数据进行比较
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

