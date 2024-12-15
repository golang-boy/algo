/*
 * @lc app=leetcode.cn id=1122 lang=golang
 * @lcpr version=20004
 *
 * [1122] 数组的相对排序
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {

	/*
	   1. 在arr2中，按arr2排序
	   2. 不在其中的，按升序排在arr1末尾

	   使用计数排序搞定
	*/

	// arr1中元素计数

	count := make([]int, 1001)
	ans := []int{}

	for _, num := range arr1 {
		count[num]++
	}

	for _, num := range arr2 {
		for count[num] > 0 {
			ans = append(ans, num)
			count[num]--
		}
	}

	for val := 0; val < 1001; val++ {
		for count[val] > 0 {
			ans = append(ans, val)
			count[val]--
		}
	}

	return ans

}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,3,2,4,6,7,9,2,19]\n[2,1,4,3,9,6]\n
// @lcpr case=end

// @lcpr case=start
// [28,6,22,8,44,17]\n[22,28,8,6]\n
// @lcpr case=end

*/

