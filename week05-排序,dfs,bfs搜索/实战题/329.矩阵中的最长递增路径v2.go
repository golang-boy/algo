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

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	g := make([][]int, m*n) // 出边数组
	deg := make([]int, m*n)
	dist := make([]int, m*n) // 距离

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			for k := 0; k < 4; k++ {
				// 下一个点
				nx := x + dx[k]
				ny := y + dy[k]

				// 处理边界
				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				}

				if matrix[nx][ny] > matrix[x][y] {
					// 从x,y节点连接一条边到nx，ny
					addEdge(g, deg, node(x, y, n), node(nx, ny, n))
				}
			}
		}
	}

	q := []int{}

	for i := 0; i < m*n; i++ {
		if deg[i] == 0 {
			q = append(q, i)
			dist[i] = 1 // 起点的路径长度为1
		}
	}

	for len(q) != 0 {
		// 出队
		x := q[0]
		q = q[1:]

		for _, y := range g[x] {
			deg[y]--
			dist[y] = max(dist[y], dist[x]+1)
			if deg[y] == 0 {
				q = append(q, y)
			}
		}
	}

	for i := 0; i < m*n; i++ {
		ans = max(ans, dist[i])
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

func node(x, y, n int) int {
	// n为列的长度
	// 总结点数量为 m*n
	return x*n + y
}

func addEdge(g [][]int, deg []int, from, to int) {
	deg[to]++
	// 从from可以到达多个to点

	// fmt.Println(from, to)

	if g[from] == nil {
		g[from] = make([]int, 0)
	}
	g[from] = append(g[from], to)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
总结：


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

