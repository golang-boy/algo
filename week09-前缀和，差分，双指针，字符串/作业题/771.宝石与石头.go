/*
 * @lc app=leetcode.cn id=771 lang=golang
 * @lcpr version=20004
 *
 * [771] 宝石与石头
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numJewelsInStones(jewels string, stones string) int {
	/*
		jewels表示宝石类型，stones表示拥有的石头类型

		判断石头中有哪些是宝石，返回宝石数量

		区分大小写

		问题的目标是在stones中找到所有在jewels中的字符，并计数
	*/

	q := map[rune]int{}

	for _, ch := range jewels {
		q[ch] = 0
	}

	for _, ch := range stones {
		if _, ok := q[ch]; ok {
			q[ch]++
		}
	}

	ans := 0
	for _, v := range q {
		ans += v
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "aA"\n"aAAbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "z"\n"ZZ"\n
// @lcpr case=end

*/

