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

	 */

	hash := sortString(p)
	ans := []int{}

	for i := 0; i < len(s)-len(p)+1; i++ {
		if hash == sortString(s[i:i+len(p)]) {
			ans = append(ans, i)
		}
	}
	return ans
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

/*
总结：
	每次移动循环计算，得到hash值，然后与p的hash进行比较
	hash值计算方式是排序，排序后异位词都一样

	方法虽然可以，但时间复杂度太高 n*m*log(m)

*/

// @lc code=end

/*
// @lcpr case=start
// "cbaebabacd"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abab"\n"ab"\n
// @lcpr case=end

*/

