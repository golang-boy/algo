/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	/*

		   关键信息
		   1. 宽度为1
		   2. 序列为高度
		   3. 输出为接住的雨水量
		        1. 左右边界中低的高度-当前的高度

				    ans := left.height - heights[i]
					if ans > 0 {
					    左边可以接住
					} else {
					    左边不可以接住
					}

				2. 累加起来就是接住的雨水量


	*/

	stack := []struct {
		height int
		width  int
	}{}
	ans := 0

	for i := 0; i < len(height); i++ {

		acc := 0
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top.height > height[i] {
				break
			}

			// 更新答案，并出栈
			bottom := top.height // 前一个的高度
			acc += top.width
			stack = stack[:len(stack)-1] // 出栈

			if len(stack) == 0 {
				continue
			}
			pretop := stack[len(stack)-1] // 前一个的前一个高度

			// 前后俩个哪个低，取哪个的高度
			up := min(pretop.height, height[i])
			ans += (up - bottom) * acc
		}

		stack = append(stack, struct {
			height int
			width  int
		}{
			height: height[i],
			width:  1 + acc,
		})
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
 单调递减栈:
	1. 当前的高度小于栈顶的高度，入栈
	2. 大于时，栈顶出栈，计算下限，宽度 (这是一个合并简化的过程, 简化为当前栈顶元素，当前元素, 栈顶第二个元素)
	3. 查看出栈后，当前栈顶的元素height, 取height和当前元素height的最小值, 拿到上限
	4. 计算雨水量  (上限 - 下限) * 宽度
	5. 计算完毕后，将当前高度，累积的宽度信息入栈, 下次循环时，就不用计算
*/

// @lc code=end

