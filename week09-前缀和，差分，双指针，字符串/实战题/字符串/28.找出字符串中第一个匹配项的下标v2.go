/*
 * @lc app=leetcode.cn id=28 lang=golang
 * @lcpr version=20004
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func strStr(haystack string, needle string) int {
	/*
		rabin-karp基于hash的高效字符串搜索算法

		定长n字符串s，长度m模式串t, 求t在s中出现

		计算s的每个长度为m的子串的hash值, 宽度为m的滑动窗口滑过s

		检测与t的hash值是否相等


		输入原始字符串，和模式子串
		输出模式子串在原字符串中的第一个位置索引

		朴素的思路:

			循环遍历原字符串的每个字符，在子串中依次匹配

			for i:=0;i<len(haystack);i++{
				// [l,r], 内部循环依次扫描比较
			}
	*/

	n := len(haystack)
	m := len(needle)

	b := uint64(131)

	ans := -1

	H := make([]uint64, n+1)

	// H表示每个位置字符对应的hash值
	//  0-26，总共27个数字代表26个字母+一个零位
	//    b表示进制数，字符串中越往后，位数越低
	//     前面的高位，需要*b,然后加上低位的数字
	//     如 123560
	//        1
	//        12
	//        123
	//        1235
	//        12356
	//
	for i := 1; i <= n; i++ {
		// p = 1e9+7
		// a = 1, b = 2, z = 26
		H[i] = H[i-1]*b + uint64(haystack[i-1]-'a'+1)
		// H[i] = (H[i-1]*b + uint64(haystack[i-1]-'a'+1) % p
	}

	Hneedle := uint64(0)
	powBM := uint64(1)

	// 从0开始，计算模式串的hash
	for _, ch := range needle {
		Hneedle = Hneedle*b + uint64(ch-'a'+1)
		// Hneedle = (Hneedle*b + uint64(ch-'a'+1))%p
		// 循环几次就是几位，每位乘以进制数
		powBM = powBM * b % p
	}

	// 循环比较
	for l := 1; l <= n-m+1; l++ {
		r := l + m - 1

		// 当前数 - 最高位*进制的多次位数平方, 等于[l,r]区间的hash值
		if H[r]-H[l-1]*powBM == Hneedle {
			// if ((H[r]-H[l-1]*powBM)%p+p)%p == Hneedle {
			ans = l - 1
			break
		}
	}

	// 利用自动溢出方式,不进行取模运算, 溢出后会将高位去掉,保留地位的，相当于取模
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// "sadbutsad"\n"sad"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n"leeto"\n
// @lcpr case=end

*/

