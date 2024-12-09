/*
 * @lc app=leetcode.cn id=154 lang=golang
 * @lcpr version=20004
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findMin(nums []int) int {

	/*

	  1. 旋转数组, 升序排列后旋转
	  2. 有重复值, 最小值也可能有多个重复的, 这种情况下得找到第一个

	  nums[mid] <= nums[right]
	*/

	left := 0
	right := len(nums) - 1

	for left < right {

		mid := (left + right) / 2

		if nums[mid]-nums[right] < 0 {
			right = mid
		} else if nums[mid]-nums[right] > 0 {
			left = mid + 1
		} else {
			right--
		}
	}

	fmt.Println(right, nums[right])
	fmt.Println(left, nums[left])
	return nums[left]

}

/*
总结：
     通过二分法进行解题，
	    求下边界时（第一个(>=target或者<=target的)） right = mid
		    满足条件时，最小的那个, 或者最前面的那个

		求上边界时 left = mid
		    满足条件时，最大的那个，或者最后面的那个
			这种情况，mid计算时需要补1

		需要注意是，条件一定要用if elseif  else，否则会有问题

	二分的前提是：
	   1. 目标函数单调
	   2. 有边界
	   3. 索引访问
*/

// @lc code=end

/*
// @lcpr case=start
// [1,3,5]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,0,1]\n
// @lcpr case=end

*/

