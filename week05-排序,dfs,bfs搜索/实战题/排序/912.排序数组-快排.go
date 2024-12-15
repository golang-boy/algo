/*
 * @lc app=leetcode.cn id=912 lang=golang
 * @lcpr version=20004
 *
 * [912] 排序数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func sortArray(nums []int) []int {

	/*
		排序算法分为二大类：
		   1. 基于比较的排序，时间复杂度下界为O(nlogn)
		      * 交换排序
		          冒泡，快速排序: 从数组中取中值，小的放左边，大的放右边
		      * 插入排序
		          简单插入排序，希尔排序
				  每次从前到后，考虑每个未排序的数据，在已排序的序列中找位置插入
		      * 选择排序
		          简单选择排序，堆排序
		          每次从未排序的数据找最小值，放到以排序列表末尾
		      * 归并排序
		          二路归并，多路归并
				  把数组前一半，后一半分别排序，最终合并

		   2. 基于非比较的排序，时间复杂度受元素范围和分布等多种因素影响
	*/

	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pIndex := partition(nums, left, right)
		quickSort(nums, left, pIndex)
		quickSort(nums, pIndex+1, right)
	}
}

func partition(nums []int, left, right int) int {

	p := nums[left]

	for left <= right {
		for nums[left] < p {
			left++
		}

		for nums[right] > p {
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
// [5,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,1,2,0,0]\n
// @lcpr case=end

*/

