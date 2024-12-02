
/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=20003
 *
 * [23] 合并 K 个升序链表
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
func mergeKLists(lists []*ListNode) *ListNode {

	h := &MinHeap{}

	var head = &ListNode{}
	var cur = head

	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	for h.Len() > 0 {
		x := heap.Pop(h).(*ListNode)
		cur.Next = x
		cur = x
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
	}

	return head.Next

}

type MinHeap []*ListNode

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Swap(i, j int) {
	(m)[i], (m)[j] = (m)[j], (m)[i]
}

func (m MinHeap) Less(i, j int) bool {
	return (m)[i].Val < (m)[j].Val
}

func (m *MinHeap) Push(x any) {
	(*m) = append((*m), x.(*ListNode))
}

func (m *MinHeap) Pop() any {
	x := (*m)[m.Len()-1]
	(*m) = (*m)[:m.Len()-1]
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/

