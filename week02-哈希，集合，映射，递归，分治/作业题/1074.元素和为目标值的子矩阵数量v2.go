/*
 * @lc app=leetcode.cn id=1074 lang=golang
 *
 * [1074] 元素和为目标值的子矩阵数量
 */

// @lc code=start
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	/*
		   关键信息：
			 输入： 一个二维矩阵，和目标值
			 输出： 非空子矩阵的数量
			 逻辑： 子矩阵元素和等于目标值的子矩阵的数量

		   思路：

			1. 根本目标是找矩阵，找到计算,然后比较
			2. 从上往下找, 从左往右
			   1. 根据当前位置，遍历下一个或多个的位置, 二维矩阵中，收缩的例子
				for i, _:= range matrix {
				   // i为当前位置, 从当前位置开始，往后找，如何找呢？

				   为什么要在这里定义?
					sum := make([]int, len(matrix[0])) // sum为数组，数组长度为列的宽度

				   for _, row := range matrix[i:] {
					   // 从当前位置开始，遍历后面的, 对应的上面会收缩的逻辑
				   }
				}
			   2. 计算每行中每一列的值，求sum中对应列
					上面缩进时，重新申请空间sum
					for c, v := range row {
						sum[c] += v // 更新每列的元素和
					}
					一行中每列的累积和

					然后判断，这一行中的各个列的sum值与目标的关系，循环，依次往右累加
					找到后ans累加

					下一行时，sum中还保留着上一行的记录，sum经过累加后，有两行的信息
				3. 前缀和(前面列的sum)加入到map中，判断最后一个等不等目标值时，通过当前所有列的累计和-k,等于前面列的累计和
				  快速在map中找到，减少计算量

			 这种是属于变形

			 能不能用分治?
				1.子问题是什么？如何分?
				2.计算完毕，需要合，怎么合？
				3.有没有明显的状态层级?

			 官方解法：
				枚举遍历,二维矩阵上下边界

				  1. 每行为一个单位, 从左往右
					for i := range matrix { // 枚举上边界
						sum := make([]int, len(matrix[0]))
						for _, row := range matrix[i:] { // 枚举下边界
							for c, v := range row {
								sum[c] += v // 更新每列的元素和
							}
							ans += subarraySum(sum, target)
						}
					}
					return
				  2.

	*/

	ans := 0
	scaned := map[int]bool{}
	r(0, 0, 0, 0, &matrix, &ans, target, scaned)
	return ans
}

/*
 总结：

*/

// @lc code=end

/*
// @lcpr case=start
// [[0,1,0],[1,1,1],[0,1,0]]\n0\n
// @lcpr case=end

// @lcpr case=start
// [[1,-1],[-1,1]]\n0\n
// @lcpr case=end

// @lcpr case=start
// [[904]]\n0\n
// @lcpr case=end

*/

