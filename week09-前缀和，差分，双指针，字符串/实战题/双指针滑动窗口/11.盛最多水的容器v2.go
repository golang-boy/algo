/*
 * @Author: 刘慧东
 * @Date: 2025-01-10 16:15:21
 * @LastEditors: 刘慧东
 * @LastEditTime: 2025-01-10 16:51:42
 */

/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=20004
 *
 * [11] 盛最多水的容器
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxArea(height []int) int {

	/*

		输入数组，值为高度, 索引为宽度

		输出为可存储的最大水量

		最大水量  =  max(min(height[i],height[j]) * (j-i)   i<j)

		双指针, 什么时候左移，什么时候右移

		  i的值小于j的值时，i移动，否则j移动，为什么？

		  	固定一端后，移动另一端，高度更高的, 能容更多水，因此需要另一端找比当前值更高的

	*/

	n := len(height)

	ans := 0

	j := n - 1

	for i := 0; i < j; {
		ans = max(ans, min(height[i], height[j])*(j-i))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}

	}

	// [1,2,1]
	// 1 * 2 = 2
	// 1 * 1 = 1
	// 1 * 1 = 1

	return ans
}

/*

总结:
	什么时候用前缀和?
	什么时候用单调队列?
	什么时候双指针扫描?

	如果满足区间减法性质：
	    [l,r]可以由[1,r]和[1,l-1]导出
	此时可以使用前缀和与差分思想


	如果固定一个点，另一个点维护的信息是一个点，此时双指针扫描

	如果是多个点构成的候选集合，则单调队列

*/

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [1,8,6,2,5,4,8,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

*/

