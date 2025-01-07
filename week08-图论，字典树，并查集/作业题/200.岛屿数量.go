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

	m := len(grid)
	n := len(grid[0])

	fa := make([]int, m*n)

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				id := nums(i, j, n)
				fa[id] = id
				// 有一个1则加加
				ans++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, grid, fa, &ans)
		}
	}
	return ans

}

func dfs(x, y int, grid [][]byte, fa []int, ans *int) {

	if grid[x][y] == '0' {
		return
	}

	m := len(grid)
	n := len(grid[0])

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		if nx < 0 || ny < 0 || nx >= m || ny >= n {
			continue
		}

		if grid[nx][ny] == '0' {
			continue
		}

		// 等于1时,

		rx := find(fa, nums(nx, ny, n))
		ry := find(fa, nums(x, y, n))

		if rx != ry {
			fa[rx] = ry
			// 每合并一个则--
			(*ans)--

		}
	}
}

func nums(x, y, n int) int {
	return x*n + y
}

func find(fa []int, x int) int {
	if fa[x] == x {
		return x
	}

	fa[x] = find(fa, fa[x])
	return fa[x]
}

/*
总结：
	初始时，每个1为一个岛屿，合并后，岛屿减一, 最后通过并查集合并后，剩余的则为岛屿数量

*/

// @lc code=end

/*
// @lcpr case=start
// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]\n
// @lcpr case=end

*/

