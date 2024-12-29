
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
	 */

	min := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {

		if min > prices[i] {
			min = prices[i]
		}
		maxProfit = max(maxProfit, prices[i]-min)
	}

	return maxProfit
}

/*
总结：
    max(maxProfit, 第i天的价格 - i前面最小的值)

	如何算最小值
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
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/

