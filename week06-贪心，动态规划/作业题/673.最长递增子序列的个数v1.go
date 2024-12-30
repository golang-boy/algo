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

		f[i] 表示在位置i时，最长递增子序列的个数
		l[i] 表示位置在i时，最长子序列的长度

		递增是严格的， nums[i] > nums[i-1] 表示递增

		if nums[i] > nums[i] - 1 {

			最长序列长度加一，个数不变
			f[i] = f[i-1]
		}

		如果小于等于时，找到比当前值小的最大的位置，判断当前长度是否是最长的

		f[i] =  max(max([nums[0],nums[i-2]] < nums[i]) + 1,  f[i-1])

	*/

	f := make([]int, len(nums))
	l := make([]int, len(nums))

	f[0] = 1
	l[0] = 1
	for i := 1; i < len(nums); i++ {

		if nums[i] > nums[i-1] {
			l[i] = l[i-1] + 1
			f[i] = f[i-1]
			continue
		}

		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] < nums[i] {
				break
			}
		}

		l[i] = l[j] + 1
		if l[i-1] == l[i] {
			f[i]++
		} else {
			l[i] = max(l[i-1], l[i])
		}

		f[i] = max(f[i-1], f[i])

	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = max(nums[i], ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,5,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n
// @lcpr case=end

*/

