/*
 * @lc app=leetcode.cn id=374 lang=golang
 * @lcpr version=20004
 *
 * [374] 猜数字大小
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left := 1
	right := n
	ans := 0

	for left <= right {
		mid := (left + right) / 2

		switch guess(mid) {
		case -1:
			right = mid - 1
		case 0:
			ans = mid
			return ans
		case 1:
			left = mid + 1
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 10\n6\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

// @lcpr case=start
// 2\n1\n
// @lcpr case=end

*/

