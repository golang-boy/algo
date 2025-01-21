/*
 * @lc app=leetcode.cn id=560 lang=golang
 * @lcpr version=20004
 *
 * [560] 和为 K 的子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func subarraySum(nums []int, k int) int {
	/*
		输入整数数组nums和k,返回连续子数组元素和为k的子数组个数

		求取 sum[l,r] = k

		sum[l,r] = sum[r] - sums[l-1] == k

		sum[l-1] = sum[r] - k
	*/

	n := len(nums)

	sum := make([]int, n+1)

	sum[0] = 0
	ans := 0

	for i := 1; i < n+1; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	// for i := 1; i <= n; i++ {
	// 	s := sum[i] - k
	// 	for j := 0; j < i; j++ {
	// 		if s == sum[j] {
	// 			ans++
	// 		}
	// 	}
	// }

	// count用于计数，记录每个位置和为某个值的数量
	// 如果和的值比较小，可以使用索引表示，则使用count比较快速
	// 此处也可以使用map存储和为某个值数量的计数, 一轮循环边计算边统计

	count := make(map[int]int)
	count[sum[0]]++

	for i := 1; i <= n; i++ {
		s := sum[i] - k
		ans += count[s] // 在count查找计数
		count[sum[i]]++
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

*/

