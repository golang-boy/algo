/*
 * @lc app=leetcode.cn id=143 lang=golang
 * @lcpr version=20003
 *
 * [143] 重排链表
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
func reorderList(head *ListNode) {
	r(head, &head)
}

func r(head *ListNode, h **ListNode) {
	if head == nil {
		return
	}

	r(head.Next, h)

	if head.Next != nil {
		fmt.Printf("%p (%d) %p(%d) %p(%d)\n", head, head.Val, head.Next, head.Next.Val, *h, (*h).Val)
	}

	if head == (*h).Next || head == *h {
		return
	}
	cur := &head
	for {
		if *cur == nil {
			break
		}
		if *cur == *h {
			return
		}
		cur = &((*cur).Next)
	}

	// 插入末尾
	temp := (*h).Next
	(*h).Next = head
	head.Next = temp
	*h = temp
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

*/

