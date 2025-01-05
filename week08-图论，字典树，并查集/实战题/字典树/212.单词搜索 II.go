/*
 * @lc app=leetcode.cn id=212 lang=golang
 * @lcpr version=20004
 *
 * [212] 单词搜索 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findWords(board [][]byte, words []string) []string {
	/*
		输入字符表格，和单词列表

		输出在表格中的所有单词, 要求按字母序在相邻单元格内的字母构成


		1. 根据words创建trie树
		2. 在图中进行深度遍历，每遍历一个字符，在trie树中进行同步查找
		3. trie中找到后，就清理该字符，层层递归

	*/

	// 1. 根据words创建trie树

	root := MakeNode()
	for _, s := range words {
		insert(s, root)
	}

	// 2. 搜索图中每个点, 判定是否在trie中
	m := len(board)
	n := len(board[0])

	visit := make([][]bool, m)
	ans := map[string]struct{}{}
	res := []string{}
	str := ""

	for i := range visit {
		visit[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visit[i][j] = true
			// 与trie同时深度搜索
			dfs(board, i, j, root, visit, ans, str)
			visit[i][j] = false
		}
	}

	for k := range ans {
		res = append(res, k)
	}

	return res
}

func dfs(board [][]byte, x, y int, cur *Node, visit [][]bool, ans map[string]struct{}, str string) {
	if cur == nil {
		return
	}

	m := len(board)
	n := len(board[0])

	ch := rune(board[x][y])
	if cur.child[ch] == nil {
		// 找到末尾了没有
		return
	}
	// 递归处理trie树的值
	next := cur.child[ch]

	str += string(ch)
	if next.count > 0 {
		// 找到了单词，追加到结果中
		// *ans = append(*ans, str)
		ans[str] = struct{}{}
	}

	if len(next.child) == 0 {
		delete(cur.child, ch)
		return
	}

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	// for循环处理board[x][y]的四个方向
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		if nx < 0 || ny < 0 || nx >= m || ny >= n {
			continue
		}

		if visit[nx][ny] {
			continue
		}

		visit[nx][ny] = true
		dfs(board, nx, ny, next, visit, ans, str)
		visit[nx][ny] = false
	}
	// 还原, 这步应该可有可无, 因为str不是全局的
	// str = str[:len(str)-1]
	return
}

type Node struct {
	count int
	child map[rune]*Node
}

func insert(s string, root *Node) {
	cur := root
	for _, ch := range s {
		// 该字符未找到
		if cur.child[ch] == nil {
			cur.child[ch] = MakeNode()
		}
		cur = cur.child[ch]
	}
	cur.count++
}

func MakeNode() *Node {
	return &Node{
		count: 0,
		child: make(map[rune]*Node),
	}
}

/*
总结：
	1. 建立一颗trie树，将words列表中的词放入
	2. 二维表格中根据方向数组进行深度遍历，且通过visit数组进行去重记录
	3. 遍历表格式中的每个字符，每取一个字符进行递归时，在trie树中进行判断，是否是一个单词，是就追加到结果中
*/

// @lc code=end

/*
// @lcpr case=start
// [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n["oath","pea","eat","rain"]\n
// @lcpr case=end

// @lcpr case=start
// [["a","b"],["c","d"]]\n["abcb"]\n
// @lcpr case=end

*/

