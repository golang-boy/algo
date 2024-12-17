/*
 * @Author: 刘慧东
 * @Date: 2024-12-17 11:54:20
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-12-17 11:57:20
 */
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

	length := len(nums)

	if length < 2 {
		return nums
	}

	mid := length / 2

	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	return merge(left, right)
}

func merge(nums []int, left, right, mid int) {
	temp := make([]int, right-left+1)

	i, j, k := left, mid+1, 0

	for ; i <= mid && j <= right; k++ {

	}

}

func merge(left, right []int) []int {
	ans := []int{}

	if len(left) == 0 && len(right) == 0 {
		return ans
	} else if len(left) != 0 && len(right) == 0 {
		ans = append(ans, left...)
		return ans
	} else if len(right) != 0 && len(left) == 0 {
		ans = append(ans, right...)
		return ans
	}

	if left[0] >= right[0] {
		ans = append(ans, right[0])
		ans = append(ans, merge(left, right[1:])...)
	} else {
		ans = append(ans, left[0])
		ans = append(ans, merge(left[1:], right)...)
	}

	return ans
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

