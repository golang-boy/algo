/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
func calculate(s string) int {
	/*
		关键信息：
		  1. 运算符只有：+ - * / ( ), 每个有俩操作数
		  2. 操作数，都为正数，且在 32 位整数范围内
		  3. 有空格，需要过滤掉
		  4. 不同符号有不同的优先级, * / 优先级大于 + -, 总共俩级优先级

		思路：
		    1. 将中缀表达式，转为后缀表达式

			 情况一： 顺序计算

			    3*2+2

				3 2 * 2 +

			 情况二： 优先级高的先计算

			    3+2*2

				3 2 2 * +

			1. 为数字时，如tokens列表,
			2. 为运算符时,
			     栈空时,入栈;
			     栈不空时, 比较优先级
				     当前优先级高, 入栈
				     当前优先级低, 运算符栈出栈，写入tokens列表，当前符号入栈

			 情况三： 多个连续数字

			    33 + 2*22

			 情况四：有空格

			   3              + 2*2

			2. 再计算后缀表达式


		复杂度分析：
		  时间复杂度：O(n), 遍历一遍字符串
	*/
	opStack := make([]rune, 0)

	getLevel := func(op rune) int {
		if op == '*' || op == '/' {
			return 1
		}
		return 0
	}

	tokens := []string{}
	num := ""

	// 1. 转换为后缀表达式
	for _, c := range s {
		// 情况三: 识别数字
		if c >= '0' && c <= '9' {
			num += string(c)
			continue
		} else {
			// 遇到不是数字时，计算nums
			if num != "" {
				tokens = append(tokens, num)
				num = ""
			}
		}

		//  情况四: 过滤空格
		if c == ' ' {
			continue
		}

		// 计算优先级与栈顶优先级比较
		for len(opStack) != 0 && getLevel(c) <= getLevel(opStack[len(opStack)-1]) {
			lastOp := opStack[len(opStack)-1]
			opStack = opStack[:len(opStack)-1]
			tokens = append(tokens, string(lastOp))
		}
		opStack = append(opStack, c)
	}

	if num != "" {
		tokens = append(tokens, num)
	}

	for len(opStack) != 0 {
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		tokens = append(tokens, string(op))
	}

	fmt.Println(tokens)
	// 2. 计算后缀表达式
	r := evalRPN(tokens)
	return r
}

func evalRPN(tokens []string) int {
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
