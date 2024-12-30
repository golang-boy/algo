
/*
 * @lc app=leetcode.cn id=122 lang=golang
 * @lcpr version=20004
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int) int {
	// 将索引从1开始
	prices = append([]int{0}, prices...)

	// 初始化, 值设置为最小值
	f := make([][]int, len(prices))
	// f[i] 表示在第i天最大收益
	for i := 0; i < len(prices); i++ {
		f[i] = make([]int, 2)

		// 0表示不持仓, 1表示持仓
		f[i][0] = math.MinInt
		f[i][1] = math.MinInt
	}
	// 第一个元素值为0
	f[0][0] = 0

	for i := 1; i < len(prices); i++ {

		// 买, 已有收益减去价格
		f[i][1] = max(f[i][1], f[i-1][0]-prices[i])
		// 卖, 已有收益加上价格
		f[i][0] = max(f[i][0], f[i-1][1]+prices[i])

		for j := 0; j < 2; j++ {
			// 不活动
			f[i][j] = max(f[i][j], f[i-1][j])

			// 因为从第一天开始进行递推，与f[i][j]取最大是一种简化写法
			// 每次循环时，f[i][j]表示当前最大，或者前面过程中的最大
		}

	}
	// 不持仓时，最大的收益
	return f[len(prices)-1][0]
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/

