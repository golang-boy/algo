/*
 * @lc app=leetcode.cn id=70 lang=golang
 * @lcpr version=20004
 *
 * [70] 爬楼梯
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func climbStairs(n int) int {
	/*


	   1 <= i <= n

	   f(i) 到达第i个台阶有多少种方法

	   f(i) = f(i - 1) + 1
	   f(i) = f(i - 2) + 2

	   f(i) = f(i - 2) + f(i - 1)

	*/

	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	f := make([]int, n+1)

	f[0] = 1
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/

