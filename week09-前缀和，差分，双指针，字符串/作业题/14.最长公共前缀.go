/*
 * @lc app=leetcode.cn id=14 lang=golang
 * @lcpr version=20004
 *
 * [14] 最长公共前缀
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestCommonPrefix(strs []string) string {

	/*
		思路：
			1. 将字符放入二维表格，但这样可能会比较稀疏
			2. 取第一个循环在其他字符串中依次判断
	*/

	ans := ""

	for i := range strs[0] {
		isPrefix := true
		for _, str := range strs[1:] {
			if i > len(str)-1 || str[i] != strs[0][i] {
				isPrefix = false
				break
			}
		}
		if isPrefix {
			ans += string(strs[0][i])
		} else {
			break
		}
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// ["flower","flow","flight"]\n
// @lcpr case=end

// @lcpr case=start
// ["dog","racecar","car"]\n
// @lcpr case=end

*/

