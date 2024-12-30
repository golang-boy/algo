
/*
 * @lc app=leetcode.cn id=198 lang=golang
 * @lcpr version=20004
 *
 * [198] 打家劫舍
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rob(nums []int) int {

	/*
		输入nums表示相邻房间存放的金额，非负整数数组

		输出不触动警报情况下，一夜能偷窃的最高金额

		什么时候可能触动警报？
		     相邻两家房间同一晚上被闯入

		// 方法一、
		f[i] 表示在i号房屋偷窃后的最高金额

		f[i] = max(f[i], f[i-j]+nums[i])  j -> [0,i-2] // 偷
		f[i] = max(f[i], f[i-1]) // 不偷

		f[len(nums)-1]

		// 方法二、

		f[i][j] 表示在i号房屋偷窃后的最高金额, j -> [0,1] 0表示偷，1表示不偷

		f[i][0] = max(f[i][0], f[i-1][1]+nums[i])
		f[i][1] = max(f[i][1], f[i-1][j])

		max(f[len(nums)-1][0], f[len(nums)-1][1])

	*/

	nums = append([]int{0}, nums...)

	f := make([][]int, len(nums))
	for i := range f {
		f[i] = make([]int, 2)

		for j := range f[i] {
			f[i][j] = -100000000
		}
	}

	// f[0][0] = 0
	f[0][1] = 0

	for i := 1; i < len(nums); i++ {
		for j := 0; j < 2; j++ {
			// 不偷到偷
			f[i][0] = max(f[i][0], f[i-1][1]+nums[i])

			// 不偷
			f[i][1] = max(f[i][1], f[i-1][j])
		}
	}

	return max(f[len(nums)-1][0], f[len(nums)-1][1])

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
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,7,9,3,1]\n
// @lcpr case=end

*/

