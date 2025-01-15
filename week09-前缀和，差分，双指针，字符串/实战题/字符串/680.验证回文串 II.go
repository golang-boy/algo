/*
 * @lc app=leetcode.cn id=680 lang=golang
 * @lcpr version=20004
 *
 * [680] 验证回文串 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func validPalindrome(s string) bool {

	ans := false

	for l, r := 0, len(s)-1; l < r; {
		if s[l] != s[r] {
			return ans || check(l+1, r, s, 1) || check(l, r-1, s, 1)
		}
		l++
		r--
	}
	return true
}

func check(l, r int, s string, times int) bool {
	if times == 0 {
		return false
	}
	for l < r {
		if s[l] != s[r] {
			return false || check(l+1, r, s, 0) || check(l, r-1, s, 0)
		}

		l++
		r--
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "aba"\n
// @lcpr case=end

// @lcpr case=start
// "abca"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n
// @lcpr case=end

*/

