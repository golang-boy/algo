/*
 * @lc app=leetcode.cn id=58 lang=golang
 * @lcpr version=20004
 *
 * [58] 最后一个单词的长度
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLastWord(s string) int {

	/*
		倒序遍历字符串，遇到空格统计一次长度
	*/

	ans := 0

	ss := []byte(s)
	i := len(s) - 1

	for ; i >= 0 && ss[i] == ' '; i-- {
	}

	for ; i >= 0 && ss[i] != ' '; i-- {
		ans++
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "Hello World"\n
// @lcpr case=end

// @lcpr case=start
// "   fly me   to   the moon  "\n
// @lcpr case=end

// @lcpr case=start
// "luffy is still joyboy"\n
// @lcpr case=end

*/

