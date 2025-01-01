
/*
 * @lc app=leetcode.cn id=312 lang=golang
 * @lcpr version=20004
 *
 * [312] 戳气球
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxCoins(nums []int) int {

	/*
		输入nums，戳破第i个气球，可获得 g(i) = nums[i-1]*nums[i]*nums[i+1]这么多硬币
		如果i-1或i+1越界，当做数字为1的气球
		输出, 戳破所有气球时，获得硬币的最大数量

			ans := max(g[i],g[j])

		题目看不懂

		哦，戳破后就没有了。示例1[3,1,5,8]中，
		  戳破1，变为[3,5,8]
		  戳破5，变为[3,8]
		  戳破3，变为[8]
		  戳破8，结束,得最大

		戳破后，左右发送变化

		看能否分出子问题?
			1. 任何时候，都是要戳破一段区间的气球
			2. 戳破一个后，同样处理剩余的气球
			3. 朴素的方法就是递归，但nums变化后，每个位置上的都要进行尝试，空间复杂度太高

		戳破p气球后，子问题[l, p-1], [p+1, r]与[l,r]不是同类子问题
		先戳破[l,p-1]和[p+1,r], 最后戳破p, 满足同类子问题

		定义：
			f[l][r]表示戳破区间l->r之间所有的气球，获取的最大数量

		转移方程：
			f[l][r] max(f[l][p-1] + f[p+1][r] + nums[p] * nums[l-1]*nums[r+1])   p in [l,r]

		目标：
			f[1][n]

	*/

	n := len(nums)
	// 前后各增加一个数
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	f := make([][]int, n+2)
	for i := range f {
		f[i] = make([]int, n+2)
	}

	// 区间动态规划，最外层循环区间长度
	for length := 1; length <= n; length++ {

		// 此处要处理长度为length的区间，算出左右端点开始循环，l的范围 [1,n+1-length)], 循环左端点
		for l := 1; l <= n+1-length; l++ {
			// length固定情况下，左端点循环右移，计算f[l][r],即从l到r区间的硬币数量
			// length的移动，表示右端点移动, 综合看，每次右端点移动时，左端点所有结果已经计算完毕

			// 计算出右端点
			r := l + length - 1
			for p := l; p <= r; p++ {
				f[l][r] = max(
					f[l][r],
					f[l][p-1]+f[p+1][r]+nums[p]*nums[l-1]*nums[r+1],
				)
			}
		}
	}

	return f[1][n]

}

/*
总结：
	区间规划是一些特殊类型的题目，这种解题思路需要了解掌握

*/

// @lc code=end

/*
// @lcpr case=start
// [3,1,5,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,5]\n
// @lcpr case=end

*/

