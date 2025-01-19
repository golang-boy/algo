
/*
 * @lc app=leetcode.cn id=205 lang=golang
 * @lcpr version=20004
 *
 * [205] 同构字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isIsomorphic(s string, t string) bool {

	/*
		1. 先判断长度是否相同，如果不同返回false
		2. 创建一个hash表, 用索引i遍历s和t
		3. hash表中存储映射关系，找不到存进去，找到取出来比较
		4. 相等则继续，不相等，结束返false
	*/

	if len(s) != len(t) {
		return false
	}

	q := map[byte]byte{}
	p := map[byte]byte{}
	ans := true

	for i := 0; i < len(s); i++ {
		if ch, ok := q[s[i]]; ok {
			if ch != t[i] {
				ans = false
				break
			}
		}

		if ch, ok := p[t[i]]; ok {
			if ch != s[i] {
				ans = false
				break
			}
		}

		q[s[i]] = t[i]
		p[t[i]] = s[i]
	}

	return ans

}

// @lc code=end

/*
// @lcpr case=start
// "egg"\n"add"\n
// @lcpr case=end

// @lcpr case=start
// "foo"\n"bar"\n
// @lcpr case=end

// @lcpr case=start
// "paper"\n"title"\n
// @lcpr case=end

*/

