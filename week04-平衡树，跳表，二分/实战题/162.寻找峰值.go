/*
 * @lc app=leetcode.cn id=162 lang=golang
 * @lcpr version=20004
 *
 * [162] 寻找峰值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findPeakElement(nums []int) int {

	/*
	  输入：整数数组，大于左右相邻的元素为峰值
	     nums[i] != nums[i + 1]  ==》 相邻的肯定不等

	  输出：可能有多个峰值,返回任意一个峰值


	  条件为比左右相邻的大，即为峰值

	  left < mid < right


	*/

	left := 0
	right := len(nums) - 1

	for left < right {
		lmid := (left + right) / 2
		rmid := lmid + 1

		// 取两个点，判断是递增还是递减
		if nums[lmid] <= nums[rmid] {
			// 递增则移动左侧
			left = lmid + 1
		} else {
			// 递减则移动右侧, 每次移动都是一半的范围
			right = rmid - 1
		}
	}
	return left

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,3,5,6,4]\n
// @lcpr case=end

*/

