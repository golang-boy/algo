/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=20004
 *
 * [15] 三数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func threeSum(nums []int) [][]int {
	/*
		输入nums整数数组(有正有负)
		索引不同的三个数，和为0
		输出所有不重复的情况

		思路：
			遍历每个i，从i后面选俩数判断


	*/

	n := len(nums)
	ans := [][]int{}

	sort.Ints(nums)

	for i := 0; i < n; i++ {

		// 产生时，进行去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//  思路重点： 遍历每个i，从i后面选俩数判断
		sum := twoSum(nums, i, 0-nums[i])
		for _, s := range sum {
			ans = append(ans, append([]int{nums[i]}, s...))
		}
	}

	return ans
}

func twoSum(nums []int, ei int, k int) [][]int {

	ans := [][]int{}

	n := len(nums)
	j := n - 1

	for i := ei + 1; i < n; i++ {

		// 产生时，进行去重
		// 必须从ei+1开始，
		if i > ei+1 && nums[i] == nums[i-1] {
			continue
		}

		for ; i < j && nums[i]+nums[j] > k; j-- {
		}

		if i < j && nums[i]+nums[j] == k {
			ans = append(ans, []int{nums[i], nums[j]})
		}
	}

	return ans

}

/*
总结：

	三数之和，在两数之和的基础上增加一数。

	1. 遍历每个数, 从该数之后的数字中选2个进行判断
	2. 另一个重点是去重，要么结果里去重，要么产生时去重
	3. 两数之和的模型为：
		先对数组排序，双指针左右移动，大于k时，右侧移动，否则左侧移动，时间复杂度为O(n)


*/

// @lc code=end

/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

*/

