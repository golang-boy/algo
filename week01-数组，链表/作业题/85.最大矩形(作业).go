/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	/*
	   先算第一行的柱状图
	   然后再算第二行的柱状图，更新高度
	*/

	heights := make([]int, len(matrix[0]))
	maxArea := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		fmt.Println(heights)
		area := largestRectangleArea(heights)
		maxArea = max(maxArea, area)
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)

	stack := []struct {
		height int
		width  int
	}{}
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		acc := 0
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if top.height < heights[i] {
				break
			}

			stack = stack[:len(stack)-1]
			acc += top.width
			maxArea = max(maxArea, top.height*acc)
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

// @lc code=end

