/*
 * @lc app=leetcode.cn id=684 lang=golang
 * @lcpr version=20004
 *
 * [684] 冗余连接
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	/*



	 */

	fa := make([]int, 1001)

	for i := 0; i < 1001; i++ {
		fa[i] = i
	}

	ans := []int{}

	for _, edge := range edges {
		if find(fa, edge[0]) == find(fa, edge[1]) {
			ans = edge
			break
		}
		join(fa, edge[0], edge[1])
	}

	return ans
}

func join(fa []int, x, y int) {

	x = find(fa, x)
	y = find(fa, y)

	if x != y {
		fa[x] = y
	}
}

func find(fa []int, x int) int {
	if fa[x] == x {
		return x
	}

	fa[x] = find(fa, fa[x])
	return fa[x]
}

/*
总结：
	1. 初始化并查集, 每个点都是孤立的
	2. 循环处理每一条边，边两端的断点进入并查集查根，如果两个点都连接到同一个根，则说明已经合并连接过
	3. 如果不相等，则将两个端点进行合并操作
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

