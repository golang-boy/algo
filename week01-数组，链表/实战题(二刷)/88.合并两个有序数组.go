/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {

	/*
	 */

	// 方法一、加进去重新排序
	// for i := range nums2 {
	// 	nums1[m+i] = nums2[i]
	// }

	// sort.Ints(nums1)

	i := m - 1
	j := n - 1

	k := n + m - 1

	for ; k >= 0; k-- {

		if i < 0 && j >= 0 {
			break
		}

		if j < 0 {
			break
		}

		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}

	if i < 0 && j >= 0 {
		for ; j >= 0; j-- {
			nums1[k] = nums2[j]
			k--
		}
	}

}

// @lc code=end

