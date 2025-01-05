/*
 * @lc app=leetcode.cn id=130 lang=golang
 * @lcpr version=20004
 *
 * [130] 被围绕的区域
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func solve(board [][]byte) {

	/*
	   并查集实现
	      将矩阵边上的所有O都连入同一个并查集
	*/

	m := len(board)
	n := len(board[0])

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	outside := m * n

	// 初始化并查集
	fa := make([]int, m*n+1) // m*n个点+外部的一个点
	for i := 0; i < m*n+1; i++ {
		fa[i] = i
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}

			// 等于O的情况时，查看四个方向
			for k := 0; k < 4; k++ {
				ni := i + dx[k]
				nj := j + dy[k]

				if ni < 0 || nj < 0 || ni >= m || nj >= n {
					// 等于O，且该方向到达边界, 则将该点加入并查集
					// 怎么加呢？
					unionSet(fa, nums(n, i, j), outside)
				} else {
					// 下一个还是O, 且没有到边界，这时需要将当前的与下一个进行连接
					if board[ni][nj] == 'O' {
						unionSet(fa, nums(n, i, j), nums(n, ni, nj))
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 遍历矩阵，元素为O，并且并查集查询结果与根不等时，改为X
			if board[i][j] == 'O' && find(fa, nums(n, i, j)) != find(fa, outside) {
				board[i][j] = 'X'
			}
		}
	}
}

func find(fa []int, i int) int {
	if fa[i] == i {
		return i
	}
	fa[i] = find(fa, fa[i])
	return fa[i]
}

func unionSet(fa []int, i, j int) {

	i = find(fa, i)
	j = find(fa, j)

	if i != j {
		fa[i] = j
	}
}

func nums(n, i, j int) int {
	return i*n + j
}

// @lc code=end

/*
// @lcpr case=start
// [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["X"]]\n
// @lcpr case=end

*/

