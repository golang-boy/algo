/*
 * @lc app=leetcode.cn id=51 lang=golang
 * @lcpr version=20004
 *
 * [51] N 皇后
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func solveNQueens(n int) [][]string {

	/*
	   输入是整数n,表示棋盘大小nxn个格子,n个皇后放在这些格子，每行要放一个皇后

	   同一行，同一列，左斜线，右斜线

	   输出格式可以转换为01表示，1表示皇后所在的位置
	   输出一个矩阵，标出皇后所在位置

	   [
	      [0,1,0,0],
	      [0,0,0,1],
	      [1,0,0,0],
	      [1,0,1,0],
	   ]

	   皇后位置变化，需要记录位置，判断时，查看同一行或列是否已有皇后

	   位置使用i,j表示行列

	   尝试放在第一行第一列，然后尝试第二行，第三行等，直到最末尾,
	   看每行是否都能放下皇后,如果能放下，这记录这个该行的这个结果，直到末尾或者放不下

	   然后尝试放第一行第二列


	*/

	ans := [][][]int{}

	used := make([]bool, n)

	state := make([][]int{}, n)
	for i := 0; i < n; i++ {
		state[i] = make([]int, n)
	}

	// 行
	for i := 0; i < n; i++ {
		dfs(i, state, n, &ans, used)
	}

}

func dfs(i int, state [][]int, n int, ans *[][][]int, used []bool) {

	if i == n-1 {
		return
	}

	for j := 0; j < n; j++ {
		if state[i][j] == 0 {
		}

	}

}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

