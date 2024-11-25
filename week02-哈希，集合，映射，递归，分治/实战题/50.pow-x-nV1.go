/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {

	/*
		分治算法练习
	*/

	if n > math.MaxInt32 && x < 1 && x > 0 {
		return 0.0
	}

	if x == 0.0 {
		return 0.0
	}

	if n == 0 {
		return 1.0
	}
	if n > 0 {
		return x * myPow(x, n-1)
	}

	if n < 0 {
		r := (x * myPow(x, -n-1))
		return 1.0 / r
	}

	return 0
}

/*
  上述问题划分时，分为一个和后面的n-1个, 这种方式和循环去乘没什么区别，依旧是一个乘以一个。
    因此，上述方法不通过，消耗太多内存

  正确的方式是，分为一半的一半，每次一半，最后合起来完成计算
*/

// @lc code=end

