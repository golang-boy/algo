
/*
 * @lc app=leetcode.cn id=123 lang=golang
 * @lcpr version=20004
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int) int {

	prices = append([]int{0}, prices...)

	f := make([][][]int, len(prices))
	k := 2

	for i := 0; i < len(prices); i++ {
		f[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = make([]int, k+1) // 加1是为确保从1开始, 简化程序
			for h := 0; h <= k; h++ {
				f[i][j][h] = -100000
			}
		}
	}
	f[0][0][0] = 0

	for i := 1; i < len(prices); i++ {
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
	ans := -100000

	for h := 0; h <= k; h++ {
		ans = max(ans, f[len(prices)-1][0][h])
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
// [3,3,5,0,0,3,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

