/*
 * @lc app=leetcode.cn id=56 lang=golang
 * @lcpr version=20004
 *
 * [56] 合并区间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func merge(intervals [][]int) [][]int {

	/*
		输入：区间集合
		输出：返回不重叠的区间集合

		重叠的判定：
		   1. 前一个下标小于等于后一个下标, 后一个下标大于等于前一个下标, 此时为重叠,可以合并
		   2. 合并后的新的区间，再与下一个进行合并
	*/

	ans := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	pre := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= pre[0] {
			if intervals[i][1] >= pre[1] {
				if intervals[i][0] <= pre[1] {
					pre[0] = pre[0]
					pre[1] = intervals[i][1]
				} else if intervals[i][0] > pre[1] {
					temp := make([]int, 2)
					copy(temp, pre)
					ans = append(ans, temp)
					pre[0] = intervals[i][0]
					pre[1] = intervals[i][1]
				}
			} else {
				pre[0] = pre[0]
				pre[1] = pre[1]
			}
		}
	}

	temp := make([]int, 2)
	copy(temp, pre)
	ans = append(ans, temp)

	return ans
}

/*
总结：
    1. 先排序,确定上限的范围，然后确定下限的范围, 确认合并的条件
	2. 然后将结果写入答案
	3. 最后一次时，由于没将结果写入，需要循环完毕后写入一次

*/

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[2,6],[8,10],[15,18]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,4],[4,5]]\n
// @lcpr case=end

*/

