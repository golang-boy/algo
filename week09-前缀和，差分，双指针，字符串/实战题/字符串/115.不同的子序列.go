
/*
 * @lc app=leetcode.cn id=115 lang=golang
 * @lcpr version=20004
 *
 * [115] 不同的子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numDistinct(s string, t string) int {

	/*
		f[i][j] t中第j字符前的子串出现在s中前i字符的个数

		_rabbbit   _rabbit

		f[i][0] = 1
		f[1][1] = f[0][0]     if s[1] == t[1]
		f[2][2] = f[1][1]     if s[2] == t[2]
		f[3][3] = f[2][2]     if s[3] == t[3]
		f[4][4] =


	*/

}

// @lc code=end

/*
// @lcpr case=start
// "rabbbit"\n"rabbit"\n
// @lcpr case=end

// @lcpr case=start
// "babgbag"\n"bag"\n
// @lcpr case=end

*/

