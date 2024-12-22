/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=20004
 *
 * [200] 岛屿数量
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numIslands(grid [][]byte) int {

	/*

	   输入一个二维矩阵,矩阵大小范围 1-300, 每个小格子值0或1
	   输出数量，找到相邻的1构成的岛屿数量

	   广度优先搜索，找一个1，依次找相邻的格子，找到后累加结果计数，注意判重

	   如何进行程序化?
	      1. 广度优先搜索，先要构造一个图？邻接表？
	              邻接表在该题中不易构建
	      2. 如何构造在图中的遍历?数据结构怎么表示?
	              方向数组？方向加一然后搜索？

	     网格类，地图类搜索题目，搜索找路径，或联通块

	*/

	t := MakeNumIslands(grid)
	return t.Calc()
}

type NumIslands struct {
	dx      []int
	dy      []int
	visited [][]bool
	ans     int
	m, n    int

	grid [][]byte
}

func MakeNumIslands(grid [][]byte) *NumIslands {

	m := len(grid)
	n := len(grid[0])

	//           左,不变,不变,右
	dx := []int{-1, 0, 0, 1}
	//           不变,下,上,不变
	dy := []int{0, -1, 1, 0}

	visited := make([][]bool, 0)
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}

	ans := 0

	return &NumIslands{
		m:       m,
		n:       n,
		dx:      dx,
		dy:      dy,
		visited: visited,
		ans:     ans,

		grid: grid,
	}
}

func (t *NumIslands) Calc() int {
	// 找每个为1的点
	for i := 0; i < t.m; i++ {
		for j := 0; j < t.n; j++ {
			if t.grid[i][j] == '1' && !t.visited[i][j] {
				t.ans++
				t.bfs(i, j)

			}
		}
	}
	return t.ans
}

func (t *NumIslands) bfs(sx, sy int) {
	q := [][]int{}

	// 先入队
	q = append(q, []int{sx, sy})
	t.visited[sx][sy] = true

	for len(q) != 0 {
		// 取队头
		x := q[0][0]
		y := q[0][1]
		q = q[1:] // 是要取队头，不是取队尾

		// 扩展队头
		// 从四个方向进行扩展
		for i := 0; i < 4; i++ {
			nx := x + t.dx[i]
			ny := y + t.dy[i]

			// 处理边界
			if nx < 0 || nx >= t.m || ny < 0 || ny >= t.n {
				continue
			}

			// 处理0的情况
			if t.grid[nx][ny] != '1' {
				continue
			}

			// 去重
			if t.visited[nx][ny] {
				continue
			}
			q = append(q, []int{nx, ny})
			t.visited[nx][ny] = true
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]\n
// @lcpr case=end

*/

