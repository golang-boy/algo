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
		输入整数数组, 找最大和连续子数组
		输出最大和


		1. 求取 子区间和s[l,r]最大值
		2. 要使s[l,r] = s[r]-s[l-1] 最大，需求s[l-1]值最小
		3. 因此需要计算每个索引i位置的最小前缀和, 初始化preMin和s数组


		s的长度和n+1
	*/

	n := len(nums)
	s := make([]int, n+1)
	preMin := make([]int, n+1)

	// 计算前缀和
	s[0] = 0
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]
	}

	// 初始化preMin
	preMin[0] = s[0]
	for i := 1; i <= n; i++ {
		preMin[i] = min(preMin[i-1], s[i])
	}

	ans := math.MinInt
	for i := 1; i <= n; i++ {
		// 在i之前找一个j，使s[i]-s[j]最大，因此s[j]需要取前i-1索引对应的前缀和的最小值
		ans = max(ans, s[i]-preMin[i-1])
	}

	return ans
}

/*
总结：
	上述思路是使用前缀min方法来解决问题,可以当做模板题来记忆
*/

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

