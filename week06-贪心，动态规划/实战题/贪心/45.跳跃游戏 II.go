/*
 * @lc app=leetcode.cn id=45 lang=golang
 * @lcpr version=20004
 *
 * [45] 跳跃游戏 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func jump(nums []int) int {

	/*
	   输入：nums[i] 从i向前跳的最大长度, 每一跳，可以在最大长度内的任一位置
	   输出：跳到最末尾的最小次数, （约束为必定能调大）


	   每次跳跃时，选最大的跳，
	   如果超出，则说明能跳到
	   如果不够，则试着减一,回到上一层

	   如果 nums[i+1] 比nums[i]更大, 表示nums[i+1]跳的更远，因此需要跳到nums[i+1], 否则跳到nums[i]

	*/
	ans := 0

	for i := 0; i < len(nums)-1; {

		// 当前为i, reached为当前能跳的最远的索引
		reached := i + nums[i]

		// 跳到了末尾，则结束
		if reached >= len(nums)-1 {
			return ans + 1
		}

		maxReached := reached
		maxIndex := i

		//如果到不了，则在该[i+1, reached]范围找一个跳的最远的(跳得索引最大的)
		for j := i + 1; j <= reached; j++ {
			if j+nums[j] > maxReached { // 如果比当前的可达索引还大，则更新
				maxReached = j + nums[j]
				maxIndex = j
			}
		}
		i = maxIndex
		ans++
	}

	return ans
}

/*

总结：

    贪心，找最远的，如果当前不能确定策略，进行决策扩展，多看一步


*/

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end


// @lcpr case=start
// [2,3,0,1,4]\n
// @lcpr case=end

*/

