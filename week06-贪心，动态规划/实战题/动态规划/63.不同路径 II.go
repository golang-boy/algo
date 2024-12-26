/*
 * @lc app=leetcode.cn id=63 lang=golang
 * @lcpr version=20004
 *
 * [63] 不同路径 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	/*
	   输入：坐标矩阵地图, 有障碍物矩阵中表示1
	   输出：到达右下角的不同路径数量

	   path = path(a) + path(b)
	*/
	visited := make([][]int, len(obstacleGrid))
	for i := range visited {
		visited[i] = make([]int, len(obstacleGrid[0]))
	}

	return countPath(obstacleGrid, 0, 0, visited)
}

func countPath(obstacleGrid [][]int, row, col int, visited [][]int) int {

	if visited[row][col] != 0 {
		return visited[row][col]
	}

	if obstacleGrid[row][col] == 1 {
		return 0
	}

	if len(obstacleGrid)-1 == row && len(obstacleGrid[0])-1 == col {
		return 1
	}

	if len(obstacleGrid)-1 == row {
		ans := countPath(obstacleGrid, row, col+1, visited)
		visited[row][col] = ans
		return ans
	}
	if len(obstacleGrid[0])-1 == col {
		ans := countPath(obstacleGrid, row+1, col, visited)
		visited[row][col] = ans
		return ans
	}
	ans := countPath(obstacleGrid, row+1, col, visited) + countPath(obstacleGrid, row, col+1, visited)
	visited[row][col] = ans
	return ans
}

/*

总结：
    定义好状态，推导出状态方程，确定状态空间位树状结构，进行dfs，同时使用visited进行记忆化搜索



*/

// @lc code=end

/*
// @lcpr case=start
// [[0,0,0],[0,1,0],[0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[0,0]]\n
// @lcpr case=end

*/

