/*
 * @lc app=leetcode.cn id=72 lang=golang
 * @lcpr version=20004
 *
 * [72] 编辑距离
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minDistance(word1 string, word2 string) int {

	/*
		输入俩个单词
		输出从word1到word2的最小操作数

		word1和word2前面补空格
		f[i][j] 在从word1的子串[1,i]个位置时，变换为从word2的子串[1,j]个位置的最小操作数
	*/

	word1 = " " + word1
	word2 = " " + word2

	n := len(word1)
	m := len(word2)

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, m)
		for j := range f[i] {
			f[i][j] = -1000
		}
	}

	for i := 0; i < n; i++ {
		// 从子串[1,i]变为空串, 需要i次
		f[i][0] = i
	}

	for i := 0; i < m; i++ {
		// 从空串变为子串[1,i], 需要i次
		f[0][i] = i
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {

			if word1[i] == word2[j] {
				f[i][j] = min(min(f[i][j-1]+1, f[i-1][j]+1), f[i-1][j-1]) // 相等时，无操作
			} else {
				f[i][j] = min(min(f[i][j-1]+1, f[i-1][j]+1), f[i-1][j-1]+1) // 不等时替换，操作+1
			}
		}
	}

	return f[n-1][m-1]
}

/*
总结：

	f[i][j-1] + 1 表示从[1,i] 变为[1,j-1]字符，然后在第j个位置加一个字符

	   因此,增加操作对应的递推关系为
	   f[i][j] = f[i][j-1] + 1

	f[i-1][j]+1) 表示从[1,i-1]变为[1,j], 将前i-1个字符变化到[1,j], 把第i个字符给删除不要

	   对应删除操作的递推关系
	   f[i][j] = f[i-1][j] + 1

	f[i-1][j-1] 将前i-1个字符变为前j-1个字符，最后这一个字符如果相等就不用任何操作，如果不等则操作需要加1,进行替换

	   对应替换或者不变的操作

	   f[i][j] = f[i-1][j-1] (+ 1)


	结果取值 f[n-1][m-1]

	递推关系需要有初始值，否则不能递推起来，本题中递推关系就是从[1,i]到[0,0]的操作，需要删除i次, 以及[0,0]到[1,j]需要j次增加操作


	状态方程中三个操作不易想到，需要费一些时间，是难点；另一个难点是f[i][j]的定义是难点，这两点搞不定，这题做不出

*/

// @lc code=end

/*
// @lcpr case=start
// "horse"\n"ros"\n
// @lcpr case=end

// @lcpr case=start
// "intention"\n"execution"\n
// @lcpr case=end

*/

