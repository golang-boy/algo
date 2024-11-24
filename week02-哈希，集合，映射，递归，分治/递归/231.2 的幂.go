/*
 * @lc app=leetcode.cn id=231 lang=golang
 * @lcpr version=20003
 *
 * [231] 2 的幂
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPowerOfTwo(n int) bool {

	/*
		   输入:整数n
		   逻辑：判断n是否是2的幂次方, 判定条件 n == 2的x次方

		   思路：
		      2*2*2.... 有x个, 最后结果与n比较

			  每次递归时乘2,什么时候结束？

	*/

	return r(n)

}

func r(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n%2 == 1 {
		return false
	}
	return r(n / 2)
}

/*
总结：
    没有冷静清晰的分析边界条件，直接使用math.Pow函数进行计算，忽略了计算的时间复杂度

	递归解法关键点：
	   1. 确定递归的问题，什么是重叠子问题
	         上述问题进行转换：x个2相乘 == n转换为 n是奇数还是偶数的问题
	            如果n是偶数，则除以2，产生新的n，重复上述判断
				如果n是奇数, 则必定不是2的x次幂

	   2. 确定递归边界
	      n == 0
		  n == 1     除2==1为偶数
		  n %2 == 1  余数为1表示奇数

	   3. 递归现场的保护与还原
	      无


*/

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 16\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/

