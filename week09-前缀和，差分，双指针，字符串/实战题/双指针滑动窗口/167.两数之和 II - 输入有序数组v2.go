/*
 * @lc app=leetcode.cn id=167 lang=golang
 * @lcpr version=20004
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func twoSum(numbers []int, target int) []int {
	/*
		   1. 下标从1开始索引
		   2. 不严格递增,非递减序列
		   3. 返回满足条件的俩索引

		   朴素思路：
		   		两层循环，暴力搜索枚举

		   优化思路，去除内层循环
				1. 固定右端点，看左端点的取值范围
				2. 移动一个端点，看另一个的变化
	*/

	n := len(numbers)
	j := n - 1
	for i := 0; i < n; i++ {
		// 移动一个端点，看另一个的变化
		// 右边开始移动,大于时一直右移
		for ; j > i && numbers[i]+numbers[j] > target; j-- {
		}

		if j > i && numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
		// 小于时，i左移
	}
	return []int{}
}

/*
总结：
	两层循环，想的是取优化内层循环，但实际做题时还是暴力

	为什么比上个版本更优？
		优不优看循环内外层循环执行了多少次.

		v1版本，内层循环j对于每个i都会进行判断，直到找到


24/24 cases passed (0 ms)
Your runtime beats 100 % of golang submissions
Your memory usage beats 40.63 % of golang submissions (7.6 MB)
*/

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [-1,0]\n-1\n
// @lcpr case=end

*/

