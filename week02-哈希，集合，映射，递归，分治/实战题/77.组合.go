/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=20003
 *
 * [77] 组合
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func combine(n int, k int) [][]int {

	/*
	  关键信息：
	    1. 1-n,  n个数
	    2. n个数中，选k个数
	    3. 返回组合，以为着不重复，顺序无关, 如[1,2]和[2,1]是一样的

	  分析:
	    1. 构造序列
	    2. 取第i个数
	    3. 这怎么递归呀?
	*/

	ans := [][]int{}
	chosen := []int{}

	recure(1, n, k, &ans, &chosen)
	return ans
}

func recure(i, n, k int, ans *[][]int, chosen *[]int) {
	if i == n+1 { // i从1开始，最后一个数也要参与递归
		if len(*chosen) == k { // 等于k个，满足条件
			temp := make([]int, k)
			// temp := []int{}   // 这种方式不行，copy时，不会扩容，也不会报错，结果不对
			copy(temp, *chosen)
			*ans = append(*ans, temp)
		}
		return
	}

	recure(i+1, n, k, ans, chosen) // 不选i, chosen中没i
	*chosen = append(*chosen, i)
	recure(i+1, n, k, ans, chosen) // 选i, chosen中有i
	*chosen = (*chosen)[:len(*chosen)-1]
}

// @lc code=end

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/

