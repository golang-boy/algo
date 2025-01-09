
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
			左右确定俩位置(i,j)， 在此区间使用k进行搜索
	*/

	n := len(nums)
	ans := [][]int{}

	return ans
}

/*
总结：
	暴力枚举中关于本题去重的思路
		1. 排序后，转为字符串进行
		2. 使用结构体，map中的key需要时可以比较的
			前提是结构体的所有字段都是可比较的
*/

type Triplet struct {
	a, b, c int
}

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

