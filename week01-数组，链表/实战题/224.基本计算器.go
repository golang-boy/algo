/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
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
			return 2
		}
		if op == '+' || op == '-' {
			return 1
		}

		if op == '(' {
			return 0
		}
		return 0
	}

	tokens := []string{}
	num := ""
	needzero := true

	// 1. 转换为后缀表达式
	for _, c := range s {
		// 情况三: 识别数字
		if c >= '0' && c <= '9' {
			num += string(c)
			needzero = false
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
		if c == '(' {
			opStack = append(opStack, c)
			needzero = true
			continue
		}
		if c == ')' {
			// 遇到右括号时，需要找到左括号中间的所有符号，进入token
			for len(opStack) != 0 {
				lastOp := opStack[len(opStack)-1]
				if lastOp == '(' {
					opStack = opStack[:len(opStack)-1] // 出栈
					break
				}
				opStack = opStack[:len(opStack)-1] // 出栈
				tokens = append(tokens, string(lastOp))
			}
			needzero = false
			continue
		}

		if (c == '-' || c == '+') && needzero {
			tokens = append(tokens, "0") // 正负号前补零
		}
		// 计算优先级与栈顶优先级比较
		for len(opStack) != 0 && getLevel(c) <= getLevel(opStack[len(opStack)-1]) {
			lastOp := opStack[len(opStack)-1]
			tokens = append(tokens, string(lastOp))
			opStack = opStack[:len(opStack)-1]
			fmt.Println(string(lastOp))
		}
		opStack = append(opStack, c)
		needzero = true
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

/*

  操作结构:
    一个栈，用于存储运算符
    一个tokens结构，用于存储后缀表达式

  循环遍历字符串
    1. 遇到数字，写入tokens
	2. 遇到运算符，与栈顶比较优先级,
	   找到栈中所有大于等于当前运算符优先级的运算符，写入tokens, 并且出栈
	3. 当前运算符入栈

  最后一个数字写入到num后，没机会放入tokens中,
    因此循环完毕需要再写入一次, 或者一开始末尾补个空格

  最后将栈中所有运算符，写入tokens中


  针对括号的处理
      1. 先左括号入栈,
	  2. 当遍历到右括号时, 需要找到左括号中间的所有符号，进入tokens

  针对正负数的处理：
      通过补零的方式，将正负数转为正数，统一处理
	  补零的时机：
	    1. 正负号前，补零
		2. 确定补零的位置：
		    开始时
			正负号遍历完毕后, 下一个字符可能是正负号时
			左括号后面可能为正负号时
*/

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

