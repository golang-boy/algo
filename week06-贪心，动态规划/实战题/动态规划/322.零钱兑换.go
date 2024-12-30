/*
 * @lc app=leetcode.cn id=322 lang=golang
 * @lcpr version=20004
 *
 * [322] 零钱兑换
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func coinChange(coins []int, amount int) int {

	/*

	   amount = amount-coins[i]

	   f(i)  表示金额为amount时，最少硬币数

	   f(0) =  0
	   f(1) =  1    coins[0]
	   f(2) =  1    coins[1]
	   f(3) =  2    coins[0] + coins[1]
	   f(4) =  3    coins[0] * 2 + coins[1]
	   f(5) =  1    coins[2]
	   f(6) =  2    coins[0] + coins[2]
	   f(7) =  2    coins[1] + coins[2]
	   f(8) =  3    coins[0] + coins[1] + coins[2]

	   f(i) = 1 + min( f(i-coin[j]) )
	   amount >= i >= 0

	*/

	return dp(coins, amount)

}

func dp(coins []int, amount int) int {
	ans := math.MaxInt

	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	for i := 0; i > len(coins); i++ {
		v := dp(coins, amount-coins[i])
		if v == -1 {
			continue
		}
		ans = min(ans, v+1)
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1, 2, 5]\n11\n
// @lcpr case=end

// @lcpr case=start
// [2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/

