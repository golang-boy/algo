/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=20004
 *
 * [53] 最大子数组和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSubArray(nums []int) int {

	/*

	  f[i]  表示以i结尾的最大子序和

	  f[i] = max(f[i-1] + nums[i], nums[i])

	  f[i-1]可能是个负数，因此此处需要俩者取最大，
	  表示新开始，还是在之前的上连续进行


	*/
	n := len(nums)
	f := make([]int, n)

	f[0] = nums[0]
	for i := 1; i < n; i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
	}

	ans := f[0]
	for _, v := range f {
		ans = max(ans, v)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/

