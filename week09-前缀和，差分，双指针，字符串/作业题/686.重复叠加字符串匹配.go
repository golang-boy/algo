/*
 * @lc app=leetcode.cn id=686 lang=golang
 * @lcpr version=20004
 *
 * [686] 重复叠加字符串匹配
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func repeatedStringMatch(a string, b string) int {

	/*
		a叠加多次后，b为a的子串
		输出叠加的次数

		要使条件满足:
		   1. 则a必为b的子串
		   2. 找到整个a在b中的位置
		   3. 往左往右找 abc,bcd,或者ab,cd,a,d
		   4. b像个环形数组
		   	   先定位a在b中的位置，如果找到,则计数
			   然后在环形数组中查找是否还有其他a(或者尾部的和头部的拼起来是否为a),如果是, 则+2
	*/

}

// @lc code=end

/*
// @lcpr case=start
// "abcd"\n"cdabcdab"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"aa"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"wxyz"\n
// @lcpr case=end

*/

