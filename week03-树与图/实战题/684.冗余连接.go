/*
 * @lc app=leetcode.cn id=684 lang=golang
 * @lcpr version=20003
 *
 * [684] 冗余连接
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findRedundantConnection(edges [][]int) []int {

	/*
		关键信息
			1. 二维数组，表示一张图，内层数组表示俩节点间有边
			2. 输出一维数组，表示要删哪条边

			3. 删除边后，图中没有环，不连通
			4. 如果有多种方法，则删除edges中最后一条满足条件的边

		分析：
			1. 需要将边的存储方式，变为邻接数组的存储形式，便于遍历
			2. 找到出边大于1的节点，尝试依照顺序删除一个边，深度优先遍历
			3. 看有没有环,有则不对，在选一条边进行删除
	*/

	var graph [][]int
	var n = 0

	//拿到图结构

	//  找最大节点，用于分配入边数组空间大小
	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]
		n = max(n, max(x, y))
	}

	var hascycle = false

	graph = make([][]int, n+1) // 存储 0 -> n 个点, 总共n+1
	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]

		if graph[x] == nil {
			graph[x] = []int{}
		}

		if graph[y] == nil {
			graph[y] = []int{}
		}

		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)

		// 找环
		//  每加一条边，就找一次, 刚开始肯定没有，直到最后一条加上后有了，此时就返回
		var visited = make([]bool, n+1) // 每次遍历图时，这个visited要重置
		dfs(x, 0, graph, &visited, &hascycle)
		if hascycle == true {
			return edges[i]
		}
	}

	return []int{}
}

func dfs(x int, fa int, g [][]int, visited *[]bool, hascycle *bool) {
	(*visited)[x] = true
	for i := 0; i < len(g[x]); i++ {
		y := g[x][i]
		if y == fa {
			continue
		}

		if !(*visited)[y] {
			dfs(y, x, g, visited, hascycle)
		} else {
			*hascycle = true
			return
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
总结：

   分析与实际实现时有差异的
      需要从实际程序执行的角度，确定问题的处理
	      而这需要对数据结构的处理有清晰的认识，基于数据结果进行算法设计，以及问题分析

	进行数据结构转换的过程中，相当于重建图

	因为是基环树，根据定义只有一条多余的边
	因此重建图时，前面的都是无环，直到找到某一条加入的边后形成环,直接返回该边


    深度遍历中的关键点：
	    1. 图中可能有环，访问过的节点需要标记
		2. 从父节点进入，不能再回去




*/

// @lc code=end

/*
// @lcpr case=start
// [[1,2], [1,3], [2,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2], [2,3], [3,4], [1,4], [1,5]]\n
// @lcpr case=end

*/

