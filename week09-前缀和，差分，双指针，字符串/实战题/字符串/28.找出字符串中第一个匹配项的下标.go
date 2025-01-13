/*
 * @lc app=leetcode.cn id=28 lang=golang
 * @lcpr version=20004
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func strStr(haystack string, needle string) int {
	/*
		rabin-karp基于hash的高效字符串搜索算法

		定长n字符串s，长度m模式串t, 求t在s中出现

		计算s的每个长度为m的子串的hash值, 宽度为m的滑动窗口滑过s

		检测与t的hash值是否相等

	*/

	/*

		输入原始字符串，和模式子串
		输出模式子串在原字符串中的第一个位置索引

		朴素的思路:

			循环遍历原字符串的每个字符，在子串中依次匹配

			for i:=0;i<len(haystack);i++{
				// [l,r], 内部循环依次扫描比较
			}
	*/

	n := len(haystack)
	m := len(needle)

	ans := -1

	if n < m {
		return ans
	}

	for i := 0; i < n-m+1; i++ {
		l := 0

		for ; l < m && haystack[i+l] == needle[l]; l++ {
		}

		if l == m {
			ans = i
			break
		}
	}

	return ans

}

// @lc code=end

/*
// @lcpr case=start
// "sadbutsad"\n"sad"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n"leeto"\n
// @lcpr case=end

*/

