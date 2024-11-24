/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {

	/*
	   排列问题
	*/

	var res = [][]int{}
	var chosen = make([]int, 0, len(nums))
	var used = make([]bool, len(nums))

	recure(nums, 0, &res, chosen, used)

	return res
}

func recure(nums []int, j int, res *[][]int, chosen []int, used []bool) {

	if j == len(nums) {
		temp := make([]int, len(chosen))
		copy(temp, chosen)
		*res = append(*res, temp)
		return
	}

	var uniq = map[int]bool{}

	for i := 0; i < len(nums); i++ {

		if !used[i] {

			if uniq[nums[i]] {
				continue
			}

			chosen = append(chosen, nums[i])
			used[i] = true
			recure(nums, j+1, res, chosen, used)
			chosen = (chosen)[:len(chosen)-1]
			used[i] = false

			uniq[nums[i]] = true
		}
	}
}

// @lc code=end

