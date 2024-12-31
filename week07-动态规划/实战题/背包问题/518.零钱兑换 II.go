/*
 * @lc app=leetcode.cn id=518 lang=golang
 * @lcpr version=20004
 *
 * [518] 零钱兑换 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func change(amount int, coins []int) int {

	/*
	  零钱兑换 =》完全背包模型

	  零钱 == 物品，体积 == 面值， 价值 == 1 求min

	  =====>   完全背包问题

	  输入金额amount和coins数组


	  输出凑成总金额的硬币组合

	  f[i] 金额为i时，硬币组合数组
	*/

	coins = append([]int{0}, coins...)

	f := make([]int, amount+1)

	// 因为是计数，1表示唯一的方法
	f[0] = 1

	for i := 1; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			f[j] += f[j-coins[i]]
		}
	}

	return f[amount]
}

// @lc code=end

/*
// @lcpr case=start
// 5\n[1, 2, 5]\n
// @lcpr case=end

// @lcpr case=start
// 3\n[2]\n
// @lcpr case=end

// @lcpr case=start
// 10\n[10]\n
// @lcpr case=end

*/

