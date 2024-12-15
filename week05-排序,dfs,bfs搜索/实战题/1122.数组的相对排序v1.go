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
	 */

	q := map[int]int{}

	for i, v := range arr2 {
		q[v] = i
	}

	slices.SortFunc(arr1, func(a, b int) int {

		aPos, ok := q[a]
		if !ok {
			aPos = len(arr2)
		}
		bPos, ok := q[b]
		if !ok {
			bPos = len(arr2)
		}

		// 如果都在arr2中，按已有顺序排
		if aPos < bPos {
			return -1
		}

		// 如果都不在arr2中，按大小排(顺序不可能相等)
		if aPos == bPos {
			if a < b {
				return -1
			}
			if a == b {
				return 0
			}
		}

		// todo: 得看看 sort的源码是怎么回事?学会slices包的使用

		return 1
	})

	return arr1

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

