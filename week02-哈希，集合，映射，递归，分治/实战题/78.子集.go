/*
 * @lc app=leetcode.cn id=78 lang=golang
 * @lcpr version=20003
 *
 * [78] 子集
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func subsets(nums []int) [][]int {
	/*

		  关键信息：
		    1. 数组，元素不同
			2. 返回 数组，数组里包含所有子集构成的数组
			3. 子集不能重复

		  解题思路：
		    1. 选不选的问题
			  从0开始，选下标0或不选下标0
			2. 递归到最底下
	*/

	/*
	   递归的关键:
	     1. 定义需要递归的问题
	     2. 确定递归边界
	     3. 保护与还原现场

	     递归相当于有个隐藏的栈，将扫描未扫描的不停的放入栈中，直到最底下，然后不停的出栈
	        因此，需要先确定最底层的逻辑, 在考虑返回后怎么处理
	*/

	var chosen = []int{}
	var res = [][]int{}

	fmt.Printf("ptr %d  %p\n", 0, &chosen)

	getSubset(nums, 0, &res, &chosen)

	return res
}

// i表示选择哪一个
// 输入参数必须是指针，否则入参的值不会改变，（go入参都是值传递）
func getSubset(nums []int, i int, res *[][]int, chosen *[]int) {
	// 索引i累积到最大时，将结果放入结果集

	// 为什么要等于len(nums)而不是len(nums)-1?
	//   因为最后一个元素，也要经历选和不选的过程
	if i == len(nums) {
		temp := make([]int, len(*chosen))
		copy(temp, *chosen)
		// temp := *chosen  // 不行，切片是引用类型，指向的实际数据还是一个，修改一个
		*res = append(*res, temp) // 直接将*chosen append到res中也不行，chosen变化时，res对应的位置也会变化

		// 必须要copy一份，否则res中的值会随着chosen的变化而变化
		// fmt.Printf("res %d  %v\n", i, *res)
		return
	}

	// fmt.Printf("unchosen %d  %v\n", nums[i], *chosen)
	getSubset(nums, i+1, res, chosen)  // 不选下标i,
	*chosen = append(*chosen, nums[i]) // 选下标i

	// fmt.Printf("chosen %d  %v\n", nums[i], *chosen)
	getSubset(nums, i+1, res, chosen)    // 选下标i后，继续后面的
	*chosen = (*chosen)[:len(*chosen)-1] // 还原现场
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

