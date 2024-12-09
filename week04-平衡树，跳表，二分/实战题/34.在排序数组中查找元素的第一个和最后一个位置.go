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
	/*
	  1.nums是有序，有相同值    => 可以使用二分查找
	  2.找目标值的多个起止位置
	  3.存在返回索引，不在返回【-1，-1】

	  从左边找，找下边界
	  从右边找，找上边界
	*/

	// 找第一个 >= target的
	left := 0
	right := len(nums)

	downbound := -1
	upbound := -1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	downbound = left

	left = -1
	right = len(nums) - 1
	for left < right {
		mid := (left + right + 1) / 2

		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	upbound = right

	if downbound > upbound {
		return []int{-1, -1}
	}

	return []int{downbound, upbound}

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

