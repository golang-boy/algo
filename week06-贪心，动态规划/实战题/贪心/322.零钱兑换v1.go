/*
 * @Author: 刘慧东
 * @Date: 2024-12-25 09:52:13
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-12-25 09:52:14
 */
/*
 * @lc app=leetcode.cn id=322 lang=golang
 * @lcpr version=20004
 *
 * [322] 零钱兑换
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func coinChange(coins []int, amount int) int {

	/*

		能不能二分？
		    能否确定最大值和最小值俩个边界,如果不能就不能使用二分

		贪心: 每次取最优决策, 该题中先尽可能取大

		取局部最优不能得到正确的解

		使用dfs搜索
		    组合问题, 每次选取时，可以是三个中的任何一个
			amount是状态，索引不是状态

		代码如何表示?
		    本题中看，需要使用递归，递归最大的，非叶子尝试选取其他, 递归叶子

		    变量： amount，coins的索引
	*/

	ans := math.MaxInt
	chosen := []int{}

	dfs(amount, coins, chosen, &ans)

	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func dfs(amount int, coins []int, chosen []int, ans *int) {
	if amount == 0 {
		// 统计结果
		if *ans > len(chosen) {
			*ans = len(chosen)
		}
		return
	}

	if amount < 0 {
		return
	}

	for i := 0; i < len(coins); i++ {
		chosen = append(chosen, coins[i])
		dfs(amount-coins[i], coins, chosen, ans)

		chosen = chosen[:len(chosen)-1]
	}
}

/*
总结：
    执行超时并不是算法的问题，是因为coins太大，
    导致执行时间过长，因此dfs搜索的方式不行
*/

// @lc code=end

/*
// @lcpr case=start
// [1, 2, 5]\n11\n
// @lcpr case=end

// @lcpr case=start
// [1, 2, 5]\n100\n
// @lcpr case=end

// @lcpr case=start
// [2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/

