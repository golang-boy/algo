/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=20004
 *
 * [5] 最长回文子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestPalindrome(s string) string {

	n := len(s)
	ansLen := 0
	ansStart := 0

	for c := 0; c < n; c++ {
		l, r := c-1, c+1

		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}

		if r-l-1 > ansLen {
			ansLen = r - l - 1
			ansStart = l + 1
		}
	}

	for c := 1; c < n; c++ {
		l, r := c-1, c

		// 以间隙为中心时，左右相等则往外扩散继续判断
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		// 停止时，说明l和r到达不是回文串的位置
		// 因此，[l+1, r-1]是以c-1和c间的空为中心的回文子串
		// 子串长度为(r-1)-(l+1) + 1 = r-l-1
		if r-l-1 > ansLen {
			ansLen = r - l - 1
			ansStart = l + 1
		}
	}

	return s[ansStart : ansStart+ansLen]

}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

