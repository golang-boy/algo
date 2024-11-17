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
		   计算时，如果是数，则保存起来, 如果是符号，同时读取下一个操作数，
		     1. 得判断下一个符号和当前符号的优先级,
			        如果优先级相同，还得继续判断下下个符号的优先级, 先计算优先级高的

					先进来的数，可能会后被计算    => 一个存储数的栈
					先进入的运算符,可能会后被计算 => 一个存储运算符的栈

			 2. 数栈什么时候入栈? 什么时候出栈？运算符栈什么时候入栈? 什么时候出栈?
			    1. 当前符号为高优先级时，出数栈, 读取下一个操作数, 进行运算, 结果入数栈
				2. 当前符号为低优先级时，入数栈, 符号入运算符栈
				3. 当循环到字符末尾时，需要处理运算符栈中剩余的运算符, 所有运算符处理完毕，则计算完成
				4. 表达式总是有效，则不需要判断数与符号不匹配的情况

		 复杂度分析：
		  时间复杂度：O(n), 遍历一遍字符串
	*/
	numStack := make([]int, 0)
	opStack := make([]byte, 0)

	num := []byte{}
	skip := false

	getLevel := func(op byte) int {
		if op == '*' || op == '/' {
			return 1
		}
		return 0
	}

	for i, c := range s {
		if skip {
			skip = false
			continue
		}
		switch c {
		case '+', '-':
			currlevel := getLevel(c)
			if len(opStack) != 0 && getLevel(opStack[len(opStack)-1]) >= currlevel {
				// 栈顶出栈, 计算，结果入数栈
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]

				num1 := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]

				// 操作数怎么获取？操作数还是下一个字符,需要预取，否则无法获取
				// 计算完毕，结果入数栈，循环完毕后，需要跳过当前字符,
				// 因此，需要有个标识判断是否跳过当前字符

				if i+1 > len(s)-1 {
					// 超出范围，并且表达式不完整
					break
				}
				num2 := s[i+1] // 需要考虑多个数字的情况
				skip = true

				r := calc(num1, num2, op)
				numStack = append(numStack, r)

				// 当前操作符需要入栈, 什么时候计算呢？
				opStack = append(opStack, byte(c))
			} else {
				// 如果当前符号比栈顶优先级高，针对两级的情况,此时需要进行计算
				// opStack = append(opStack, byte(c))
			}
		case '*', '/':

		default:
			if c >= '0' && c <= '9' {
				num = append(num, c)
			}
			if c == ' ' {
				if len(num) != 0 {
					n, _ := strconv.Atoi(string(num))
					numStack = append(numStack, n)
					num = []byte{}
				}
			}
		}
	}
}

/*
总结：
    1. 不能像上面那样写程序，需要处理的细节太多，容易出错
	2. 一个循环中不可能把所有逻辑和细节都处理了
	3. 需要先进行token的识别的，即词法分析, 然后进行语法分析, 构造出表达式树, 最后进行计算

	而这些已经超出了算法题的范畴，需要从编译的角度进行程序处理
*/

// @lc code=end

