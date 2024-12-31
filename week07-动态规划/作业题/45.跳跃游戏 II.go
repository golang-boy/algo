/*
 * @lc app=leetcode.cn id=45 lang=golang
 * @lcpr version=20004
 *
 * [45] 跳跃游戏 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func jump(nums []int) int {

	/*
		f[i] 到i需要的最小跳跃次数

		f[i] = min(f[i], f[j] >= i)
	*/

	nums = append([]int{1}, nums...)

	f := make([]int, len(nums))
	for i := range f {
		f[i] = 100000
	}

	f[0] = 0

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}

	return f[len(nums)-1] - 1
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,0,1,4]\n
// @lcpr case=end

*/

