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

	/*
		尝试用语言内置的有序集合库，或写一棵平衡树，来解决

		遍历数组，每次将元素和对应的索引放入树中，
		   当达到k时，取个值最大的, 当移动到下一个时，将最小的索引删除

		   平衡树需要具备取值最大，和索引最小的能力

	*/

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

