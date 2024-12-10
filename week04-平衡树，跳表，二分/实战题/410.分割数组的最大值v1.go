/*
 * @lc app=leetcode.cn id=410 lang=golang
 * @lcpr version=20004
 *
 * [410] 分割数组的最大值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func splitArray(nums []int, k int) int {
	/*
	   输入：数组元素都为正数,  分组数k
	   输出：分割后每个分组求和，找到最大值,
	         这个最大值在所有情况的分组中最小


	   朴素的思路是，先分组，然后求每个分组的和，然后比较分组和的大小


	   递归思路：

	       求子集，每个子集求和
	*/

	ans := math.MaxInt
	sum := 0
	chosen := []int{}

	for i := 0; i < len(nums); i++ {
		dfs(i, nums, k, &sum, chosen)

		ans = min(ans, sum)

	}
	return ans

}

func dfs(i int, nums []int, k int, ans *int, chosen []int) {

	if i == len(nums) {
		if len(chosen) == k {
			s := sum(chosen)
			*ans = max(*ans, s)
		}
		return
	}

	dfs(i+1, nums, k, ans, chosen)
	chosen = append(chosen, nums[i])
	dfs(i+1, nums, k, ans, chosen)
	chosen = chosen[:len(chosen)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(chosen []int) int {
	ans := 0
	for _, c := range chosen {

		ans += c

	}
	return ans

}

/*
   总结：递归求子集思路不对, 不是子集
*/

// @lc code=end

/*
// @lcpr case=start
// [7,2,5,10,8]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,4,4]\n3\n
// @lcpr case=end

*/

