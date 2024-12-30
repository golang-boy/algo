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

	c := 2
	for i := range f {
		f[i] = make([][]int, 2)
		for j := range f[i] {
			f[i][j] = make([]int, c+1)
			for k := 0; k <= c; k++ {
				f[i][j][k] = -1000
			}
		}
	}

	f[0][0][0] = 0

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k <= c; k++ {

				if k > 0 {
					// 买
					f[i][1][k] = max(f[i][1][k], f[i-1][0][k-1]-prices[i])
				}
				// 卖
				f[i][0][k] = max(f[i][0][k], f[i-1][1][k]+prices[i])

				// 不做
				f[i][j][k] = max(f[i][j][k], f[i-1][j][k])
			}
		}
	}

	ans := -1000

	for i := 0; i <= c; i++ {
		// 取最后一位时，不持仓情况下的某一笔的最大值
		ans = max(f[len(prices)-1][0][i], ans)
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

