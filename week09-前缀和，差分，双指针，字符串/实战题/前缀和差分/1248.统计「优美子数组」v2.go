/*
 * @lc app=leetcode.cn id=1248 lang=golang
 * @lcpr version=20004
 *
 * [1248] 统计「优美子数组」
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	/*
		输入nums和k,连续子数组中有k个奇数数字，则为优美子数组
		输出优美子数组数量

		子段问题可以考虑前缀和去处理，或者看是否能转换为前缀和问题

		为什么？
			处理子段问题，【l,r】时间复杂度是O(n2)，前缀和方式降低到O(n)

		前缀和思路：
			1. 将奇数变为1，偶数变为0。此时nums和的范围为[0,n]
			2. s[l,r] = (nums[l]%2+...+nums[r]%2) = s[r]-s[l-1] = k
			3. s[l-1] = s[r] - k
			4. 上述表达式，假设r为i,l-1为寻找的另一半，通过循环i,即可得到j的前缀和
			5. 看j的前缀和的数量
			6. 数量统计使用count数组或者map进行，此处范围为[0,n], 因此使用数组
	*/

	ans := 0
	n := len(nums)

	s := make([]int, n+1)

	// 和的范围为[0,n]总共n+1个数
	count := make([]int, n+1)

	// 求前缀和, 并转化
	s[0] = 0
	for i := 1; i < n+1; i++ {
		s[i] = s[i-1] + nums[i-1]%2
	}

	count[s[0]]++ // 和为0的位置有一个
	for i := 1; i <= n; i++ {
		if s[i]-k >= 0 {
			ans += count[s[i]-k] // 和为S[i]-k的位置为满足条件的位置，进行累计
		}
		count[s[i]]++ // 循环时每个s[i]位置进行累加
	}

	return ans
}

/*
总结：

O(n)
​

*/

// @lc code=end

/*
// @lcpr case=start
// [1,1,2,1,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,4,6]\n1\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,1,2,2,1,2,2,2]\n2\n
// @lcpr case=end

*/

