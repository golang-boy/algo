/*
 * @Author: 刘慧东
 * @Date: 2024-11-22 16:52:08
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-22 16:52:08
 */
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

	if x == 0 {
		return 0.0
	}

	// if x < 1 && x > -1 && n > math.MaxInt32 {
	// 	return 0.0
	// }

	if n > 0 {
		return calc(x, n)
	}

	if n < 0 {
		return 1.0 / calc(x, -n)
	}

	return 1.0
}

func calc(x float64, n int) float64 {
	tmp := myPow(x, n/2)
	ans := tmp * tmp // 通过计算，少很多次递归
	if n%2 == 1 {
		return ans * x
	}
	return ans
}

// @lc code=end

