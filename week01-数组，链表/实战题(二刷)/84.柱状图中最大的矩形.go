/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {

	/*
		明确输入输出：
			输入非负整数数组
			输出矩形最大面积

		左侧比相邻右侧高时，不延伸，否则就延伸计算面积, 宽度加1
		到了2和3这里，除向右延伸外，还需要向左延伸

		什么时候不能延伸？
			遇到比自己低的, 则停止延伸

		本题需要考虑前面相邻的元素, 遇到比自己小时，停止延伸
			此时需要计算面积, 同时向左延伸

		循环的特点是从左往右循环, 如果前一个元素比较自己小,需要考虑右延伸,此时宽度+1,
	*/

	s := []struct {
		height int
		width  int
	}{}
	ans := 0
	heights = append(heights, 0)

	for i := 0; i < len(heights); i++ {
		acc := 0
		// 该循环为左延伸, 需要累加宽度
		for len(s) != 0 {
			top := s[len(s)-1]
			if heights[i] > top.height {
				// 此时右延伸
				break
			}

			acc += top.width
			ans = max(ans, top.height*acc)
			s = s[:len(s)-1]
		}

		// 右延伸
		s = append(s, struct {
			height int
			width  int
		}{
			height: heights[i],
			width:  acc + 1,
		})
	}

	return ans

}

// @lc code=end

