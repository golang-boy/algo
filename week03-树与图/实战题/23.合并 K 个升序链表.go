
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

	h := &Heap{}
	heap.Init(h)
	var head *ListNode = &ListNode{}
	var tail *ListNode = head

	for _, n := range lists {
		if n != nil {
			heap.Push(h, n)
		}
	}

	for h.Len() != 0 {
		n := heap.Pop(h).(*ListNode)
		tail.Next = n
		tail = tail.Next

		if n.Next != nil {
			heap.Push(h, n.Next)
		}

	}

	return head.Next

}

type Heap []*ListNode

func (n Heap) Len() int {
	return len(n)
}

func (n Heap) Less(i int, j int) bool {
	return n[i].Val < n[j].Val
}

func (n Heap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *Heap) Push(x any) {
	*n = append(*n, x.(*ListNode))
}

func (n *Heap) Pop() any {

	old := *n
	x := old[len(old)-1]
	*n = old[:len(old)-1]
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

