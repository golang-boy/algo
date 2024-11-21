
/*
 * @lc app=leetcode.cn id=46 lang=golang
 * @lcpr version=20003
 *
 * [46] 全排列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func permute(nums []int) [][]int {
	/*
	   关键信息：
	     1. 无重复数字
	     2. 子集顺序相关
	     3. 结果内部顺序无关

	   思路：
	     1. 划分问题子集
	       也是选不选的问题, 但是与子集和组合不一样
	     2. 边界条件是什么？

	*/

	ans := [][]int{}
	chosen := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	// 从0开始
	recure(0, nums, &ans, used, chosen)

	return ans
}

func recure(i int, nums []int, ans *[][]int, used []bool, chosen []int) {

	if i == len(nums) { //递归到边界
		temp := make([]int, len(chosen))
		copy(temp, chosen)
		*ans = append(*ans, temp)
	}

	for j := 0; j < len(nums); j++ {
		if !used[j] {

			// 选择j
			chosen = append(chosen, nums[j])
			used[j] = true

			// 进入下一层
			recure(i+1, nums, ans, used, chosen)

			// 返回上一层，撤销选择,回到上一层的状态
			chosen = (chosen)[:len(chosen)-1]
			used[j] = false
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

