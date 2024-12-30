/*
 * @Author: 刘慧东
 * @Date: 2024-12-30 15:05:42
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-12-30 16:09:46
 */

/*
 * @lc app=leetcode.cn id=213 lang=golang
 * @lcpr version=20004
 *
 * [213] 打家劫舍 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	nums = append([]int{0}, nums...)

	f := make([][]int, len(nums))
	for i := range f {
		f[i] = make([]int, 2)

		for j := range f[i] {
			f[i][j] = -100000000
		}
	}

	f[1][1] = 0
	// f[1][0] = -100000000 // 不偷1

	// 不偷1，偷n
	for i := 2; i < len(nums); i++ {
		for j := 0; j < 2; j++ {
			// 偷
			f[i][0] = max(f[i][0], f[i-1][1]+nums[i])

			// 不偷
			f[i][1] = max(f[i][1], f[i-1][j])
		}
	}

	ans := max(f[len(nums)-1][1], f[len(nums)-1][0])

	// 偷1，不偷n
	f[1][0] = nums[1]
	f[1][1] = 0

	for i := 2; i < len(nums); i++ {
		for j := 0; j < 2; j++ {
			// 偷
			f[i][0] = max(f[i][0], f[i-1][1]+nums[i])

			// 不偷
			f[i][1] = max(f[i][1], f[i-1][j])
		}
	}

	ans = max(ans, f[len(nums)-1][1])

	return ans

}

/*
总结：
	[1, !n]  [!1, n] [1, n] [!1, !n]
	从1开始时，包含了上述四种情况，现在只需要1和n不能同时存在的场景, 因此通过两次dp打破约束

	不选1是一种情况，需要从第2个开始, 到达末尾时，ans = max(f[n][0], f[n][1])
	选1是一种情况，循环从第2个开始, 到达末尾后，ans = max(ans, f[n][1]) 表示不选n的最大值

*/

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

