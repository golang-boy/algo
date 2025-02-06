/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {

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

			// 计算
			// 当前元素比前一个高或者相等时, 宽度累加
			bottom := top.height
			acc += top.width

			stack = stack[:len(stack)-1]
			// 需要看前一个的前一个情况,与当前比较，谁是上限
			if len(stack) == 0 {
				break
			}

			pretop := stack[len(stack)-1]
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

// @lc code=end

