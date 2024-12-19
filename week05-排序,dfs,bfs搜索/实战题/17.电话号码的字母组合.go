/*
 * @lc app=leetcode.cn id=17 lang=golang
 * @lcpr version=20004
 *
 * [17] 电话号码的字母组合
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func letterCombinations(digits string) []string {

	/*
		输入数字，返回对应的字母的组合列表

		1. 建立数字与字母的对应关系
		2. 遍历输入的数字,找到对应的字母集合
		3. 依次从不同的数组中取字母,获取组合


		输入23不变，
		依次考虑哪个数字对应的字母
		     变量为index，字符串为变量
		该俩变量为状态量
	*/

	ans := []string{}
	str := []byte{}
	dfs(0, digits, str, &ans)
	return ans

}

func dfs(index int, digits string, str []byte, ans *[]string) {
	if index == len(digits) {
		if len(str) != 0 {
			*ans = append(*ans, string(str))
		}
		return
	}

	for _, ch := range []byte(alphabet[digits[index]]) {

		// 一种是str为共享状态，每次append完后，需要恢复状态
		// str = append(str, ch)
		// dfs(index+1, digits, str, ans)
		// str = str[:len(str)-1]

		// 另一种是传参时，str不变，每次都是一个新的
		dfs(index+1, digits, append(str, ch), ans)
	}

}

var alphabet = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// @lc code=end

/*
// @lcpr case=start
// "23"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

// @lcpr case=start
// "2"\n
// @lcpr case=end

*/

