/*
 * @lc app=leetcode.cn id=153 lang=golang
 * @lcpr version=20004
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findMin(nums []int) int {

	// left := 0
	// right := len(nums)
	// for left < right {
	// 	mid := (right + left) / 2

	// 	if nums[mid] >= nums[right] {
	// 		right = mid
	// 	} else {
	// 		left = mid + 1
	// 	}
	// }

	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (right + left) / 2

		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[right]
}

/*

总结：

   一个是判定的不等式，如果需要包含，则为right=mid

   套用模版时，
   一定要注意，该问题是否有解，有解 left，right初始化时为0，n-1

   如果无解，模板1.1为 0，n,  模版1.2为-1，n-1, 同时mid计算时，需要额外加个1
*/

// @lc code=end

/*
// @lcpr case=start
// [3,4,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,7,0,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [11,13,15,17]\n
// @lcpr case=end

*/

