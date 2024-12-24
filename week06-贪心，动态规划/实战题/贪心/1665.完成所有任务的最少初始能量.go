/*
 * @lc app=leetcode.cn id=1665 lang=golang
 * @lcpr version=20004
 *
 * [1665] 完成所有任务的最少初始能量
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumEffort(tasks [][]int) int {
	/*

		输入：任务二维矩阵，一维表示第一个任务，二维元素表示任务消耗的实际能量，和最低能量
		       1 <= actual​i <= minimumi <= 10000
		输出：完成所有任务最少初始能量

		初始能量 >= 起始能量最大的任务

		先完成哪个，再完成哪个？
		 1. 按最小起始能量来排序，先完成起始能量需要的最大的


		二分是需要在有序的序列上求值

	*/

	slices.SortFunc(tasks, func(a, b []int) int {
		return (a[0] - a[1]) - (b[0] - b[1])
	})

	ans := 0
	for i := len(tasks) - 1; i >= 0; i-- {
		ans = max(tasks[i][1], ans+tasks[i][0])
	}
	return ans
}

/*
总结：
	1. 唯一能想到的是尝试二分
	2. 另外一个靠主管猜测是mnimum-actual进行排序，但无法证明

	本题主要是需要推导公式, 属于贪心中的邻项交换法

*/

// @lc code=end

/*
// @lcpr case=start
// [[1,2],[2,4],[4,8]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,3],[2,4],[10,11],[10,12],[8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,7],[2,8],[3,9],[4,10],[5,11],[6,12]]\n
// @lcpr case=end

*/

