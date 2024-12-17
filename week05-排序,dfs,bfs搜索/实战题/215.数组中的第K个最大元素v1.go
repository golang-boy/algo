/*
 * @lc app=leetcode.cn id=215 lang=golang
 * @lcpr version=20004
 *
 * [215] 数组中的第K个最大元素
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findKthLargest(nums []int, k int) int {
	/*
		   输入： 整数序列，整数k
		   输出：第k个最大元素
		        元素中有可能重复

			堆排序,弹出第k个
	*/

	h := &MaxHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
	}

	for i := 1; i < k; i++ {
		heap.Pop(h)
	}

	x := heap.Pop(h)
	return x.(int)
}

type MaxHeap []int

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func (m MaxHeap) Less(i, j int) bool {
	return (m)[i] >= (m)[j]
}
func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m MaxHeap) Len() int {
	return len(m)
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/

