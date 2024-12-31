
/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=20004
 *
 * [55] 跳跃游戏
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func canJump(nums []int) bool {
	/*
		输入nums,非负整数, 每个元素表示从该位置可跳跃的最大长度
		输出是否能到达最后一个下标

		初始位于第一个位置，判断是否能到达最后一个位置

		f[i] 在从0开始,能否到达i位置

		f[0] = true

		f[i] = f[i] || (f[j] && nums[j]+j >= i)

	*/

	nums = append([]int{1}, nums...)

	f := make([]bool, len(nums))

	f[0] = true
	for i := 1; i < len(nums); i++ {

		for j := 0; j < i; j++ {
			// for j := i - 1; j >= 0; j-- {
			// 1. f[i] 递推 进行 ||
			// 2. f[j] && j + nums[j] >= i  前一步到能否到i
			f[i] = f[i] || (f[j] && nums[j]+j >= i)
		}
	}

	return f[len(nums)-1]
}

/*
总结：

	类比背包问题

		nums[i]对应物品

		[1,i) 表示物品的体积或者重量

		假设我们在某一位置 i，我们可以将这个位置视为一个“物品”，
		其“重量”是从这个位置跳跃的距离，而“价值”可以简单地视为我们能否到达下一个位置。


		动态规划，计数问题要达到不重不漏
		最优化问题只要做到不漏就行

*/

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1,0,4]\n
// @lcpr case=end

*/

