
/*
 * @lc app=leetcode.cn id=121 lang=golang
 * @lcpr version=20004
 *
 * [121] 买卖股票的最佳时机
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int) int {
	/*
		定义：
		    f[i] 在第i天时，最大的利润
		决策：
		    买： 下一天如果涨, prices[i] < prices[i+1]

			f[i] = f[i-1] -

			卖： 下一天如果涨，prices[i] < prices[i+1]

		边界：
			i in [0, len(prices)-1]
		目标：
	*/

	f := make([]int, len(prices))
	f[0] = 0
	minIndex := 0
	for i := 1; i < len(prices); i++ {

		min := prices[i]
		for j := minIndex; j < i; j++ {
			// 在j位置买入，i位置卖出
			if prices[j] < min {
				min = prices[j]
				minIndex = j
			}
		}

		f[i] = max(prices[i]-min, f[i-1])
	}

	return f[len(prices)-1]
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
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/

