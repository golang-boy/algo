/*
 * @lc app=leetcode.cn id=743 lang=golang
 * @lcpr version=20004
 *
 * [743] 网络延迟时间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func networkDelayTime(times [][]int, n int, k int) int {

	dist := make([]int, n+1)

	for i := range dist {
		dist[i] = 1000000
	}

	dist[k] = 0

	for i := 1; i <= n-1; i++ {

		flag := false

		for _, edge := range times {
			x := edge[0]
			y := edge[1]
			z := edge[2]

			if dist[x]+z < dist[y] {
				dist[y] = dist[x] + z
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, dist[i])
	}

	if ans == 1000000 {
		return -1
	}
	return ans
}

/*
总结：
	最短路问题中，bellman-ford经典算法
	Bellman-Ford 算法是一种用于寻找加权图中最短路径的算法，特别适用于处理含有负权边的图

	图：由顶点和边组成的集合。边可以有正权或负权。
	最短路径：从一个源顶点到其他所有顶点的路径，权重和最小。

	算法步骤
		初始化：
			将源顶点到自身的距离设为 0，其他所有顶点的距离设为无穷大。

		松弛操作：
			对于图中的每一条边 (u, v)，检查是否可以通过 u 以更小的总权重到达 v：

				如果 distance[u] + weight(u, v) < distance[v]，则更新 distance[v]。

		重复步骤 2：
			重复执行上述松弛操作 |V| - 1 次（|V| 是图中顶点的数量）。这是因为在最坏情况下，最短路径可能需要经过 |V| - 1 条边。
			V表示顶点数

		检测负权回路：
			再次对所有边进行松弛操作。如果在此过程中仍能更新任何顶点的距离，则说明图中存在负权回路。


*/

// @lc code=end

/*
// @lcpr case=start
// [[2,1,1],[2,3,1],[3,4,1]]\n4\n2\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,1]]\n2\n1\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,1]]\n2\n2\n
// @lcpr case=end

*/

