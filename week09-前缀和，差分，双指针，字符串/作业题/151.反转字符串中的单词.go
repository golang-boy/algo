/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=20004
 *
 * [151] 反转字符串中的单词
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reverseWords(s string) string {

	st := []string{}

	temp := []rune{}
	for _, ch := range s {
		if ch == ' ' {
			if len(temp) != 0 {
				st = append(st, string(temp))
				temp = []rune{}
			}
		} else {
			temp = append(temp, ch)
		}
	}
	if len(temp) != 0 {
		st = append(st, string(temp))
	}

	ans := ""
	for len(st) != 0 {
		ans += st[len(st)-1]
		st = st[:len(st)-1]
		if len(st) != 0 {
			ans += " "
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/

