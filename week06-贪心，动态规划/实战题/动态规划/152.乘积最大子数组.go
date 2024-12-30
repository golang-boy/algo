/*
 * @Author: 刘慧东
 * @Date: 2024-12-27 10:16:44
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-12-27 11:34:55
 */
/*
 * @lc app=leetcode.cn id=152 lang=golang
 * @lcpr version=20004
 *
 * [152] 乘积最大子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProduct(nums []int) int {

	/*

	   nums[i] * f(i-1) > 0

	   f[i] 结束位置为i时，连续子数组最大的乘积

	   f[i] =  f[i-1]*nums[i]

	   前两个是连续情况时的关系，后一个是如果乘积为0时，新连续数组的开始

	   fmin[i] = min(fmin[i-i] *nums[i], fmax * nums[i], nums[i])
	   fmax[i] = max(fmin[i-1] *nums[i], fmax * nums[i], nums[i])

	*/

	fmin := make([]int, len(nums))
	fmax := make([]int, len(nums))

	fmin[0] = nums[0]
	fmax[0] = nums[0]

	for i := 1; i < len(nums); i++ {

		fmin[i] = min(min(fmin[i-1]*nums[i], fmax[i-1]*nums[i]), nums[i])
		fmax[i] = max(max(fmin[i-1]*nums[i], fmax[i-1]*nums[i]), nums[i])
	}

	max := math.MinInt

	for i := 0; i < len(nums); i++ {
		if max < fmax[i] {
			max = fmax[i]
		}
	}
	return max

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,-2,4]\n
// @lcpr case=end

// @lcpr case=start
// [-2,0,-1]\n
// @lcpr case=end

*/

