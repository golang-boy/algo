/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {

	/*
		关键信息：
		 1. 递增排列数组, 给定一个数组
		 2. 有重复项,去重
		 3. 返回去重后的数组长度
		 4. 原地操作 ==> 意味着交换

		思路：
		  1. 一个数组，需要遍历，根据过滤条件，将符合条件的元素进行计数, 最后返回该计数
		  2. 变化的信息是什么？
		     * 因为要遍历，所以需要一个变量来记录遍历的位置，即i
			 * 最后返回结果是个数，所以需要一个变量来记录个数，即sum
			 * 去重是个比较操作，因此需要一个变量来记录比较的位置，即j
		  3. 什么时候结束遍历？
		     * 当j遍历到最后一个元素时，结束遍历

		 i
		[1,1,2]
		   j

		 i
		[1,1,2]
		     j

		   i
		[1,2,2]
		       j

	*/

	i := 0
	j := i + 1
	sum := 1 // 初始值为1，因为第一个元素一定不会重复

	filter := func(i, j int) bool {
		if nums[i] == nums[j] { // 相等时，i不变，j++, 不发生赋值操作
			return false
		}

		// if nums[i] < nums[j] { // 小于时，i++, 将j赋值给i, 发生赋值操作, j++指向下一个
		// 	return true
		// }

		// if nums[i] > nums[j] { // 大于时，i++，赋值，j++
		// 	return true
		// }

		return true
	}

	for {
		if j >= len(nums) { // 退出条件, 最后一个元素
			break
		}

		if filter(i, j) {
			i++
			nums[i] = nums[j]
			j++
			sum++
		} else {
			j++
		}
	}

	return sum

}

// @lc code=end

