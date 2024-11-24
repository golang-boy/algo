/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	/*
		  子集问题
		    逻辑处理：
			  1. 找到与nums拥有相同度的最短子集
			      ans = min(len(subset), ans)
			  2. 返回子集长度 len(subset)
			  3. 最短连续子数组 ==>  连续的
	*/

	var chosen = []int{}
	var ans int = math.MaxInt32

	d := getDegree(nums)
	getSubset(0, nums, &ans, chosen, d)
	return ans
}

func getSubset(i int, nums []int, ans *int, chosen []int, d int) {

	if i == len(nums) {
		if d == getDegree(chosen) {
			*ans = min(*ans, len(chosen))
		}
		return
	}

	getSubset(i+1, nums, ans, chosen, d)

	chosen = append(chosen, nums[i])
	getSubset(i+1, nums, ans, chosen, d)
	chosen = chosen[:len(chosen)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getDegree(nums []int) int {
	q := map[int]int{}

	for _, n := range nums {
		q[n]++
	}

	max := 0
	for _, v := range q {
		if v > max {
			max = v
		}
	}
	return max
}

/*

失败，未审清楚题目, 当做子集问题去处理
*/

// @lc code=end

