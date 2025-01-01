
/*
 * @lc app=leetcode.cn id=1000 lang=golang
 * @lcpr version=20004
 *
 * [1000] 合并石头的最低成本
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func mergeStones(stones []int, k int) int {

	/*
		   输入: n堆石头排一排，stones[i]中表示第i堆中石头的数量

		   输出: 所有石头合并为一堆的最低成本，无法合并返-1

		     成本指移动k堆石头的总数量, 合并完毕后组成一个新的堆
		     如果合并完后，不满足k堆，则无法合并为一堆

		    从哪一堆开始合并呢？
		        可以从末尾开始，处理玩一堆，再从末尾开始处理
		        先处理这个区间 [n-k, n-1]

			上述思路不对，仍旧是需要划分子问题
			原始问题是从l到r堆石头合成一堆，成本最小
				子问题是需要缩小规模的,但不一定能合成一堆
				因此f[l,r,i], 从l开始到r连续的多堆石头，合成i堆最低成本

	*/

}

// @lc code=end

/*
// @lcpr case=start
// [3,2,4,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [3,5,1,2,6]\n3\n
// @lcpr case=end

*/

