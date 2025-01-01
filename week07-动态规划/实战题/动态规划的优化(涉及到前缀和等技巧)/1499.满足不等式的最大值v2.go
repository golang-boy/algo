/*
 * @lc app=leetcode.cn id=1499 lang=golang
 * @lcpr version=20004
 *
 * [1499] 满足不等式的最大值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findMaxValueOfEquation(points [][]int, k int) int {

	/*
		输入points和整数k,points按x已升序

		输出 yi+yj + |xi-xj| 最大值, |xi - xj| <= k

		针对v1进行优化
		   1. 去掉绝对值
		   2. 分离内外层循环

	*/
	ans := math.MinInt

	deque := []int{}

	for i := 0; i < len(points); i++ {

	}

	for j := 1; j < len(points); j++ {

		xj := points[j][0]
		yj := points[j][1]

		for i := 0; i < j; i++ {

			xi := points[i][0]
			yi := points[i][1]

			if xj-xi <= k {
				ans = max(ans, yj+xj+yi-xi)
			}
		}
	}
	return ans
}

func Abs(ans int) int {
	if ans > 0 {
		return ans
	}
	return -ans
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[2,0],[5,10],[6,-10]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[0,0],[3,0],[9,2]]\n3\n
// @lcpr case=end

*/

