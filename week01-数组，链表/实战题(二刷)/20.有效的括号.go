/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(ss string) bool {
	st := []byte{}

	for i := range ss {
		switch ss[i] {
		case ']':
			if len(st) == 0 {
				return false
			}
			ch := st[len(st)-1]
			st = st[:len(st)-1]
			if ch != '[' {
				return false
			}
		case '}':
			if len(st) == 0 {
				return false
			}
			ch := st[len(st)-1]
			st = st[:len(st)-1]
			if ch != '{' {
				return false
			}
		case ')':
			if len(st) == 0 {
				return false
			}
			ch := st[len(st)-1]
			st = st[:len(st)-1]
			if ch != '(' {
				return false
			}
		default:
			st = append(st, ss[i])
		}
	}

	if len(st) != 0 {
		return false
	}

	return true

}

// @lc code=end

