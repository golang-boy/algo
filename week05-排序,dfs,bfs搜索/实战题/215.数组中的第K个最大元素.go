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

	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, left, right, index int) int {
	if left >= right {
		return nums[left]
	}

	pIndex := partition(nums, left, right)
	if index <= pIndex {
		return quickSort(nums, left, pIndex, index)
	} else {
		return quickSort(nums, pIndex+1, right, index)
	}
}

func partition(nums []int, left, right int) int {

	pivot := left
	for left <= right {

		for nums[left] < nums[pivot] {
			left++
		}
		for nums[right] > nums[pivot] {
			right--
		}

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

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/

