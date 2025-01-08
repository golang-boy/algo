/*
 * @lc app=leetcode.cn id=1109 lang=golang
 * @lcpr version=20004
 *
 * [1109] 航班预订统计
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	/*
		输入:航班预定表，和n个航班
			bookings[i] 表示第i条预定记录, 在[first,last] 航班范围内的每个航班预定的对应seats座位数
		输出:返回一个长度为n的数组，每个元素表示每个航班预定的座位总数


		预定记录数量为 m:=len(bookings)

		差分思想结论：
			1. B1 = A1, Bi=Ai-A(i-1)  差分数组B的前缀和为原数组A
			2. 原数组A的[l,r]区间都加d, 差分数组的变化为：B(l)+d, B(r+1)-d

	*/

	// m := len(bookings)

	// 范围为1,n, 前面+一个，后面+一个, 0,n+1
	delta := make([]int, n+2) // 多一个位置是为了减去边界多余判断
	sum := make([]int, n+1)

	for i := range bookings {
		f := bookings[i][0]
		l := bookings[i][1]
		seats := bookings[i][2]

		// 2. 原数组A的[l,r]区间都加d, 差分数组的变化为：B(l)+d, B(r+1)-d
		delta[f] += seats
		delta[l+1] -= seats
	}

	// 求差分数组delta的前缀和, 结果为累加后的A
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + delta[i]
	}

	return sum[1:]
}

/*
总结：

	该题中并没有求原数组的差分数组，为什么直接可以？
		因为原始数组都是0，在0的基础上进行+操作, 因此可行

	1. 先算原始数组A[l,r]区间+d, 等价于 B[l]+d , B[r+1]-d
	2. 然后针对差分数组B，进行前缀和还原

	时间复杂度O(n+m)
*/

// @lc code=end

/*
// @lcpr case=start
// [[1,2,10],[2,3,20],[2,5,25]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,10],[2,2,15]]\n2\n
// @lcpr case=end

*/

