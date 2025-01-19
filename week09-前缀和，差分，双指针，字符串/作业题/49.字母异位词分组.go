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

	/*
		将给定数组中的字符串的异位词进行分组
	*/
	q := map[string][]string{}

	for i := range strs {
		key := sortString(strs[i])
		if _, ok := q[key]; ok {
			q[key] = append(q[key], strs[i])
		} else {
			q[key] = []string{strs[i]}
		}
	}

	ans := [][]string{}

	for _, v := range q {
		ans = append(ans, v)
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

