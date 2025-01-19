/*
 * @lc app=leetcode.cn id=438 lang=golang
 * @lcpr version=20004
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findAnagrams(s string, p string) []int {

	/*

		1. 使用滑动窗口方法, 使用map统计p中每个字符的频率
		2. 异位词中每个字符频率都是相等的，这样只需要计算一次，即可进行比较


	*/
	ans := []int{}

	if len(s) < len(p) {
		return ans
	}

	q := map[byte]int{}
	sq := map[byte]int{}

	for i := 0; i < len(p); i++ {
		q[p[i]]++
		sq[s[i]]++
	}

	if calc(sq, q) {
		ans = append(ans, 0)
	}

	for l := len(p); l < len(s); l++ {
		sq[s[l]]++
		sq[s[l-len(p)]]--

		if sq[s[l-len(p)]] == 0 {
			delete(sq, s[l-len(p)])
		}

		if calc(sq, q) {
			ans = append(ans, l-len(p)+1)
		}

	}

	return ans
}

func calc(s map[byte]int, q map[byte]int) bool {
	if len(s) != len(q) {
		return false
	}

	for k, v := range s {
		if q[k] != v {
			return false
		}
	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// "cbaebabacd"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abab"\n"ab"\n
// @lcpr case=end

*/

