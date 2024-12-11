/*
 * @lc app=leetcode.cn id=74 lang=golang
 * @lcpr version=20004
 *
 * [74] 搜索二维矩阵
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {

	/*
		   输入：矩阵与目标值
		   输出：是否包含目标值

		   思路：看起来是二分搜索，不同点在于在矩阵中搜索
		     1. 放到一个数组中，直接二分？
		     2. 在每一行进行二分查找？
		     3. 是有规律的, 先确定行，再确定列?
			     每行加够n时，到下一行

	*/

	m := len(matrix)
	n := len(matrix[0])

	left := 0
	right = m*n - 1
	// 这个边界条件试了很久，整个矩阵想成一个数组，数组最后一个元素为len(nums)-1
	// 对于矩阵来说，就是m*n-1, 就是最后一个元素

	for left <= right {
		mid := (left + right) / 2

		x, y := toXy(mid, n)
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func toXy(v, n int) (int, int) {
	return v / n, v % n
}
func toV(x, y, n int) int {
	return x*n + y
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n13\n
// @lcpr case=end

*/

