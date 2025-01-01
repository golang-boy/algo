
/*
 * @lc app=leetcode.cn id=918 lang=golang
 * @lcpr version=20004
 *
 * [918] 环形子数组的最大和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	/*
		输入nums为环形整数数组
		 	nums[i] 的下一个元素是 nums[(i+1)%n], 前一个元素nums[(i-1+n)%n]
			每个子集中，每个元素只能一次

		输出该数组中非空子数组的最大可能的和

		疑问：
			为什么非要环形数组？普通数组的子数组最大和不行？环形的意义在哪里？
				3+5 与 5+3一样

		思路：
			给定的nums中，选i与不选i，计算值比较大小，递归求子集再计算看起来能解该题

			动态规划呢？

	*/

}

/*
	总结:
		1. 最大子序和的进阶版
		2. 涉及到前缀和
		3. 在i-n <= j <= i-1范围内，找满足条件的最小值, 即滑动窗口求最小值的问题
		4. 使用单调队列 + 前缀和
		5. 问题需要转换为模型

*/

// @lc code=end

/*
// @lcpr case=start
// [1,-2,3,-2]\n
// @lcpr case=end

// @lcpr case=start
// [5,-3,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,-2,2,-3]\n
// @lcpr case=end

*/

