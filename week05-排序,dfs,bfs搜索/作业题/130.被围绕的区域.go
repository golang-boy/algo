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
		输入board,二维byte矩阵, m,n 在[1,200],值要么是'X',要么是'O'

		无输出，但
		 1. 需要在输入的board上进行操作，
		 2. 将所有'O'替换为'X'
		 3. 捕获被围绕的区域

		围绕：区域由O构成，X围绕这些O
		且这些O不在board边缘,没有被围绕的不用替换


		靠近边缘的O不用被替换,题目是否可以转换为找到边缘的O,其他位置的都被替换?

		如何确定是否是边缘?

		m行n列，0, m-1行都是边界，0,n-1列都是边界, i,j不等于这些点时, 进行替换

		不行，区域内的才能被替换,因此
		   1. 需要先知道区域
		   2. 然后看是否被捕获
		   3. 捕获后才能被替换


		  首先找值为O的所有区域,
		  然后判断O是否满足条件(如果区域中有一个O位于边缘，则这个区域不进行替换)

		  深度优先时,每找到一个点就判断是否处于边缘，是否等于O

		  不能在找的过程中进行判断, 需要先将所有连通区域找到，然后遍历该图进行判断

		  找节点的过程可以使用bfs，也可以dfs


	*/

	m := len(board)
	n := len(board[0])

	zones := [][][]int{}
	zone := [][]int{}
	visited := make([]bool, m*n)
	// [][][]int{
	//   [][]int{
	// 	   []int{1, 1},
	// 	   []int{1, 2},
	// 	   []int{2, 2},
	//   },
	// }

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			dfs(i, j, board, &zones, zone, visited)

		}
	}

	for _, zone := range zones {
		flag := false
		for _, p := range zone {
			i := p[0]
			j := p[1]
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				flag = true
				break // 当前区域结束
			}
		}

		if !flag {
			for _, p := range zone {
				i := p[0]
				j := p[1]
				board[i][j] = 'X'
			}

			fmt.Println(zone)
		}
		fmt.Println("==============")
	}

	return
}

func dfs(i, j int, board [][]byte, zones *[][][]int, zone [][]int, visited []bool) {

	// 什么时候结束?
	//   没有相邻的O时结束返回, 上下左右扩展后都没有,或者到达边界

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	m := len(board)
	n := len(board[0])

	if i < 0 || j < 0 || i > m-1 || j > n-1 || board[i][j] == 'X' {
		temp := make([][]int, len(zone))
		for i := range zone {
			temp[i] = make([]int, len(zone[0]))
			copy(temp[i], zone[i])
		}
		*zones = append(*zones, zone)
		return
	}

	if board[i][j] == 'O' {
		zone = append(zone, []int{i, j})
		visited[i*n+j] = true
	}
	// 扩展
	for k := 0; k < 4; k++ {
		nx := i + dx[k]
		ny := j + dy[k]

		// 不在zone中
		if nx >= 0 && ny >= 0 && nx < m && ny < n && !visited[nx*n+ny] {
			dfs(nx, ny, board, zones, zone, visited)
		}
	}
}

/*
总结：
    应该从边开始遍历，而不是每个节点都遍历一次
	结束条件没有确定好就开始套模板

*/

// @lc code=end

/*
// @lcpr case=start
// [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["X","O","X"],["X","O","X"],["X","O","X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["X"]]\n
// @lcpr case=end

*/

