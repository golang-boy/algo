/*
 * @lc app=leetcode.cn id=面试题 08.07 lang=golang
 * @lcpr version=20003
 *
 * [面试题 08.07] 无重复字符串的排列组合
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func permutation(S string) []string {

	/*
	   排列问题
	*/

	var ans = []string{}
	var chosen = []byte{}

	var chars = []byte(S)
	var used = make([]bool, len(chars))

	r(0, chars, &ans, chosen, used)
	return ans
}

func r(i int, chars []byte, ans *[]string, chosen []byte, used []bool) {
	if i == len(chars) {
		// temp := make([]byte, len(chosen))
		// copy(temp, chosen)
		*ans = append(*ans, string(chosen))
		return
	}

	for j := 0; j < len(chars); j++ {
		if !used[j] {
			chosen = append(chosen, chars[j])
			used[j] = true
			r(i+1, chars, ans, chosen, used)
			chosen = chosen[:len(chosen)-1]
			used[j] = false
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// "qwe"\n
// @lcpr case=end

// @lcpr case=start
// "ab"\n
// @lcpr case=end

*/

