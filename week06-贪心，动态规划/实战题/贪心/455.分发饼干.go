/*
 * @lc app=leetcode.cn id=455 lang=golang
 * @lcpr version=20004
 *
 * [455] 分发饼干
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findContentChildren(g []int, s []int) int {

	/*
		输入：孩子数组，元素代表胃口值, 饼干数组，元素代表饼干对应的满足胃口值
		输出：满足孩子胃口值的最大数值

		1. 饼干和孩子数组先排序
		2. 取最大，进行判断，满足结果+1，同时移动
		2. 不满足，孩子移动，饼干不移动

	*/

	sort.Ints(g)
	sort.Ints(s)

	ans := 0
	i := len(g) - 1
	j := len(s) - 1
	for i >= 0 && j >= 0 {
		if s[j] >= g[i] {
			i--
			j--
			ans++
		} else {
			i--
		}
	}

	return ans

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n[1,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[1,2,3]\n
// @lcpr case=end

*/

