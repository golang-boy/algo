/*
 * @lc app=leetcode.cn id=322 lang=golang
 * @lcpr version=20004
 *
 * [322] 零钱兑换
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func coinChange(coins []int, amount int) int {

	/*
		定义:
			f[i] 最少的硬币个数, 在金额为i时

		方程：
			f[i] = min(f[i-coins[j]]+1 且 j in [0,len(coins)])

		边界：
			i in [0, amount]

		目标：
			f[amount]
	*/

	f := make([]int, amount+1)

	f[0] = 0
	for i := 1; i <= amount; i++ {
		f[i] = 10002
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				f[i] = min(f[i], f[i-coins[j]]+1)
			}
		}
	}

	if f[amount] >= 10002 {
		f[amount] = -1
	}
	return f[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1, 2, 5]\n11\n
// @lcpr case=end

// @lcpr case=start
// [2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/

