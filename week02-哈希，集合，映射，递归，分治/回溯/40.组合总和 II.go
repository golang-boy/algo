/*
 * @lc app=leetcode.cn id=40 lang=golang
 * @lcpr version=20003
 *
 * [40] 组合总和 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {

	/*
	 输入：数组，1-50， 目标值
	 输出：子集
	 限制：元素只能用一次
	*/
	var res = [][]int{}
	var chosen = []int{}
	var used = map[int]struct{}{}

	sort.Ints(candidates)

	// 拿到子集，和等于t
	r(target, candidates, &res, chosen, 0, used)

	return res
}

func key(a []int) string {
	return fmt.Sprintf("%v", a)
}

func r(t int, cs []int, res *[][]int, chosen []int, start int, used map[int]struct{}) {
	if t < 0 {
		return
	}

	if t == 0 {
		temp := make([]int, len(chosen))
		copy(temp, chosen)
		*res = append(*res, temp)
		return
	}

	var dedup = map[int]bool{}

	for i := start; i < len(cs); i++ {
		if t-cs[i] < 0 {
			break
		}

		if _, ok := used[i]; !ok {
			if dedup[cs[i]] {
				continue
			}
			dedup[cs[i]] = true

			chosen = append(chosen, cs[i])
			used[i] = struct{}{}
			r(t-cs[i], cs, res, chosen, i, used)
			chosen = chosen[:len(chosen)-1]
			delete(used, i)
		}
	}
}

/*
   组合题目
    1. 递归通过 t-cs[i], 进行，其他用于维护状态
	2. 重点是剪枝
	   1. 剪枝通过控制start来实现,第一个分支是全部的，其他的依次右移
	      通过控制选择空间，减少不必要的重复计算
	   2. 同一拨遍历操作时，后续相同元素通过dedup进行去重，前面做过，后面不用针对相同元素再做一遍

	时间： 1:2:20

*/

// @lc code=end

/*
// @lcpr case=start
// [10,1,2,7,6,1,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2,5,2,1,2]\n5\n
// @lcpr case=end

*/

