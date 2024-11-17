/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleAreaV1(heights []int) int {

	/*
		    关键信息：
			  输入： 表示高度的数组
			  输出： 最大矩形面积

			  变量：
			     宽度会累加变化
				 高度随着循环变化

			分析：

			   area :=  height * (width * n)

			   1. 需要循环变量heigths数组
			   2. 从当前heights[i]的高度与后面的依次比较,直到小于当前height，此时，找到最大的宽度(向右扩展)
			          如果前一个的height > 当前height (向左扩)
			   3. 计算面积，并更新最大面积
			   4. 继续循环，从heights[i+1]开始

	*/

	maxArea := 0

	for i := 0; i < len(heights); i++ {

		width := 1
		for j := i + 1; j < len(heights); j++ {
			if heights[i] > heights[j] {
				break
			}
			width++
		}

		// 宽度需要保存, 后续遇到更小的height时，宽度+1 * height

		// 计算宽度，并计算面积
		area := heights[i] * width
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

// @lc code=end

