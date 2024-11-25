/*
 * @lc app=leetcode.cn id=1074 lang=golang
 *
 * [1074] 元素和为目标值的子矩阵数量
 */

// @lc code=start
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	n := len(matrix)
	m := len(matrix[0])

	ans := 0

	for i := 0; i < n; i++ {
		sum := make([]int, m)
		for j := i; j < n; j++ {
			pre := 0
			q := map[int]int{0: 1} // 预置0:1,正好pre-target == 0，表示从左往右，等于目标时，依次命中

			for k := 0; k < m; k++ {
				sum[k] += matrix[j][k]

				// 从左往右,
				//   k==0时， 第一个格子， k==1时，前俩个格子，k==2时，三个格子
				pre += sum[k]
				if v, ok := q[pre-target]; ok {
					ans += v
				}
				q[pre] += 1

				// 如何确定第二个，第三个格子是否满足条件
				//   pre-target == 前一个pre时，拿到中间格子是否命中的计数
				//  这是最精妙的设计，通过计算拿到结果
			}
		}
	}

	return ans
}

// @lc code=end

