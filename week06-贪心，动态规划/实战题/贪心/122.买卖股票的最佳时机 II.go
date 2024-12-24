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

	/*
		输入： 股票的价格
		输出： 最大利润

		规则：最多持有一股, 可以买，可以卖, 同一天，可以买也可以卖

		选定第i天，判断后续的哪一天利润最大,循环处理

		max(p[j]-p[i])
	*/

	ans := 0

	// 在第i天买入，第j天卖出

	for i := 1; i < len(prices); i++ {
		ans += max(prices[i]-prices[i-1], 0)
	}

	// 卖出后如果还能买，则继续

	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/*

总结：
    什么时候买入，什么时候卖出

	没有时，第i天涨时买入 p[i-1] < p[i]
	有时，第i天涨时卖出 p[i-1] < p[i]
	否则，不卖,不买,不产生利润
*/

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

