/*
 * @lc app=leetcode.cn id=714 lang=golang
 * @lcpr version=20004
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int, fee int) int {

	prices = append([]int{0}, prices...)

	f := make([][]int, len(prices))

	for i := range f {
		f[i] = make([]int, 2)
		for j := range f[i] {
			f[i][j] = -1000000
		}
	}

	f[0][0] = 0

	// 从i开始进行递推
	for i := 1; i < len(prices); i++ {

		// i == 1时买入此f[0][0] = 0，买入后f[1][1] = 0-prices[i]

		// prices[0] 是追加进去的，i等于1时，是第一项, 买第一项时f[1][1] = - prices[1]
		f[i][1] = max(f[i][1], f[i-1][0]-prices[i])

		// 此时卖出
		// 一次完整的交易过程，买入时没有手续费，卖出时有手续费
		f[i][0] = max(f[i][0], f[i-1][1]+prices[i]-fee)
		for j := 0; j < 2; j++ {
			f[i][j] = max(f[i][j], f[i-1][j])
		}
	}

	return f[len(prices)-1][0]
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
// [1, 3, 2, 8, 4, 9]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,3,7,5,10,3]\n3\n
// @lcpr case=end

*/

