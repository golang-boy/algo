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

	n := len(s)
	m := len(p)

	f := make([][]bool, len(s)+1)
	for i := range f {
		f[i] = make([]bool, len(p)+1)
	}

	s = " " + s
	p = " " + p

	f[0][0] = true

	for j := 2; j <= m; j += 2 {
		if p[j] == '*' {
			f[0][j] = true
		} else {
			break
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j] >= 'a' && p[j] <= 'z' {
				f[i][j] = f[i-1][j-1] && s[i] == p[j]
			} else if p[j] == '.' {
				f[i][j] = f[i-1][j-1]
			} else {
				// p[j] == '*' 情况, 可以匹配前面s[i-1]位置任何值，同时需要参考更前面的是否匹配，更前面的为f[i][j-2]
				f[i][j] = f[i][j-2]

				// p[j] 为*时，如果p[j-1] 为. (此时f[i][j] = f[i][j-2]) 或者为字母等于s[i]时，此时也匹配，同时要参考更前面i-1位置的
				if s[i] == p[j-1] || p[j-1] == '.' {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			}

		}
	}

	return f[n][m]
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

