/*
 * @lc app=leetcode.cn id=304 lang=golang
 * @lcpr version=20004
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	m := len(matrix[0])

	sum := make([][]int, n+1)
	for i := range sum {
		sum[i] = make([]int, m+1)
	}

	// 计算前缀和
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	return NumMatrix{
		sum: sum,
	}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1++
	row2++
	col1++
	col2++
	return this.sum[row2][col2] - this.sum[row2][col1-1] - this.sum[row1-1][col2] + this.sum[row1-1][col1-1]
}

/*
总结：
    二维前缀和思想：
        s[i][j] 表示从（0,0）到（i,j)的和

        s[i][j] = s[i-1][j]+s[i][j-1]-s[i-1][j-1] + a[i][j]

        sum[p,q,i,j] 表示从（p,q)到（i,j)的区域和

        sum[p,q,i,j] = s[i][j] - s[i][q-1]-s[p-1][j] + s[p-1][q-1]

        sumRegion实践复杂度为O(1)
*/

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

/*
// @lcpr case=start
// ["NumMatrix","sumRegion","sumRegion","sumRegion"][[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]\n
// @lcpr case=end

*/

