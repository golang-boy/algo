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
			f[i][j] = f[i][j - 2] || (f[i-1][j] && p[j-2] == s[i-1]) || p[j-2] == '.'
		}

		f[0][0]  = ( s[0] == p[0]  || p[0] == '.' )

		f[1][1]  = f[0][0]  ||  if　p[1] == '*'  p[0] == s[1] || if p[1] == '.'

	*/

	f := make([][]bool, len(s)+1)
	for i := range f {
		f[i] = make([]bool, len(p)+1)
	}

	f[0][0] = true

	for j := 1; j <= len(p)+1; j++ {
		if p[j-1] == '*' {
			f[0][j] = f[0][j-2]
		}
	}

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

