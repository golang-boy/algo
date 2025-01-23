/*
 * @lc app=leetcode.cn id=10 lang=golang
 * @lcpr version=20004
 *
 * [10] 正则表达式匹配
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isMatch(s string, p string) bool {
	/*
		输入s和p, s为字符串，p为模式串。.表示匹配任一字符，*匹配0个或多个前面的元素
		输出能否匹配

		f[i][j] 在s的前i个字符[0,i-1]是否可以匹配p中前j个字符[0,j-1]


		if p[j-1] == '.' || s[i-1] == p[j-1] {
			f[i][j] = f[i-1][j-1]
		}

		if p[j-1] == '*' {
			f[i][j] = f[i][j - 2] ||
		}




	*/

}

// @lc code=end

/*
// @lcpr case=start
// "aa"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"a*"\n
// @lcpr case=end

// @lcpr case=start
// "ab"\n".*"\n
// @lcpr case=end

*/

