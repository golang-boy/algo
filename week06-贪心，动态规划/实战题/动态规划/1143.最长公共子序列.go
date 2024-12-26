/*
 * @lc app=leetcode.cn id=1143 lang=golang
 * @lcpr version=20004
 *
 * [1143] 最长公共子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {

	/*
		     输入俩字符串
			 输出公共子序列的长度

	*/

	m := len(text1)
	n := len(text2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
总结：
      思路一：

	        二维表格中，最长公共子序列


*/

// @lc code=end

/*
// @lcpr case=start
// "abcde"\n"ace"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"def"\n
// @lcpr case=end

*/

