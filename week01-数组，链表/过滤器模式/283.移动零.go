/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	/*

									关键信息：
									    1. 给定一个数组
									    2. 保持非零元素的顺序
										3. 零元素移到数组末尾
										4. 数组内操作，不使用额外空间

									变化的信息：
									    1. 数组的访问涉及到变量, 需要一个变量来记录遍历的位置
										2. 根据过滤条件，交换俩元素的位置, 需要变量来记录交换的位置

								     i
									[0,1,0,3,12]         // 如果要保证顺序性，只能从前往后遍历
									   j                 // 如果i==0 时，j不为0时，交换, i++, j++

				                       i
									[1,0,0,3,12]
									       j

				                         i
									[1,3,0,0,12]
									          j

				                            i
									[1,3,12,0,0]
									             j

						             i
									[0,0,0,3,12]    // 如果i==0,j==0,j++，i不能变,不发生交换
						               j

		                               i
									[1,3,12,0,0] // 不等于0时，不发生交换，i++,j++
									     j

	*/

	i := 0
	j := i + 1

	filter := func(i, j int) int {
		if nums[i] == 0 && nums[j] != 0 {
			return 0
		}
		if nums[i] == 0 && nums[j] == 0 {
			return 1
		}

		if nums[i] != 0 {
			return -1
		}
		return -1 // 其他情况下都往前走
	}

	for {

		if j >= len(nums) {
			break
		}

		if filter(i, j) == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if filter(i, j) == 1 {
			j++
		} else {
			i++
			j++
		}
	}
	return

}

// @lc code=end

