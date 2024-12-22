/*
 * @lc app=leetcode.cn id=329 lang=golang
 * @lcpr version=20004
 *
 * [329] 矩阵中的最长递增路径
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestIncreasingPath(matrix [][]int) int {
	/*
	  输入一个矩阵m*n, 找最长递增路径

	  1. 需要在矩阵中进行搜索
	  2. 根据单元格中对应的值，找最长递增的路径
	      matrix[i][j] > pre    way += 1
	  3. 直到遍历到没有,需要去重记录
	      1. 周围的值小于等于当前值时停止
	      2. 另一种是到达边界时停止

	  一种是dfs，使用方向数组，每次搜到结尾时，与ans比较大小，保留大的
	      ans = max(ans, res)

	  一种是bfs，找层数最大的路径

	  怎么进行程序化？
	*/

	m := len(matrix)
	n := len(matrix[0])

	visited := make([][]bool, 0)
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			bfs(i, j, matrix, visited, &ans)
		}
	}
	return ans
}

func bfs(x, y int, matrix [][]int, visited [][]bool, ans *int) {

	dx := []int{-1, 0, 0, -1}
	dy := []int{0, -1, 1, 0}

	m := len(matrix)
	n := len(matrix[0])

	q := [][]int{
		[]int{x, y}, // 入队第一个位置
	}

	visited[x][y] = true
	depth := 0

	for len(q) != 0 {

		// 出队
		x = q[0][0]
		y = q[0][1]
		q = q[1:]
		for i := 0; i < 4; i++ {
			// 下一个点
			nx := x + dx[i]
			ny := y + dy[i]

			// 处理边界
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}

			// 去重
			if visited[nx][ny] {
				continue
			}

			// 这种情况递增
			if matrix[nx][ny] > matrix[x][y] {
				q = append(q, []int{nx, ny})
				visited[nx][ny] = true
				depth++
			}
		}
	}

	*ans = max(*ans, depth)
	return
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

/*
总结：
    情况没有考虑完全，使用bfs和dfs思路没问题，细节上考虑不全

	bfs和dfs要求每个点只能访问一次,按题目要求，算最长路径时，到当前点的路径也是要计算路径的

	因此该题需要在矩阵中，构建好有向无环图，针对该图进行拓扑排序，找最长的路径

	如何构造呢?


*/

// @lc code=end

/*
// @lcpr case=start
// [[9,9,4],[6,6,8],[2,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,4,5],[3,2,6],[2,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1]]\n
// @lcpr case=end

*/

