/*
 * @Author: 刘慧东
 * @Date: 2024-12-30 13:51:28
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-12-30 13:57:36
 */

/*
 * @lc app=leetcode.cn id=309 lang=golang
 * @lcpr version=20004
 *
 * [309] 买卖股票的最佳时机含冷冻期
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int) int {

	prices = append([]int{0}, prices...)
	f := make([][][]int, 2)

	for i := range f {
		f[i] = make([][]int, 2)
		for j := range f[i] {
			f[i][j] = make([]int, 2)
			for k := range f[i][j] {
				f[i][j][k] = -1000000
			}
		}
	}

	f[0][0][0] = 0

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {

				f[i&1][0][0] = -1000000
				f[i&1][0][1] = -1000000
				f[i&1][1][0] = -1000000
				f[i&1][1][1] = -1000000

				//  0的时候不在冷冻期，1时在冷冻期
				// 买入
				f[i&1][1][0] = max(f[i&1][1][0], f[i-1&1][0][0]-prices[i])

				// 卖出
				f[i&1][0][1] = max(f[i&1][0][1], f[i-1&1][1][0]+prices[i])

				f[i&1][j][0] = max(f[i&1][j][0], f[i-1&1][j][k])
			}
		}
	}

	// 取冷冻期和非冷冻期的最大值

	return max(f[len(prices)-1&1][0][1], f[len(prices)-1&1][0][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
总结：

    上述这种方式使用滚动数组不行，只能使用表格优化后的方式？


*/

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

