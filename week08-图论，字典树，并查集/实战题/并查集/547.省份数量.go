/*
 * @lc app=leetcode.cn id=547 lang=golang
 * @lcpr version=20004
 *
 * [547] 省份数量
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findCircleNum(isConnected [][]int) int {
	/*
	   省份：直接或间接相连的城市为一个省份
	   输入一个二维矩阵, 返回其中省份的数量

	   思路：
	   1. 矩阵中相连的城市，构成图，进行深度搜索统计
	   2. 每个城市一个集合，相连的合起来，最后看有几个集合(并查集)
	*/

	n := len(isConnected)
	// makeSet
	fa := make([]int, n)

	// 初始化并查集
	for i := 0; i < n; i++ {
		// 每个节点保存其父节点的索引
		// 初始化时，自己的父节点为自己
		fa[i] = i
	}

	/* 为什么一个数组能搞定并查集？
	   1. 并查集本质是只关心每个节点所在的集合，不关心集合对应的树结构
	   2. 每个节点所在的集合有根节点确定

	   并查集核心是树结构，每个子节点包含指向父节点的信息。
	   俩元素合并，即所在集合合并, 元素的父节点信息需要指向新的父节点

	   集合合并相当于俩树结构合并，较浅的树合并到较深的上面（按秩合并）+ 路径压缩

	   路径压缩指向同一个根节点的节点, 将其父节点设置为根节点，从深层树结构变为浅层树结构

	*/

	// 每条边合并俩集合

	// 给定的矩阵为邻接矩阵
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				unionSet(fa, i, j)
			}
		}
	}

	// 有几个根，表示有几个省份
	ans := 0
	for i := 0; i < n; i++ {
		// i的根等于i,表示一个独立集合
		if find(fa, i) == i {
			ans++
		}
	}
	return ans
}

func find(fa []int, i int) int {
	if i == fa[i] {
		// 找到了根节点
		return i
	}
	// 没找到时，找fa[i]节点的父节点
	// 找i的父节点，递归直到找根父节点, 返回
	fa[i] = find(fa, fa[i])
	// 为什么赋值在返回呢？
	//  这是在进行路径压缩, 优化
	return fa[i]
}

func unionSet(fa []int, i, j int) {
	// 找i所在集合的根
	x := find(fa, i)
	// 找j所在集合的根
	y := find(fa, j)

	// 根不相等，需要合并, x的根变为y
	// 此处实现时，没有进行按秩合并优化
	//  如果要按秩合并，需要记录树的深度，小的合到大的
	//    均摊复杂度为O(a(n)) 只采用其中一种优化，则为O(log(n)), 前者接近常数, a(n) <= 5
	if x != y {
		fa[x] = y
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,0],[1,1,0],[0,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0],[0,1,0],[0,0,1]]\n
// @lcpr case=end

*/

