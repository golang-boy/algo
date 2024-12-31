

/*
 * @lc app=leetcode.cn id=279 lang=golang
 * @lcpr version=20004
 *
 * [279] 完全平方数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numSquares(n int) int {

	/*
		输入一个整数n,
		输出满足和为n的完全平方数的最少数量

		即，完全平方数的和为n, 1,4,9,16都是完全平方数

		问题转换为，给定无限长的完全平方数数组，选合适的数相加，使其和为n的最少的完全平方数数量

		完全背包问题 + 计数

		最大最小时为max | min, 计数时，为+=

		完全平方数 ==> 物品,  和为n对应包的体积， 价值为1

	*/

	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt
	}

	// f[j]  和为j时，需要最少的完全平方数的最少数量

	f[0] = 0

	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			f[j] = min(f[j], f[j-i*i]+1)
		}
	}

	return f[n]
}

/*
总结：
    完全背包问题没错，但过于套模板，边界始终不知道怎么确定，看了题解
	另外关于计数问题，过于纠结示例，还是需要实践与理论结合。
	真正理解题目后，在套模板，会事半功倍

	边界不好确定
*/

// @lc code=end

/*
// @lcpr case=start
// 12\n
// @lcpr case=end

// @lcpr case=start
// 13\n
// @lcpr case=end

*/

