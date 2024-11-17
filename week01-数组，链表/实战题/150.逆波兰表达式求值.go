/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {

	/*
		关键信息：
		 1. 有效的运算符号: +, -, *, /,  都有两个操作数
		 2. 计算结果可以使用int类型存储
		 3. 表达式为后缀表达式, 操作数在前，操作符在后
		 4. 零向截断，如：13 / 5 = 2, 且不会包含除零的情况

		变化的量：

		  token需要循环处理
		    1. 遇到数字，则表示操作数，需要存下来
			2. 遇到符号，拿之前存的数进行计算，计算结果也存下来,之前的数需要清理掉
			3. 循环结束，返回最后一个数

		  两次循环,记录两个数,第三次循环，计算结果，并清理掉之前的数,保存结果
		  需要有个结构来存储数, 特点如下
		    1. 存储两个数, 数的顺序无关
			2. 计算时，需要能取出，并清理
			3. 计算完的结果，再存进去
			4. 最后再出来

		 两种方案：
		   1. 栈结构
		   2. 两个类似寄存器的变量, 一个状态位遍历

	*/

	s := []int{}

	for _, t := range tokens {

		switch t {
		case "+":
			// 出栈俩个
			if len(s) < 2 {
				return -1
			}
			num1 := s[len(s)-2]
			num2 := s[len(s)-1]

			r := num1 + num2

			s = s[:len(s)-2]
			s = append(s, r)

		case "-":
			if len(s) < 2 {
				return -1
			}
			num1 := s[len(s)-2]
			num2 := s[len(s)-1]
			s = s[:len(s)-2]

			r := num1 - num2
			s = append(s, r)
		case "*":
			if len(s) < 2 {
				return -1
			}
			num1 := s[len(s)-2]
			num2 := s[len(s)-1]
			s = s[:len(s)-2]

			r := num1 * num2
			s = append(s, r)
		case "/":
			if len(s) < 2 {
				return -1
			}
			num1 := s[len(s)-2]
			num2 := s[len(s)-1]
			s = s[:len(s)-2]

			r := num1 / num2
			s = append(s, r)
		default:
			// 默认都是数字, 不存在非数字的情况
			// 数字中有正数，有负数
			num, _ := strconv.Atoi(t)
			s = append(s, num)
		}
	}

	if len(s) != 1 {
		return -1
	}

	return s[0]
}

// @lc code=end

