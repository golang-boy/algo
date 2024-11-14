/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(ss string) bool {
	/*

		  关键信息：
		    1. 给定的是字符串  ==》 意味着需要遍历字符串中的字符
			2. 字符串中只包含 ()[]{}
			3. 各种括号需要按顺序闭合匹配

		  变化的量:

		    1. 遍历时，右括号时，需要和前面的字符进行匹配比较
			2. 左括号不需要进行匹配比较
			3. 前面对应的字符，需要先存起来
			    1. 用到时再取出来比较
				2. 要保持顺序
				3. 先进入的后出，先比较最近的
				 ==》 栈： 能保持顺序，且先入后出
	*/

	s := []rune{}

	for _, ch := range ss {
		switch ch {
		case ')':
			// 注意：比较前，先判断栈是否为空，如果为空，则直接返回false
			if len(s) == 0 {
				return false
			}
			//  栈顶比较, 如果匹配则出栈丢弃，继续下一次，不匹配则返回false
			if s[len(s)-1] == '(' {
				s = s[:len(s)-1] // 出栈
			} else {
				return false
			}
		case '}':
			// 注意：比较前，先判断栈是否为空，如果为空，则直接返回false
			if len(s) == 0 {
				return false
			}
			//  栈顶比较, 如果匹配则出栈丢弃，继续下一次，不匹配则返回false
			if s[len(s)-1] == '{' {
				s = s[:len(s)-1] // 出栈
			} else {
				return false
			}
		case ']':
			// 注意：比较前，先判断栈是否为空，如果为空，则直接返回false
			if len(s) == 0 {
				return false
			}

			// 栈顶比较, 如果匹配则出栈丢弃，继续下一次，不匹配则返回false
			if s[len(s)-1] == '[' {
				s = s[:len(s)-1] // 出栈
			} else {
				return false
			}
		default:
			// 注意：需要确定可以入栈的字符,注意条件检测
			//  入栈
			s = append(s, ch)
		}
	}

	//  判断栈里空不空，如果空则返回true，否则返回false
	if len(s) == 0 {
		return true
	}
	return false
}

// @lc code=end

