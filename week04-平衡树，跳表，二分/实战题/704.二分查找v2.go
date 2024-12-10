/*
 * @Author: 刘慧东
 * @Date: 2024-12-10 10:42:09
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-12-10 10:53:25
 */
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
	var ans = -1

	r(left, right, nums, target, &ans)
	return ans
}

func r(left, right int, nums []int, target int, ans *int) {

	if left > right {
		*ans = -1
		return
	}

	mid := (left + right) / 2

	if nums[mid] == target {
		*ans = mid
		return
	}

	if nums[mid] < target {
		left = mid + 1
		r(left, right, nums, target, ans)

	} else if nums[mid] > target {
		right = mid - 1
		r(left, right, nums, target, ans)
	}
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

