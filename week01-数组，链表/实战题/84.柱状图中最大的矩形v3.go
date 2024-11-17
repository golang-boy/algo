/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {

	/*
		    关键信息：
			  输入： 表示高度的数组
			  输出： 最大矩形面积

			  变量：
			     宽度会累加变化
				 高度随着循环变化

			分析：

			   area :=  height * (width * n)

			   1. 需要循环遍历heigths数组
			   2. 从当前heights[i]的高度与后面的依次比较,直到小于当前height，此时，找到最大的宽度(向右扩展)
			          如果前一个的height > 当前height (向左扩)
			   3. 计算面积，并更新最大面积
			   4. 继续循环，从heights[i+1]开始

			   重点需要关注面积的变化, 面积由一组的height和width组成

			    右扩展时
				    后面的高度大于当前的高度, 宽度累加, 面积 = 当前高度 * 累加的宽度, 一直累加到不能再累加为止

			    左扩展时
				    需要和前一个进行比较, 如果高度大于当前高度，则加上前面的宽度，如果小于则停止

				遍历是从左往右，如何做检测做扩展呢？
				    因为是就近检测，因此使用栈保存前一个矩形的信息
					由于从左右到右遍历时，计算过一次当前高度最大面积，所以栈中只需要保存计算过的最大矩形信息即可


			单调栈：
			   1. 栈中保存的是递增的元素
			   2. 如果遇到比栈顶小的元素，则开始计算面积
			   3. 栈顶元素出栈，计算面积，并更新最大面积
			   4. 继续出栈，直到栈顶元素小于当前元素，或者栈为空
			   5. 将当前元素入栈
			   6. 重复步骤1-5，直到遍历结束

			存什么元素？
			   存储矩形的高度和宽度
	*/

	maxArea := 0

	stack := []struct {
		height int
		width  int
	}{}

	heights = append(heights, 0) // 最后一个元素得检测一次，加一个高度为0的元素, 对面积没有影响

	for i := 0; i < len(heights); i++ {

		acc := 0
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if top.height < heights[i] {
				break
			}
			acc += top.width
			// 出栈
			maxArea = max(maxArea, top.height*acc)
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, struct {
			height int
			width  int
		}{
			height: heights[i],
			width:  acc + 1,
		})

	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*

 总结：

     单调栈的时间复杂度为o(n)

     单调栈的思路：
	    1. 栈中保存的是递增的元素
		2. 如果遇到比栈顶小的元素，如果当前元素破坏了递增性(这是一个简化的过程，简化为前一个元素, 和当前元素)
		     计算前一个柱子的扩展面积，更新最大面积, 并将不符合单调递增的柱子，都出栈
			 直到满足单调性，此时以当前元素高度， 累积宽度为基准，入栈
		3. 重复步骤2，直到遍历结束

	 应用场景
		单调栈是一种高效的解决方案，在许多与区间和边界相关的问题中表现出色。
		广泛用于解决与栈相关的各种问题，特别是在需要查找某个元素的“下一个”或“之前”的元素时
		单调栈的每个元素在入栈和出栈时最多只会被处理一次

*/

// @lc code=end

