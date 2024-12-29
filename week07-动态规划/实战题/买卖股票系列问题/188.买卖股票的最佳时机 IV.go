
/*
 * @lc app=leetcode.cn id=188 lang=golang
 * @lcpr version=20004
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(k int, prices []int) int {

	prices = append([]int{0}, prices...)

	f := make([][][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		f[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = make([]int, k+1) // 加1是为确保从1开始, 简化程序
			for h := 0; h <= k; h++ {
				f[i][j][h] = math.MinInt
			}
		}
	}
	f[0][0][0] = 0

	for i := 1; i <= len(prices)-1; i++ {
		for j := 0; j < 2; j++ {
			for h := 0; h <= k; h++ {
				// 买
				if h > 0 {
					f[i][1][h] = max(f[i][1][h], f[i-1][0][h-1]-prices[i])
				}
				// 卖
				//  持仓时+prices[i] = 等现在不持仓的收益
				f[i][0][h] = max(f[i][0][h], f[i-1][1][h]+prices[i])
				// 不做
				// 前一天的与现在的一样
				f[i][j][h] = max(f[i][j][h], f[i-1][j][h])

			}
		}
	}

	// 最多k笔交易，因此需要找最大的那个
	ans := math.MinInt

	for h := 0; h <= k; h++ {
		if ans < f[len(prices)-1][0][h] {
			ans = f[len(prices)-1][0][h]
		}
	}
	return ans
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
// 2\n[2,4,1]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[3,2,6,5,0,3]\n
// @lcpr case=end

*/

