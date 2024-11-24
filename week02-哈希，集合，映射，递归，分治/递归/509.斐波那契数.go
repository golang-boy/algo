/*
 * @lc app=leetcode.cn id=509 lang=golang
 * @lcpr version=20003
 *
 * [509] 斐波那契数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func fib(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/

