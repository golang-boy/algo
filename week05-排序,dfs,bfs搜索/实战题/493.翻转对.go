/*
 * @lc app=leetcode.cn id=493 lang=golang
 * @lcpr version=20004
 *
 * [493] 翻转对
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reversePairs(nums []int) int {

	/*
	   给俩位置，满足不等式，统计关系数量，可以使用归并排序

	*/

	ans := 0
	mergeSort(nums, 0, len(nums)-1, &ans)

	return ans
}

func mergeSort(nums []int, l, r int, ans *int) {

	if l >= r {
		return
	}

	mid := (l + r) / 2
	mergeSort(nums, l, mid, ans)
	mergeSort(nums, mid+1, r, ans)

	// 1,2,3    1,3

	// calc([1,3,2,3,1]) = calc([1,3,2]) + calc([1,3]) + calc([左一个，右一个 ])

	calc(nums, l, mid, r, ans)

	merge(nums, l, mid, r)
}

func calc(nums []int, l, m, r int, ans *int) {

	j := m + 1
	for i := l; i <= m; i++ {

		//  左右两侧比较,拿左侧的每一个i，依次与右侧的j位置比较,满足条件的j++,
		for j <= r && nums[i] > 2*nums[j] {
			j++
		}

		// j初始为mid+1, 中间加了几次，答案记几次
		*ans += j - m - 1
	}
}

func merge(nums []int, l, m, r int) {
	temp := make([]int, r-l+1)
	i := l
	j := m + 1
	k := 0

	for ; i <= m && j <= r; k++ {

		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
	}

	for ; i <= m; i++ {
		temp[k] = nums[i]
		k++
	}
	for ; j <= r; j++ {
		temp[k] = nums[j]
		k++
	}

	for i := range temp {
		nums[l+i] = temp[i]
	}

}

// @lc code=end

/*
// @lcpr case=start
// [1,3,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,4,3,5,1]\n
// @lcpr case=end

*/

