/*
 * @lc app=leetcode.cn id=704 lang=golang
 * @lcpr version=20004
 *
 * [704] 二分查找
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			// 模板值在左边
			right = mid - 1
		} else {
			// 模板值在右边
			left = mid + 1
		}
	}

	return -1
}

/*
二分查找前提：(能不能二分)
  1. 目标函数具有单调性（递增或递减）
  2. 存在上下边界
  3. 能够通过索引访问
*/

// @lc code=end

/*
// @lcpr case=start
// [-1,0,3,5,9,12]\n9\n
// @lcpr case=end

// @lcpr case=start
// [-1,0,3,5,9,12]\n2\n
// @lcpr case=end

*/

