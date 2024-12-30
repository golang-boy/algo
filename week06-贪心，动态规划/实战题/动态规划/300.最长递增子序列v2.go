/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=20004
 *
 * [300] 最长递增子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLIS(nums []int) int {
	/*

		f(i) = 1+f(i+1)
		f(i) = f(i+1)


	*/

	dp := make([]int, len(nums)+1)

	for i := 1; i < len(nums); i++ {
		max := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}

		if nums[i-1] > max {
			dp[i-1] = 1 + dp[i]
		} else {
			dp[i-1] = dp[i]
		}
	}

	ans := 0
	for i := range dp {
		ans += dp[i]
	}

	return ans
}

func lengthLis(i int, nums []int) int {

	flag := false
	for j := i + 1; j < len(nums); j++ {
		if nums[j] > nums[i] {
			flag = true
			break
		}
	}
	if flag {
		return 1 + lengthLis(i, nums)
	} else {
		return lengthLis(i, nums)
	}

	n := len(nums)
	f := make([]int, n)

	for i := 0; i < n; i++ {
		f[i] = 1
	}

	for i := 1; i < n; i++ {

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
				// f[i] = f[j] + 1
			}

		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, f[i])
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
// [10,9,2,5,3,7,101,18]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

*/

