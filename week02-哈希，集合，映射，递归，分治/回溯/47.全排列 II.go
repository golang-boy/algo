/*
 * @Author: 刘慧东
 * @Date: 2024-11-25 15:13:35
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-25 15:28:58
 */
/*
 * @lc app=leetcode.cn id=47 lang=golang
 * @lcpr version=20003
 *
 * [47] 全排列 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func permuteUnique(nums []int) [][]int {

	var ans = [][]int{}
	var used = make([]bool, len(nums))
	var chosen = []int{}

	r(0, nums, &ans, used, chosen)
	return ans
}

func r(i int, nums []int, ans *[][]int, used []bool, chosen []int) {
	if i == len(nums) {
		temp := make([]int, len(chosen))
		copy(temp, chosen)
		*ans = append(*ans, temp)
		return
	}

	var dedup = make(map[int]bool, len(nums))
	for j := 0; j < len(nums); j++ {
		if !used[j] {
			if dedup[nums[j]] {
				continue
			}
			dedup[nums[j]] = true
			chosen = append(chosen, nums[j])
			used[j] = true
			r(i+1, nums, ans, used, chosen)
			chosen = chosen[:len(chosen)-1]
			used[j] = false
		}
	}
}

/*
   循环的作用：选第一个，或者选第二个，或者选第三个
      used作用：选过的不选，本层选过后，下一层就不用选了, 排列的关键
	  dedup作用：每次循环时，如果后面有相同的，就不用再选了，前面的已经处理过了
   递归的作用：选定第一个，选定二个，选定第三个
      到达末尾收集结果
*/

// @lc code=end

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

