/*
 * @Author: 刘慧东
 * @Date: 2024-11-19 09:51:55
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-19 10:02:00
 */
/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=20003
 *
 * [1] 两数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func twoSum(nums []int, target int) []int {

	/*
		  关键信息：
		   1. 输入一个序列，目标数
		   2. 输出索引
		   3. 每种输入只有一种输出
		   4. 每个输出只允许使用一次, 顺序无关
		   5. 序列中某俩数加起来和为目标数

		  分析：
		    1. 数组中找数，需要遍历
			2. 判断当前数是否满足条件
			      条件：target - cur 是否在数组中存在
			3. 满足条件则返回索引
			      索引：{target - cur,  cur}
	*/

	ans := []int{}

	query := map[int]int{} // 存储值与索引的映射

	for i := 0; i < len(nums); i++ {
		if j, ok := query[target-nums[i]]; ok {
			ans = append(ans, j, i)
			return ans
		}
		query[nums[i]] = i
	}

	return ans
}

// @lc code=end
