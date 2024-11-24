/*
 * @lc app=leetcode.cn id=344 lang=golang
 * @lcpr version=20003
 *
 * [344] 反转字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reverseString(s []byte) {
	r(0, &s)
}

func r(i int, s *[]byte) {
	n := len(*s)

	if i >= n/2 {
		return
	}
	r(i+1, s)
	(*s)[n-i-1], (*s)[i] = (*s)[i], (*s)[n-i-1]
	return
}

/*
    总结：
	   使用递归解题
	     1. 先入栈，到达中间
		 2. 返回时，通过计算，得到另一半的位置，进行位置交换

	   时间消耗25分钟

*/

// @lc code=end

/*
// @lcpr case=start
// ["h","e","l","l","o"]\n
// @lcpr case=end

// @lcpr case=start
// ["H","a","n","n","a","h"]\n
// @lcpr case=end

*/

