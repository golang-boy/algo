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

		定义：f[i] 在索引i位置时，最长递增子序列的长度

		边界：
		    f[0] = 1

			f[1] = 2    nums[1] >  nums[0]
			f[1] = 1    nums[1] <= nums[0]

			f[2] = f[1] + 1     nums[2] > nums[1]
			f[2] = f[1] + 1     nums[2] <= nums[1]

		方程:
		    f[i] = f[i-1] + 1      nums[i] > nums[i-1]

			f[i] = max(f[j]+1, f[i])  j in [0,i]

		目标：

		    max(f)
	*/

	f := make([]int, len(nums))

	// 每个位置长度初始都为1，不可能为0
	for i := 0; i < len(nums); i++ {
		f[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		// 在前面的序列中找个取最大长度
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
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

