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


	*/

	maxArea := 0

	stack := []struct {
		height int
		width  int
	}{}

	heights = append(heights, 0) // 最后一个元素得检测一次，加一个高度为0的元素, 对面积没有影响

	for i := 0; i < len(heights); i++ {
		// 入栈

		// 当前高度小于后一个高度, 宽度累加
		// 找到当前块的最大的宽度
		accWidth := 1
		for j := i + 1; j < len(heights); j++ {
			if heights[i] <= heights[j] { // 向右找
				accWidth++
			} else {
				break // 需要向左找
			}
		}

		// for len(stack) > 0 { // 向左找找，找到前面累加的宽度
		// 	top := stack[len(stack)-1]
		// 	if top.height < heights[i] {
		// 		break
		// 	}
		// 	// 前面的高度高，累加宽度
		// 	stack = stack[:len(stack)-1]
		// 	accWidth += top.width // 重叠累加了宽度 !!!!!
		// }
		// 上述方式累加了重叠的宽度

		for k := len(stack) - 1; k >= 0; k-- { // 向左找找，找到前面累加的宽度
			top := stack[k]
			if top.height < heights[i] {
				break
			}
			// 前面的高度高，累加宽度
			accWidth += top.width
		}

		// 上述方式, 时间复杂度不能达标, 栈只会增加，不会减少

		stack = append(stack, struct {
			height int
			width  int
		}{
			heights[i],
			1,
		})
		// 更新最大面积
		maxArea = max(maxArea, heights[i]*accWidth)
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
   通过上述方式，先遍历右边，再遍历左边
     最外层的循环n次，内层需要向左向右遍历循环n次，因此时间复杂度是O(n^2)
*/

// @lc code=end

