/*
 * @lc app=leetcode.cn id=685 lang=golang
 * @lcpr version=20003
 *
 * [685] 冗余连接 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findRedundantDirectedConnection(edges [][]int) []int {

	var hascycle = false
	var n = 0

	// 找到n
	for _, e := range edges {
		ui := e[0] // 父节点
		vi := e[1] // 子节点
		n = max(max(n, vi), ui)
	}

	parent := make([][]int, n+1)
	for _, e := range edges {
		ui := e[0] // 父节点
		vi := e[1] // 子节点
		if parent[vi] == nil {
			parent[vi] = []int{}
		}
		parent[vi] = append(parent[vi], ui)
	}

	var g = make([][]int, n+1)

	for _, e := range edges {
		ui := e[0] // 父节点
		vi := e[1] // 子节点

		if g[ui] == nil {
			g[ui] = []int{}
		}
		if g[vi] == nil {
			g[vi] = []int{}
		}

		g[ui] = append(g[ui], vi)

		visited := make([]bool, n+1)
		dfs(ui, g, visited, &hascycle) // 找环
		if hascycle {
			return e
		}

	}

	// 解决不了，一节点二父的情况
	for x, f := range parent {
		if len(f) >= 2 {
			return []int{f[0], x}
		}
	}

	return []int{}
}

func dfs(n int, g [][]int, visited []bool, hascycle *bool) {
	visited[n] = true
	for _, x := range g[n] {
		if !visited[x] {
			dfs(x, g, visited, hascycle)
		} else {
			*hascycle = true
			return
		}
	}

	return
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2],[1,3],[2,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[2,3],[3,4],[4,1],[1,5]]\n
// @lcpr case=end

*/

