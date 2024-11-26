/*
 * @Author: 刘慧东
 * @Date: 2024-11-26 18:03:18
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-26 18:03:18
 */

/*
 * @lc app=leetcode.cn id=面试题 16.18 lang=golang
 * @lcpr version=20003
 *
 * [面试题 16.18] 模式匹配
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func patternMatching(pattern string, value string) bool {

	q := map[string]string{}
	set := map[string]bool{}

	return r(pattern, value, set, q)
}

func r(p, v string, set map[string]bool, q map[string]string) bool {
	if p == "" {
		return v == ""
	}

	char := string(p[0])
	if w, ok := q[char]; ok {
		if !strings.HasPrefix(v, w) {
			return false
		}
		return r(p[1:], v[len(w):], set, q)
	}

	for i := -1; i < len(v); i++ {
		word := v[:i+1]
		if !set[word] {
			set[word] = true
			q[char] = word
			if r(p[1:], v[i+1:], set, q) {
				return true
			}
			delete(set, word)
			delete(q, char)
		}
	}
	return false

}

/*

总结：和原始思路一样，只是v1版本做的过程中，陷入了嵌套模板的过程

*/

// @lc code=end

/*
// @lcpr case=start
// "abba"\n"dogcatcatdog"\n
// @lcpr case=end

// @lcpr case=start
// "abba"\n"dogcatcatfish"\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n"dogcatcatdog"\n
// @lcpr case=end

// @lcpr case=start
// "abba"\n"dogdogdogdog"\n
// @lcpr case=end

*/

