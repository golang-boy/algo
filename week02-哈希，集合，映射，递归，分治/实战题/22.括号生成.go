/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {

	/*

	 关键信息：
	   输入: 3
	   输出: ["((()))","(()())","(())()","()(())","()()()"]

	  组合问题, 但是如何写程序？

	*/

	if n == 0 {
		return []string{""}
	}

	var res = []string{}

	for i := 1; i <= n; i++ {
		a := generateParenthesis(i - 1)
		b := generateParenthesis(n - i)

		for _, v1 := range a {
			for _, v2 := range b {
				res = append(res, "("+v1+")"+v2)
			}
		}

	}

	return res
}

// @lc code=end

