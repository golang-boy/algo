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

		朴素思路：
			1. 循环遍历数组，索引为i
			2. 从i开始取连续k个奇数，取到则计数加一
			3. 直到末尾或者（n-k）为止
	*/

	ans := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		cnt := 0
		for j := i; j < n; j++ {

			if nums[j]%2 == 1 {
				cnt++
			}

			if cnt == k {
				ans += 1
			}
		}
	}

	return ans
}

/*
总结：
	朴素算法可以算出，某些执行用例耗时太多，时间复杂度不达标

	外层循环 (for i := 0; i < n; i++) 从 0 遍历到 n-1，总共运行 n 次。
内层循环 (for j := i; j < n; j++) 从 i 遍历到 n-1。在最坏情况下（即当 i = 0 时），内层循环会执行 n 次；当 i = 1 时，执行 n-1 次；依此类推，直到 i = n-1 时，内层循环执行 1 次。

计算总的操作次数

	内层循环对于每个 i 的执行次数为：
		当 i = 0 时：n 次
		当 i = 1 时：n-1 次
		当 i = 2 时：n-2 次
		...
		当 i = n-1 时：1 次
	因此，总的操作次数为：

n+(n−1)+(n−2)+...+1= n(n+1)/ 2

O(n2)
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

