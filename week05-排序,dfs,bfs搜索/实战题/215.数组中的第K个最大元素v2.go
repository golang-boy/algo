/*
 * @Author: 刘慧东
 * @Date: 2024-12-16 11:35:37
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-12-17 11:20:54
 */
/*
 * @lc app=leetcode.cn id=215 lang=golang
 * @lcpr version=20004
 *
 * [215] 数组中的第K个最大元素
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findKthLargest(nums []int, k int) int {
	/*
	   输入： 整数序列，整数k
	   输出：第k个最大元素
	        元素中有可能重复

	*/

	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, left, right, index int) int {
	if left >= right {
		return nums[left]
	}

	pindex := partition(nums, left, right)
	if pindex >= index {
		return quickSort(nums, left, pindex, index)
	}
	return quickSort(nums, pindex+1, right, index)
}

func partition(nums []int, left, right int) int {

	pivot := left + rand.IntN(right-left+1)
	p := nums[pivot] // 这里需要使用临时变量存储，否则后面移动时，中轴值会变
	for left <= right {

		// 左侧找是否能找到
		for nums[left] < p {
			left++
		}

		// 右侧找是否能找到
		for nums[right] > p {
			right--
		}

		// 找到左或者右等于pivot的值

		// 判断此时left与right的情况，如果等于
		if left == right {
			break
		}

		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	return right
}

/*
	总结：
	    快排，从小到大只排一半的元素，在这一半中找到第k大元素

*/

// @lc code=end

/*
// @lcpr case=start
// [5,2,4,1,3,6,0]\n4
// @lcpr case=end

// @lcpr case=start
// [3,3,3,3,4,3,3,3,3]\n1
// @lcpr case=end

*/

