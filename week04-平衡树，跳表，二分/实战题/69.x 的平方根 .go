/*
 * @lc app=leetcode.cn id=69 lang=golang
 * @lcpr version=20004
 *
 * [69] x 的平方根
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func mySqrt(x int) int {
	/*

	   ans * ans = x

	   ans = x /ans
	*/

	left := 1
	right := x

	for left < right {
		mid := (left + right + 1) / 2

		// 这是需要找最中间的数字，
		// 找到后判断是否 mid*mid<=x,
		// 数字的范围在1-x中间
		if mid*mid <= x {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right

}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 8\n
// @lcpr case=end

*/

