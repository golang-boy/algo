/*
 * @lc app=leetcode.cn id=344 lang=golang
 * @lcpr version=20004
 *
 * [344] 反转字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reverseString(s []byte) {

	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// @lc code=end

/*
// @lcpr case=start
// ["h","e","l","l","o"]\n
// @lcpr case=end

// @lcpr case=start
// ["H","a","n","n","a","h"]\n
// @lcpr case=end

*/

