/*
 * @lc app=leetcode.cn id=210 lang=golang
 * @lcpr version=20003
 *
 * [210] 课程表 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {

	g := make([][]int, numCourses)
	inDegree := make([]int, numCourses) // 总共这么多节点

	q := []int{}
	lesson := []int{}

	// 构造图

	for _, pre := range prerequisites {
		ai := pre[0]
		bi := pre[1]

		// g[bi]   =>   bi可到达ai
		if g[bi] == nil {
			g[bi] = []int{}
		}
		if g[ai] == nil {
			g[ai] = []int{}
		}

		g[bi] = append(g[bi], ai) // bi可以到达ai
		inDegree[ai]++            // ai的入度加一
	}

	// 找到入度为0的点
	for x, _ := range g {
		if inDegree[x] == 0 {
			q = append(q, x)
		}
	}

	// 处理队列
	for len(q) != 0 {
		x := q[0]
		q = q[1:]
		lesson = append(lesson, x)

		for _, y := range g[x] {
			inDegree[y]--
			if inDegree[y] == 0 {
				q = append(q, y)
			}
		}
	}

	if len(lesson) == numCourses {
		return lesson
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

