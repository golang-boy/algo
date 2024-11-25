/*
 * @Author: 刘慧东
 * @Date: 2024-11-25 15:29:39
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-25 18:49:43
 */
/*
 * @lc app=leetcode.cn id=90 lang=golang
 * @lcpr version=20003
 *
 * [90] 子集 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func subsetsWithDup(nums []int) [][]int {

	/*

	 */

	var res [][]int
	var chosen []int
	var used = map[string]bool{}

	sort.Ints(nums)

	r(0, nums, &res, chosen, used)
	return res
}

func key(a []int) string {
	return fmt.Sprintf("%v", a)
}

func r(i int, nums []int, res *[][]int, chosen []int, used map[string]bool) {

	if i == len(nums) {
		if used[key(chosen)] {
			return
		}
		used[key(chosen)] = true
		temp := make([]int, len(chosen))
		copy(temp, chosen)
		*res = append(*res, temp)
		return
	}

	chosen = append(chosen, nums[i])
	r(i+1, nums, res, chosen, used)
	chosen = chosen[:len(chosen)-1]
	r(i+1, nums, res, chosen, used)

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

