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

		  思路：分治
		    最小矩阵为1x1子矩阵
			横纵坐标变化

			每找到一个矩阵，判断元素和值是否等于目标值

			边界为，矩阵为1x1时
	*/

	ans := 0
	scaned := map[int]bool{}
	r(0, 0, 0, 0, &matrix, &ans, target, scaned)
	return ans
}

func r(x1, y1, x2, y2 int, matrix *[][]int, ans *int, target int, scaned map[int]bool) {
	for i := x1; i < len(*matrix); i++ {
		for j := y1; j < len((*matrix)[0]); j++ {
			for k := x2; k < len(*matrix); k++ {
				for h := y2; h < len((*matrix)[0]); h++ {

					if k < i || h < j {
						continue
					}

					key := calcKey(i, j, k, h)
					if !scaned[key] {
						// r(i, j, k, h, matrix, ans, target, scaned)

						sum := 0
						for m := i; m <= k; m++ {
							for n := j; n <= h; n++ {
								sum += (*matrix)[m][n]
							}
						}
						if sum == target {
							*ans += 1
						}
						scaned[key] = true
					}
				}
			}
		}
	}
}

func calcKey(x1, y1, x2, y2 int) int {
	return x1*1000001 + x2*10001 + y1*101 + y2
}

/*
 总结：
    上述做法纯暴力解决,map并没有起到应有的作用, 通过35个用例，还有5个没通过，执行超时
	  基本思路是
	    1. 从{x1,y1}起始点出发, {x2，y2}为矩阵的对角点, 一对点表示一个矩阵
		2. 两个点依次前进，计算之间包含的点的和
		3. 如果等于target, 则累加
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

