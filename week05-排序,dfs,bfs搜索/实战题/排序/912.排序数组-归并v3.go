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
		          冒泡，快速
		      * 插入排序
		          简单插入排序，希尔排序
				  每次从前到后，考虑每个未排序的数据，在已排序的序列中找位置插入
		      * 选择排序
		          简单选择排序，堆排序
		          每次从未排序的数据找最小值，放到以排序列表末尾
		      * 归并排序
		          二路归并，多路归并

		   2. 基于非比较的排序，时间复杂度受元素范围和分布等多种因素影响
	*/

	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)

	merge(nums, l, mid, r)
}

func merge(nums []int, l, mid, r int) {

	temp := make([]int, r-l+1)

	// 合并两个有序数组, mid+1 到r为第二个数组，0到mid为第一个数组

	i := l
	j := mid + 1
	k := 0

	for ; i <= mid && j <= r; k++ {

		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
	}

	for ; i <= mid; k++ {
		temp[k] = nums[i]
		i++
	}
	for ; j <= r; k++ {
		temp[k] = nums[j]
		j++
	}

	for i := range temp {
		nums[l+i] = temp[i]
	}

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

