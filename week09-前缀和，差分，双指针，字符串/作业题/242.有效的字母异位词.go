/*
 * @lc app=leetcode.cn id=242 lang=golang
 * @lcpr version=20004
 *
 * [242] 有效的字母异位词
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isAnagram(s string, t string) bool {

	/*
		异位词指字母相同构成的词，顺序不同

		思路：
			1. 按字母序排序，比较排序后的是否相等
			2. 俩集合分别放入map中，判断集合是否相等

	*/

	q := map[byte]int{}

	if len(t) != len(s) {
		return false
	}

	ans := true
	for i := range s {
		q[s[i]]++
	}

	for i := range t {
		if _, ok := q[t[i]]; ok {
			q[t[i]]--
		} else {
			q[t[i]] = 0
		}
	}

	for _, v := range q {
		if v != 0 {
			ans = false
			break
		}
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// "anagram"\n"nagaram"\n
// @lcpr case=end

// @lcpr case=start
// "rat"\n"car"\n
// @lcpr case=end

*/

