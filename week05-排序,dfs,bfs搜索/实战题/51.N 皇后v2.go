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

	ans := [][]int{}
	ansStr := [][]string{}

	rows := []int{}

	// 同一列
	used := make(map[int]bool, n)
	// 正对角线
	usedPlus := make(map[int]bool, n)
	// 斜对角线
	usedMinus := make(map[int]bool, n)

	dfs(0, n, &ans, rows, used, usedPlus, usedMinus)

	fmt.Println(ans)

	for _, cols := range ans {

		rowStr := []string{}
		for _, col := range cols {
			row := []byte{}
			for i := 0; i < n; i++ {
				if i == col {
					row = append(row, 'Q')
				} else {
					row = append(row, '.')
				}
			}
			rowStr = append(rowStr, string(row))
		}
		ansStr = append(ansStr, rowStr)
	}
	return ansStr
}

func dfs(row int, n int, ans *[][]int, rows []int, used, usedPlus, usedMinus map[int]bool) {
	if row == n {
		temp := make([]int, n)
		copy(temp, rows)
		*ans = append(*ans, temp)
		return
	}

	for col := 0; col < n; col++ {

		// 所有对角线有两种，col+row和col-row该值是唯一的，不论在哪一行那一列
		if !used[col] && !usedPlus[col+row] && !usedMinus[col-row] {

			used[col] = true
			usedMinus[col-row] = true
			usedPlus[col+row] = true

			// 尝试当前列是不是合适，到叶子时有结果则追加

			// 每行一个列标量
			rows = append(rows, col)
			dfs(row+1, n, ans, rows, used, usedPlus, usedMinus)
			//移除当前列，尝试下一列放置
			rows = rows[:len(rows)-1]

			used[col] = false
			usedMinus[col-row] = false
			usedPlus[col+row] = false
		}
	}
}

/*
总结：
    整体思路没问题，就思路转为程序时，出现问题
		1. 每行中皇后位置信息的表达表示，做题前没有明确下来
		     位置信息表达，最简单直接的方式，将第几列的信息放入数组中，每行对应一个索引，索引对应的值，表示该行第几列
		2. 知道是要转换，但是最终结果处理时，还是有些模糊
		3. 另一个难点是对角线的处理，这个不太容易确定其表达方式
		4. 另外一个点是如何判断以前假设的处理过还是没处理过

	本题基本的思想是，以递归的方式遍历每行，递归函数内，处理每列
	递归时，每次处理一行，递归结束后,将结果追加进ans中。递归中假设每列皇后的位置

*/

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 6\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

