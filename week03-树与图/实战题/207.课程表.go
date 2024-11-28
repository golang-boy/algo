/*
 * @lc app=leetcode.cn id=207 lang=golang
 * @lcpr version=20003
 *
 * [207] 课程表
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {

	/*
		参数
			1.必须修的课程数
			2.课程依赖图
		输出
			能否把课程都修完

		是一个有向图，如果存在环，则不能修完，否则可以修完。
		转换为邻接数组,构造图

	*/

	g := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, course := range prerequisites {
		ai := course[0]
		bi := course[1]

		if g[bi] == nil {
			g[bi] = []int{}
		}

		if g[ai] == nil {
			g[ai] = []int{}
		}
		g[bi] = append(g[bi], ai)
		inDegree[ai]++
	}

	var q = []int{}
	// 入度等于0的课程，入队
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	var lesson = []int{}
	// 处理队列数据
	for len(q) != 0 { // 队列不空
		// 弹出一个x
		x := q[0]
		q = q[1:]

		lesson = append(lesson, x)

		// 查看下一层的课程
		for _, y := range g[x] {
			inDegree[y]--
			if inDegree[y] == 0 {
				q = append(q, y)
			}
		}
	}

	return len(lesson) == numCourses
}

/*
 总结：

    该题为拓扑排序的典型题,重点理解入度的概念

	   1. 入度为0的节点，入队
	   2. 队列不空时，出队，进行逻辑处理
	   3. 判断该节点的可到达的下一层节点
	     1. 此时下一层节点的入队都--
		 2. 如果减完后，等于0,则入队
		 3. 重复上述逻辑

*/

// @lc code=end

/*
// @lcpr case=start
// 2\n[[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[[1,0],[0,1]]\n
// @lcpr case=end

*/

