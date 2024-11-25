/*
 * @lc app=leetcode.cn id=46 lang=golang
 * @lcpr version=20003
 *
 * [46] 全排列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func permute(nums []int) [][]int {
	/*
	  选择i,递归
	  不选择i,递归

	  递归的退出条件

	*/

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

	for j := 0; j < len(nums); j++ {
		if !used[j] {
			chosen = append(chosen, nums[j])
			used[j] = true
			r(i+1, nums, ans, used, chosen)
			chosen = chosen[:len(chosen)-1]
			used[j] = false
		}
	}
}

/*

   全排列：

       这种递归，相当于深度遍历，确定第一个元素后，依次确定第二个元素，直到最底层，输出结果
	   每次递归，表示每次判断下一个元素

	   循环的作用：选第一个，或者选第二个，或者选第三个
	   递归的作用：选定第一个，选定二个，选定第三个

*/

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

