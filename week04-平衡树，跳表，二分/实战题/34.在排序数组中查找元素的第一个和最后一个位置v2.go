
/*
 * @lc app=leetcode.cn id=34 lang=golang
 * @lcpr version=20004
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func searchRange(nums []int, target int) []int {
	ans := []int{}

	left := 0
	right := len(nums)

	for left < right {

		mid := (left + right) / 2

		// 左8
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid
		}
	}

	ans = append(ans, right)

	left = -1
	right = len(nums) - 1
	for left < right {
		mid := (left + right + 1) / 2
		// 右8
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid
		} else if nums[mid] == target {
			left = mid
		}
	}
	ans = append(ans, right)

	if ans[0] > ans[1] {
		return []int{-1, -1}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [5,7,7,8,8,10]\n8\n
// @lcpr case=end

// @lcpr case=start
// [5,7,7,8,8,10]\n6\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/

