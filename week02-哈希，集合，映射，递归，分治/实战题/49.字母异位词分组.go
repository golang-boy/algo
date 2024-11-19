
/*
 * @lc app=leetcode.cn id=49 lang=golang
 * @lcpr version=20003
 *
 * [49] 字母异位词分组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func groupAnagrams(strs []string) [][]string {

	// 输出一个二维数组
	ans := [][]string{}

	query := map[string][]string{}

	sort := func(s string) string {
		chs := []byte(s)

		sort.Slice(chs, func(i, j int) bool {
			return chs[i] < chs[j]
		})
		return string(chs)
	}

	for _, str := range strs {

		sortedStr := sort(str)
		if l, ok := query[sortedStr]; ok {
			query[sortedStr] = append(l, str)
		} else {
			query[sortedStr] = []string{str}
		}
	}

	for _, l := range query {
		ans = append(ans, l)
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// ["eat", "tea", "tan", "ate", "nat", "bat"]\n
// @lcpr case=end

// @lcpr case=start
// [""]\n
// @lcpr case=end

// @lcpr case=start
// ["a"]\n
// @lcpr case=end

*/

