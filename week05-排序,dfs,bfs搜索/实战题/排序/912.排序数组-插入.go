/*
 * @lc app=leetcode.cn id=912 lang=golang
 * @lcpr version=20004
 *
 * [912] 排序数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func sortArray(nums []int) []int {

	/*
		排序算法分为二大类：
		   1. 基于比较的排序，时间复杂度下界为O(nlogn)
		      * 插入排序
		          简单插入排序，希尔排序
				  每次从前到后，考虑每个未排序的数据，在已排序的序列中找位置插入
				    插到前一个元素的前面或者后面

			稳定与不稳定，排序前后，相对次序保持不变的算法，成为稳定

	*/

	for i := 0; i < len(nums); i++ {
		pre := i - 1   // 当前元素的前一个元素
		cur := nums[i] // 当前元素, 使用cur临时保存

		for pre >= 0 && nums[pre] > cur {

			// 如果大于当前，则每次循环将pre的对应元素后移一位
			nums[pre+1] = nums[pre] // 将pre元素放到下一个（第一次时，相当于放到当前位置）
			pre -= 1                // pre往前移动,准备下一次移动
		}
		nums[pre+1] = cur //移动完毕后(循环终止后)，将当前cur放入最后一个位置
	}

	return nums
}

// @lc code=end

/*
// @lcpr case=start
// [5,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,1,2,0,0]\n
// @lcpr case=end

*/

