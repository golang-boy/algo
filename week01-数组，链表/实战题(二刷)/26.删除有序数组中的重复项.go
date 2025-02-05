/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {

	// 原地删除数组中重复元素
	//  输出删除后的元素的个数
	//   数组为递增数组，存在相等的情况

	// [0,0,1,1,1,2,2,3,3,4]
	//  i j
	//  i   j
	// [0,1,1,1,1,2,2,3,3,4]

	//    i j
	// [0,1,1,1,1,2,2,3,3,4]

	//    i       j
	// [0,1,2,1,1,2,2,3,3,4]

	//      i         j
	// [0,1,2,1,1,2,2,3,3,4]
	//          i         j
	// [0,1,2,3,4,2,2,3,3,4]

	j := 0
	i := 0
	for ; i < len(nums)-1; i++ {
		for ; j < len(nums) && nums[i] == nums[j]; j++ {
		}
		if j < len(nums) {
			nums[i+1] = nums[j]
		} else {
			break
		}
	}
	return i + 1
}

// @lc code=end

