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

	for i := 0; i < n; i++ {
		for ; i < j && height[i] >= height[j]; j-- {

			// 是否需要计算
			if (j < n-1 && height[j] > height[j+1]) || j == n-1 {
				h := min(height[i], height[j])
				w := j - i
				ans = max(ans, h*w)
			}
		}

		if i < j && height[i] < height[j] {
			if (i > 0 && height[i] > height[i-1]) || i == 0 {
				h := min(height[i], height[j])
				w := j - i
				ans = max(ans, h*w)
			}
		}
	}

	// [1,2,1]
	// 1 * 2 = 2
	// 1 * 1 = 1
	// 1 * 1 = 1

	return ans
}

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

