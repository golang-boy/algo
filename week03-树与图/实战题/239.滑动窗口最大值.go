
/*
 * @lc app=leetcode.cn id=239 lang=golang
 * @lcpr version=20003
 *
 * [239] 滑动窗口最大值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {

	// 使用heap实现滑动窗口

	h := &MaxHeap{}

	ans := []int{}
	heap.Init(h)

	// 从左到右循环
	for i := 0; i < len(nums); i++ {
		// 右侧入堆
		heap.Push(h, &Item{Val: nums[i], Index: i})

		// 右边界为i,左边界为i-k+1
		// right := i
		left := i - k + 1

		if left >= 0 { // i在此时说明窗口中以容纳k个元素

			for (*h)[0].Index <= left-1 {
				heap.Pop(h)
			}
			ans = append(ans, (*h)[0].Val) // 栈顶元素
		}
	}

	return ans
}

type Item struct {
	Val   int
	Index int
}

type MaxHeap []*Item

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Val > h[j].Val
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,-1,-3,5,3,6,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/

