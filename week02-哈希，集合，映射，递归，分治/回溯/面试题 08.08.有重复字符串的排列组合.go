
/*
 * @lc app=leetcode.cn id=面试题 08.08 lang=golang
 * @lcpr version=20003
 *
 * [面试题 08.08] 有重复字符串的排列组合
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func permutation(S string) []string {

	chars := []byte(S)
	var ans = []string{}
	var used = make([]bool, len(chars))
	var chosen = []byte{}
	r(0, chars, &ans, used, chosen)
	return ans
}

func r(i int, chars []byte, ans *[]string, used []bool, chosen []byte) {

	if i == len(chars) {
		*ans = append(*ans, string(chosen))
		return
	}

	dedup := make(map[byte]bool)

	for j := 0; j < len(chars); j++ {
		if !used[j] {
			if !dedup[chars[j]] {
				dedup[chars[j]] = true

				used[j] = true
				chosen = append(chosen, chars[j])
				r(i+1, chars, ans, used, chosen)
				chosen = chosen[:len(chosen)-1]
				used[j] = false
			}
		}

	}

}

// @lc code=end

/*
// @lcpr case=start
// "qqe"\n
// @lcpr case=end

// @lcpr case=start
// "ab"\n
// @lcpr case=end

*/

