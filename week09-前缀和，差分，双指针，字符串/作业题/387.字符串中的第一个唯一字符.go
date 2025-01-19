/*
 * @lc app=leetcode.cn id=387 lang=golang
 * @lcpr version=20004
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func firstUniqChar(s string) int {

	/*
	   输入字符串s,返回第一个不重复的字符索引

	   使用map进行计数判重, 依次按顺序变量
	*/

	q := map[rune]int{}
	ans := -1

	for _, ch := range s {
		q[ch]++
	}

	for i := 0; i < len(s); i++ {
		if q[rune(s[i])] == 1 {
			ans = i
			break
		}
	}

	return ans

}

// @lc code=end

/*
// @lcpr case=start
// "leetcode"\n
// @lcpr case=end

// @lcpr case=start
// "loveleetcode"\n
// @lcpr case=end

// @lcpr case=start
// "aabb"\n
// @lcpr case=end

*/

