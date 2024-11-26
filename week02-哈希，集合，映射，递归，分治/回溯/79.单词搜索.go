/*
 * @lc app=leetcode.cn id=79 lang=golang
 * @lcpr version=20003
 *
 * [79] 单词搜索
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func exist(board [][]byte, word string) bool {

	/*
		   分析：
		      1. 先横着找，找不到，回到前一个格子，往下找
			  2. 还能往左找

			  每个格子，有三个方向，左，下，右, 依次暴力求解

			     找第一个，然后从三个方向找第二个

			  i:=0  i< n ; i++    往右
			  怎么往左？

			  1. 找第一个元素，找到后，从当前行找下一个元素，
			     如果是，继续往下找，直到不是
				 不是时，跳出当前行，从下一行开始找第一个元素（列不变） ==》 递归, 传递列信息
				 直到最后

	*/

	chars := []byte(word)

	m := len(board)
	n := len(board[0])

	found := [][]int{}
	ans := false

	for x := 0; x < m; x++ { // 遍历每行, 第x行开始找
		for y := 0; y < n; y++ { // 每一列, 从下标y开始,往右找
			if board[x][y] == chars[0] { // 当前
				found = append(found, []int{x, y})
			}
		}
	}

	if len(found) == 0 {
		return false
	}

	for _, p := range found {
		q := map[complex64]bool{}
		ans = r(p[0], p[1], board, chars, -1, q) || ans
	}
	return ans
}

func r(x, y int, board [][]byte, chars []byte, dir int, q map[complex64]bool) bool {

	m := len(board)
	n := len(board[0])
	var down, left, up, right bool

	// 到末尾后，chars还有则返回false，

	if len(chars) == 0 {
		return true
	}

	if x == -1 || y == -1 {
		return false
	}

	if x == m {
		return len(chars) == 0
	}

	if y == n {
		return len(chars) == 0
	}

	p := complex(float32(x), float32(y))
	if q[p] {
		return false
	}

	if board[x][y] == chars[0] { // 当前
		chars = chars[1:]
		q[p] = true

		if dir != 1 { // 上
			down = r(x+1, y, board, chars, 3, q) // 递归往下
		}
		if dir != 2 { // 右
			left = r(x, y-1, board, chars, 4, q) // 递归往左
		}

		if dir != 3 { // 下
			up = r(x-1, y, board, chars, 1, q) // 递归往上
		}
		if dir != 4 { // 左
			right = r(x, y+1, board, chars, 2, q) // 递归往右
		}
		q[p] = false
	}

	return right || down || left || up

}

/*
总结:

    时间花了俩小时：耗时 1:55:50, 太长时间了！
	这种节奏不对!

	这个题是回溯类问题,必定遵循模板,因此需要遵循基本思路

	   1. 确定子问题
	   2. 确定边界
	   3. 核心处理逻辑

	1. 思考花费将近半小时, 在套以前做过的上下边界，前缀和的方式, 结果发现不对
	2. 没有专注于问题本身，后来定义出子问题后，没有快速的转换为代码
	3. 时间太长，脑子有点懵了
	4. 回溯中状态保存的问题，没想透彻

*/

// @lc code=end

/*
// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"SEE"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCB"\n
// @lcpr case=end

*/

