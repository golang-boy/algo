/*
 * @lc app=leetcode.cn id=342 lang=golang
 * @lcpr version=20003
 *
 * [342] 4的幂
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPowerOfFour(n int) bool {

	if n == 1 {
		return true
	}

	if n == 0 {
		return false
	}

	div := n / 4
	rem := n % 4

	if rem != 0 {
		return false
	}

	if rem == 0 && div == 1 {
		return true
	}

	return isPowerOfFour(div)
}

/*
总结：
		 耗时 0:2:42
*/

// @lc code=end

/*
// @lcpr case=start
// 16\n
// @lcpr case=end

// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

