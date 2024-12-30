/*
 * @lc app=leetcode.cn id=416 lang=golang
 * @lcpr version=20004
 *
 * [416] 分割等和子集
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func canPartition(nums []int) bool {

	sum := 0

	for i := range nums {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	/*
		子集问题，每个数为一个物品，选出来的为总体积一半的结果

		f[i][j] 前i个数选出一些数，总和为j，是否可行

		f[i][j] = f[i-1][j-nums] || f[i-1][j]
	*/

	nums = append([]int{0}, nums...)
	f := make([]bool, sum+1)
	f[0] = true

	for i := 1; i < len(nums); i++ {
		for j := sum / 2; j >= nums[i]; j-- {

			// f[j] 与 f[j-nums[i]] 进行或运算，
			// 与以往不同在于此处不是求取最值, 而是求取是否可行
			//    因此，只要有一个条件满足即可
			f[j] = f[j] || f[j-nums[i]]
		}
	}
	return f[sum/2]
}

// @lc code=end

/*
// @lcpr case=start
// [1,5,11,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,5]\n
// @lcpr case=end

*/

