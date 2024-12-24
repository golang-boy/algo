
/*
 * @lc app=leetcode.cn id=22 lang=golang
 * @lcpr version=20003
 *
 * [22] 括号生成
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func generateParenthesis(n int) []string {

	/*
	  (A)B
	*/

	if n == 0 {
		return []string{""}
	}

	ans := []string{}
	for k := 1; k <= n; k++ {
		a := generateParenthesis(n - k)
		b := generateParenthesis(k - 1)

		for _, aa := range a {
			for _, bb := range b {
				l := "(" + aa + ")" + bb
				ans = append(ans, l)
			}
		}
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

