/*
 * @Author: 刘慧东
 * @Date: 2024-12-03 15:56:33
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-12-03 18:23:57
 */
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

	Init(h)

	for _, list := range lists {
		if list != nil {
			Push(h, list)
		}
	}

	for h.Len() > 0 {
		x := Pop(h).(*ListNode)
		cur.Next = x
		cur = x
		if cur.Next != nil {
			Push(h, cur.Next)
		}
	}

	return head.Next

}

type MinHeap []*ListNode

type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}

func Init(m Interface) {
	n := m.Len()

	// 从最后一个非叶子节点开始，依次down
	// 直到所有非叶子节点都处理完毕

	for i := n/2 - 1; i >= 0; i-- {
		down(m, i, n)
	}
}

func Push(m Interface, x any) {
	// 追加到末尾
	// up上去
	m.Push(x)
	up(m, m.Len()-1)
}

func Pop(m Interface) any {

	// pop堆顶元素
	// 与末尾交换位置，然后down
	// 最后将末尾的pop掉
	tail := m.Len() - 1
	m.Swap(0, tail)
	down(m, 0, tail)
	return m.Pop()
}

func up(m Interface, i int) {

	p := (i - 1) / 2
	c := i

	for p >= 0 {
		if m.Less(c, p) {
			m.Swap(p, c)
			c = p
			p = (p - 1) / 2
		} else {
			break
		}
	}
}

func down(m Interface, i, n int) {
	// i处的节点进行down操作

	left := 2*i + 1
	right := 2*i + 2
	p := i

	for {
		// 找到左右较最小的,然后和根比较，如果比根小，则交换，否则退出
		if left < n {
			if right < n {
				c := min(m, left, right)
				if m.Less(c, p) {
					m.Swap(c, p)
					p = c
					left = 2*p + 1
					right = 2*p + 2
				} else {
					break
				}
			} else {
				if m.Less(left, p) {
					m.Swap(left, p)
					p = left
					left = 2*p + 1
					right = 2*p + 2
				} else {
					break
				}
			}
			continue
		}

		if right < n {
			if m.Less(right, p) {
				m.Swap(right, p)

				p = right
				left = 2*p + 1
				right = 2*p + 2

			} else {
				break
			}
		} else {
			break
		}

	}
}

func min(m Interface, i, j int) int {
	if m.Less(i, j) {
		return i
	}
	return j
}

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

