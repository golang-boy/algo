/*
 * @lc app=leetcode.cn id=673 lang=golang
 * @lcpr version=20004
 *
 * [673] 最长递增子序列的个数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findNumberOfLIS(nums []int) int {

	/*
		f[i] 在i位置时，最长子序列的长度
		cnt[i] 最长子序列的个数

		f[0] = 1
		cnt[0] = 1

		if nums[i] > nums[i-1] {
		    f[i]++
		} else {
		    f[i] = max(f[i-1], f[j] + 1 in [0,i-2] )
			// 1,3,5,4,2,7
		}
	*/

	cnt := make([]int, len(nums))
	f := make([]int, len(nums))

	f[0] = 1
	cnt[0] = 1

	if len(nums) == 1 {
		return 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			f[i]++
			cnt[i] = cnt[i-1]
		} else {
			j := i - 1
			for ; j >= 0; j-- {
				if nums[j] < nums[i] {
					break
				}
			}

			if f[i-1] > f[j]+1 {
				f[i] = f[i-1]
				cnt[i] = cnt[i-1]
			} else if f[i-1] == f[j]+1 {
				f[i] = f[i-1]
				cnt[i] = cnt[i-1] + 1
			} else {
				f[i] = f[j] + 1
				cnt[i] = cnt[j]

			}
		}
	}

	ans := 0
	i := 0

	for ; i < len(nums); i++ {
		if ans < f[i] {
			ans = f[i]
		}
	}

	for j := 0; j < len(nums); j++ {
		cnt[j]

	}

	return ans

}

/*
总结：
    思路不太对，条件与状态正确， 决策的时机没想对

	最长序列的计算是从左往右，j 在[0,i] 之间

	动态规划，除了要思路对外，下面的细节也要想清楚
	1. 定义状态
	2. 状态转义方程
	3. 边界
	    循环的范围
	4. 目标
	    如何从状态中获取最终结果

*/

// @lc code=end

/*
// @lcpr case=start
// [1,3,5,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n
// @lcpr case=end

*/

