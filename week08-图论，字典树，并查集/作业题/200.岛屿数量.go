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

	var outside = n * m
	fa := make([]int, m*n+1)

	fa[nums(m-1, n-1, n)+1] = nums(m-1, n-1, n) + 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				id := nums(i, j, n)
				fa[id] = id
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, grid, fa, outside)
		}
	}

	ans := 0

	for i := 0; i < len(fa); i++ {
		if find(fa, i) != outside {
			ans++
			fmt.Println(fa[i])
		}
	}

	return ans

}

func dfs(x, y int, grid [][]byte, fa []int, outside int) {

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
			join(fa, nums(nx, ny, n), outside)
			continue
		}
		join(fa, nums(nx, ny, n), nums(x, y, n))
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

func join(fa []int, x, y int) {
	x = find(fa, x)
	y = find(fa, y)

	if x != y {
		fa[x] = y
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

