/*
 * @Author: 刘慧东
 * @Date: 2024-11-25 10:47:39
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-25 13:33:11
 */
/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=20003
 *
 * [39] 组合总和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	/*
			   输入： 数组，无重复元素，  目标数, 整型, 数都>=2
			   输出：二维数组
			   逻辑：数组中元素通过组合,计算和,等于目标值, 拿到这样的组合

			   思路：

			        t-cadidates[i]  == 0  return

		            t-cadidates[i]  != 0   t = t-cadidates[i]  递归
	*/

	res := [][]int{}
	chosen := []int{}

	sort.Ints(candidates)

	r(target, candidates, &res, chosen, 0)
	return res
}

func r(t int, cs []int, res *[][]int, chosen []int, start int) {
	if t < 0 {
		return
	}
	if t == 0 {
		temp := make([]int, len(chosen))
		copy(temp, chosen)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(cs); i++ {
		t1 := t - cs[i]
		if t1 < 0 {
			break
		}

		chosen = append(chosen, cs[i]) // 减去cs[i]时
		r(t1, cs, res, chosen, i)
		chosen = chosen[:len(chosen)-1]
	}
}

/*
总结：
   1. 找到所有组合
   2. 找的过程中进行剪枝
*/

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/

