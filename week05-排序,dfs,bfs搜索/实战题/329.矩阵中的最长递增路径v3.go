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

		该题需要在矩阵中，构建好有向无环图，针对该图进行拓扑排序，找最长的路径

		如何构造呢?
		  比当前节点大的点构造边，进而构造图
	*/

	m := len(matrix)
	n := len(matrix[0])

	ans := 0

	dist := make([][]int, m) // 距离
	for i := range dist {
		dist[i] = make([]int, n)
	}

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {

			// 从每个节点开始进行搜索, 计算
			ans = max(dfs(x, y, matrix, dist), ans)
		}
	}

	return ans

	/*
		 构造图
		    格子是点，相邻递增时连边

			点怎么表示?
			 出边数组是一个二维数组，索引为节点标号，值为可以达到的节点的标号
			 二维矩阵中，需要将坐标点x,y表示为出边数组中的索引值
			 x*n+y

	*/
}

// dfs返回距离最底层节点的距离
func dfs(x, y int, matrix [][]int, dist [][]int) int {

	m := len(matrix)
	n := len(matrix[0])
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	// 记忆搜索，如果已经有值，说明已经搜索计算过，直接返回即可
	if dist[x][y] != 0 {
		return dist[x][y]
	}

	// 等于0时，直接将当前节点赋值1,
	dist[x][y] = 1

	// 超其他方向扩展
	for k := 0; k < 4; k++ {
		// 下一个点
		nx := x + dx[k]
		ny := y + dy[k]
		// 处理边界
		if nx < 0 || nx >= m || ny < 0 || ny >= n {
			continue
		}

		if matrix[nx][ny] > matrix[x][y] {
			// 每次从下一层回来后，加一
			d := dfs(nx, ny, matrix, dist) + 1

			// 当前值与返回后的取最大的, 当前为1，下层返回后比1大，如果已经遍历过，仍旧是取最大的
			dist[x][y] = max(dist[x][y], d)
		}
	}

	// 扩展完毕或者没有扩展，直接返回
	return dist[x][y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
总结：

    深度优先搜索看起来更直观一点


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

