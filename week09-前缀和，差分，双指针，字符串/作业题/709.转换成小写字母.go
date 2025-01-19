/*
 * @lc app=leetcode.cn id=709 lang=golang
 * @lcpr version=20004
 *
 * [709] 转换成小写字母
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func toLowerCase(ss string) string {

	s := []byte(ss)

	for i := 0; i < len(s); i++ {

		if s[i] >= 'A' && s[i] <= 'Z' {
			s[i] = s[i] - 'A' + 'a'
		}
	}
	return string(s)
}

/*
总结：
	时间复杂度为O(n)
	空间复杂度为O(n)

	核心关键为：s[i] = s[i] - 'A' + 'a'

*/

// @lc code=end

/*
// @lcpr case=start
// "Hello"\n
// @lcpr case=end

// @lcpr case=start
// "here"\n
// @lcpr case=end

// @lcpr case=start
// "LOVELY"\n
// @lcpr case=end

*/

