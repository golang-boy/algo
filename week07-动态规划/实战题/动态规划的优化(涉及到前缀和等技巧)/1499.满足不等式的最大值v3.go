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

		   题目中如果没有 if xj-xi <= k {
		   该条件时，可优化为

		for (int i = 2; i <= n; i++) {
			temp = max(temp, y[i - 1] - x[i - 1]);
			ans = max(ans, y[i] + x[i] + temp);
		}
		等价于

		for (int i = 2; i <= n; i++)
			for (int j = 1; j < i; j++)
				ans = max(ans, y[i] + y[j] + x[i] - x[j]);

	*/
	ans := math.MinInt
	q := []int{}

	for i := 0; i < len(points); i++ {

		// 内层循环中，j要满足这俩个条件
		//
		// 	上界: x[j] >= x[i] - k
		// 	下界: j <= i - 1
		//
		//  满足上述条件的基础上，求表达式的最大值

		//   使用单调队列解此题
		for len(q) != 0 {
			j := q[0]
			// x[j] < x[i]-k时，都出队
			if points[j][0] >= points[i][0]-k {
				break
			}
			q = q[1:]
		}

		// 此时，队列为满足上界的j, 求最大值
		if len(q) != 0 {
			ans = max(ans, points[i][1]+points[i][0]+points[q[0]][1]-points[q[0]][0])
		}

		for len(q) != 0 {
			// y[j] - x[j] 维护递减队列
			if points[q[len(q)-1]][1]-points[q[len(q)-1]][0] > points[i][1]-points[i][0] {
				break
			}
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return ans
}

/*
总结：

	单调队列还是需要继续复习，该题目的目的，主要是考虑怎么处理和优化多层循环，通过队列来处理内层的决策变量

*/

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[2,0],[5,10],[6,-10]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[0,0],[3,0],[9,2]]\n3\n
// @lcpr case=end

*/

