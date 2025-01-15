/*
 * @lc app=leetcode.cn id=125 lang=golang
 * @lcpr version=20004
 *
 * [125] 验证回文串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPalindrome(ss string) bool {

	/*
		输入字符串，输出是否是回文


		双指针，左右开始遍历，直到左右相等结束

	*/

	ans := true

	s := []rune(ss)

	for l, r := 0, len(s)-1; l < r; {

		if s[l] >= 'A' && s[l] <= 'Z' {
			s[l] = s[l] - 'A' + 'a'
		}

		if s[r] >= 'A' && s[r] <= 'Z' {
			s[r] = s[r] - 'A' + 'a'
		}

		if (s[l] < 'a' || s[l] > 'z') && (s[l] < '0' || s[l] > '9') {
			l++
			continue
		}

		if (s[r] < 'a' || s[r] > 'z') && (s[r] < '0' || s[r] > '9') {
			r--
			continue
		}

		if s[l] != s[r] {
			ans = false
			break
		}

		l++
		r--
	}

	return ans

}

// @lc code=end

/*
// @lcpr case=start
// "A man, a plan, a canal: Panama"\n
// @lcpr case=end

// @lcpr case=start
// "race a car"\n
// @lcpr case=end

// @lcpr case=start
// " "\n
// @lcpr case=end

*/

