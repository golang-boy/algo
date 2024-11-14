/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {

	// 目标: 把nums2中的元素放入nums1中

	/*
		已知条件：
		    1. nums1有足够的空间容纳nums2中的元素,
		    2. m为nums1的有效元素个数, n为nums2的元素个数
			3. nums1和nums2都是有序的
	*/

	/* 分析：
		1. 数组合并，必定是要遍历数组的,因为有序，所以取nums1和nums2的元素进行比较，然后放入新的数组中
		2. 根据过滤条件，决定要哪个数
		3. 边界条件怎么确定？

	         i
			[1,2,3,0,0,0]
		    [2,5,6]
			 j

			[1]

	           i
			[1,2,3,0,0,0]
		    [2,5,6]
			 j
			[1,2]

	             i
			[1,2,3,0,0,0]
		    [2,5,6]
			 j
			[1,2,2]

	             i
			[1,2,3,0,0,0]
		    [2,5,6]
			   j
			[1,2,2,3]

	               i
			[1,2,3,0,0,0]    i > m 时，nums1已经遍历完了， nums2的元素直接放入nums1中
		    [2,5,6]
			   j
			[1,2,2,3]


	*/

	/*
		问：单层循环和双层循环的区别？
		答：双层循环，时间复杂度O(m*n), 单层循环时间复杂度O(m+n)

		问：为什么这里要用单层循环？
		答：因为nums1和nums2都是有序的，所以可以比较大小，然后放入新的数组中，不需要双层循环

		问：不开辟新的数组，直接在nums1上操作可以吗？
		答：可以，但是需要从后往前遍历，因为nums1后面的空间是足够的，可以存放nums2的元素
	*/

	// 正着来
	// i := 0
	// j := 0
	// ans := []int{}
	// for {

	// 	if i >= m || j >= n {
	// 		break
	// 	}

	// 	if nums1[i] <= nums2[j] {
	// 		ans = append(ans, nums1[i])
	// 		i++
	// 	}

	// 	if nums1[i] > nums2[j] {
	// 		ans = append(ans, nums2[j])
	// 		j++
	// 	}
	// }

	// for i < m {
	// 	ans = append(ans, nums1[i])
	// 	i++
	// }

	// for j < n {
	// 	ans = append(ans, nums2[j])
	// 	j++
	// }

	// copy(nums1, ans)

	// 倒着来

	i := m - 1
	j := n - 1

	gte := func(i, j int) bool {
		if i < 0 { // nums1已经遍历完了, nums2中的元素直接放入nums1中
			return false
		}

		if j < 0 { // nums2已经遍历完了, nums1中的元素直接放入nums1中
			return true
		}

		return nums1[i] >= nums2[j]
	}

	for k := m + n - 1; k >= 0; k-- {

		if gte(i, j) {
			nums1[k] = nums1[i] // 因为已经排序，从末尾开始比较，所以直接取大的，放到最后面
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}

// @lc code=end

