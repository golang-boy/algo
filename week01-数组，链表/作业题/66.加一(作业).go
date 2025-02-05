/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	/*

		  关键信息:

		     1. 输入一个非空非负整数序列
			 2. 每个位置存储单个数字
			 3. 最高位数字存放在序列的起始位置
			 4. 除0外，不会有零开头的数
			 5. 算法输出为在序列对应的数字上加一

		  约束：
		     1. 加一后，存在进位
			 2. 进位后，前一个数字也加一
			 3. 直到最前面的数字加一后，没有进位
			 4. 数组中进行上述操作，最前面的数字进位后，数组要整体向后移动
			    go中因为是切片，重新赋值即可
	*/

	// 从后向前遍历
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i] + 1

		div := num / 10
		digits[i] = num % 10

		if div == 0 {
			// 没有进位
			return digits
		}

		if i == 0 { //最前面的
			if div != 0 {
				digits = append([]int{1}, digits...)
				return digits
			}
		}
	}

	return digits

}

// @lc code=end

