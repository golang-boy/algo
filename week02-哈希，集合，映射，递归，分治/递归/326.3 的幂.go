/*
 * @lc app=leetcode.cn id=326 lang=golang
 * @lcpr version=20003
 *
 * [326] 3 的幂
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPowerOfThree(n int) bool {
	/*

	 如何使用递归来解？
	    n能被3除尽，则为true, 除不尽则为false
	*/

	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	div := n / 3
	rem := n % 3

	if rem != 0 {
		return false
	}

	if rem == 0 && div == 1 {
		return true
	}
	return isPowerOfThree(div)
}

/*
  总结：
     先判断当前是否满足条件，如果不满足，进行递归判断
	 满足直接返回,无须递归压栈
	 耗时 0:9:6

*/

// @lc code=end

/*
// @lcpr case=start
// 27\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

// @lcpr case=start
// 9\n
// @lcpr case=end

// @lcpr case=start
// 45\n
// @lcpr case=end

*/

