/*
 * @lc app=leetcode.cn id=917 lang=golang
 * @lcpr version=20004
 *
 * [917] 仅仅反转字母
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reverseOnlyLetters(ss string) string {
	/*
	 */

	l, r := 0, len(ss)-1

	s := []byte(ss)
	for l < r {

		if !isLetter(s[l]) {
			l++
			continue
		}
		if !isLetter(s[r]) {
			r--
			continue
		}
		s[l], s[r] = s[r], s[l]

		l++
		r--
	}
	return string(s)
}

func isLetter(s byte) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z')
}

// @lc code=end

/*
// @lcpr case=start
// "ab-cd"\n
// @lcpr case=end

// @lcpr case=start
// "a-bC-dEf-ghIj"\n
// @lcpr case=end

// @lcpr case=start
// "code-Q!"\n
// @lcpr case=end

*/

