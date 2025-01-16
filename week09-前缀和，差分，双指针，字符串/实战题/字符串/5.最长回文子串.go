/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=20004
 *
 * [5] 最长回文子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestPalindrome(s string) string {

	n := len(s)
	b := 131
	p := 1e9 + 7

	preH := make([]int64, n+1)
	sufH := make([]int64, n+2)
	powB := make([]int64, n+1)
	powB[0] = 1

	for i := 1; i <= n; i++ {
		preH[i] = (preH[i-1]*b + s[i-1] - 'a' + 1) % p
		powB[i] = powB[i-1] * b % p
	}

	for i := n; i >= 1; i-- {
		sufH[i] = (sufH[i+1]*b + s[i-1] - 'a' + 1) % p
	}

	// 计算[l,r]区间的哈希
	calcHash := func(l, r int) int64 {
		return ((preH[r+1]-preH[l]*powB[r-l+1])%p + p) % p
	}

	// 计算[l,r]区间反向的哈希
	calcReverseHash := func(l, r int) int64 {
		return ((sufH[l+1]-preH[r+2]*powB[r-l+1])%p + p) % p
	}

	l := 0
	r := len(s) - 1
	for l < r {
		mid := (l + r) / 2
		if isPalindrome(s, l, r, mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/

