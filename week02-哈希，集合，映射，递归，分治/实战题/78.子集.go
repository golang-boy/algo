
/*
 * @lc app=leetcode.cn id=78 lang=golang
 * @lcpr version=20003
 *
 * [78] 子集
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func subsets(nums []int) [][]int {
	/*

		  关键信息：
		    1. 数组，元素不同
			2. 返回 数组，数组里包含所有子集构成的数组
			3. 子集不能重复

		  解题思路：
		    1. 选不选的问题
			  从0开始，选下标0或不选下标0
			2. 递归到最底下
	*/

	res := [][]int{}

	getSubset(nums, 0, func(set []int) {
		res = append(res, set)
	})

	return res
}

func getSubset(nums []int, i int, fn func(set []int)) {
	if i == len(nums) {
		return
	}

	chose := append([]int{}, nums[i])

	getSubset(nums, i+1, fn)
}

/*
递归的关键:
  1. 定义需要递归的问题
  2. 确定递归边界
  3. 保护与还原现场

  递归相当于有个隐藏的栈，将扫描未扫描的不停的放入栈中，直到最底下，然后不停的出栈
     因此，需要先确定最底层的逻辑, 在考虑返回后怎么处理
*/

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

