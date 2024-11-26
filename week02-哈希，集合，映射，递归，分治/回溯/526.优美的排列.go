
/*
 * @lc app=leetcode.cn id=526 lang=golang
 * @lcpr version=20003
 *
 * [526] 优美的排列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countArrangement(n int) int {

	/*

		优美排列：  perm[i] % i = 0,   i % perm[i] == 0   索引从1开始

		值与索引相互整除

		输入n，输出满足条件的子集数量, 全排列

		思路：
		  1. 用n个数，构造子集
		  2. 判断是否满足条件
		  3. 结果+1
	*/

	var ans = 0
	var chosen = []int{}
	var used = make([]bool, n)

	r(1, n, &ans, chosen, used)

	return ans

}

func r(i int, n int, ans *int, chosen []int, used []bool) {
	if i == n+1 {
		if len(chosen) != 0 && isPerm(chosen) {
			*ans++
		}
		return
	}

	if len(chosen) != 0 && !isPerm(chosen) {
		return
	}
	for j := 1; j < n+1; j++ {

		if !used[j-1] {
			chosen = append(chosen, j)
			used[j-1] = true
			// 选i时,递归
			r(i+1, n, ans, chosen, used)
			chosen = chosen[:len(chosen)-1]
			used[j-1] = false
		}
	}
}

func isPerm(a []int) bool {
	ans := true
	for i := 0; i < len(a); i++ {
		j := i + 1
		// i, j 不为0
		if !(a[i]%j == 0 || j%a[i] == 0) {
			return false
		}
	}
	return ans
}

/*
总结：
     审题不清楚： 题目和示例，明确让找排列，结果找成子集
	 时间消耗 30分钟
*/

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

