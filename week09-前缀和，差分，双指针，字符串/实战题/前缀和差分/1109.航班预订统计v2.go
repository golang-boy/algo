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

	*/

	// m := len(bookings)

	// 范围为1,n, 前面+一个，后面+一个, 0,n+1
	delta := make([]int, n+2)
	sum := make([]int, n+1)

	for i := range bookings {
		f := bookings[i][0]
		l := bookings[i][1]
		seats := bookings[i][2]

		delta[f] += seats
		delta[l+1] -= seats
	}

	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + delta[i]
	}

	return sum[1:]
}

/*
总结：
	时间复杂度O(n)
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

