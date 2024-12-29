/*
 * @lc app=leetcode.cn id=120 lang=golang
 * @lcpr version=20004
 *
 * [120] 三角形最小路径和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumTotal(triangle [][]int) int {

	/*

		定义：
			f[i] 第i层时，最小路径和
			f[i][j] 第i层第j个元素时，路径和

		方程：
			f[0] = triangle[0][0]
			f[1] = min(triangle[1][j+0], triangle[1][j+1]) + f[0]  j in [0,i-1]


			f[i] = min(triangle[i][j+0], triangle[i][j+1])) + f[i-1]
		边界：
		目标：

	*/

}

/*
总结：
    思路不对，定义有问题。
	该从底到上，由下一层确定当前层


*/

// @lc code=end

/*
// @lcpr case=start
// [[2],[3,4],[6,5,7],[4,1,8,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[-10]]\n
// @lcpr case=end

*/

