
/*
 * @lc app=leetcode.cn id=210 lang=golang
 * @lcpr version=20004
 *
 * [210] 课程表 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {

	/*
	 1. 先构建邻接数组
	 2. 构建过程中,计算入度
	 3. 入度为零的入队
	*/

	var g [][]int = make([][]int, numCourses)
	var inDegree = make([]int, numCourses)

	var q = []int{}
	var ans = []int{}

	for _, pre := range prerequisites {
		// ai依赖bi
		ai := pre[0]
		bi := pre[1]

		if g[bi] == nil {
			g[bi] = []int{}
		}
		if g[ai] == nil {
			g[ai] = []int{}
		}

		g[bi] = append(g[bi], ai)
		inDegree[ai]++
	}

	// 入队
	for c, degree := range inDegree {
		if degree == 0 {
			q = append(q, c)
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		ans = append(ans, x)

		// 处理x节点可到达的
		for _, v := range g[x] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(ans) == numCourses {
		return ans
	}

	return []int{}
}

// @lc code=end

/*
// @lcpr case=start
// 2\n[[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[[1,0],[2,0],[3,1],[3,2]]\n
// @lcpr case=end

// @lcpr case=start
// 1\n[]\n
// @lcpr case=end

*/

