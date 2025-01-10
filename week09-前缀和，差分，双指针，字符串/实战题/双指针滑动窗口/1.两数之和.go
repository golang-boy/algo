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

	n := len(nums)

	j := n - 1

	newNums := MakeNums(nums)

	sort.Slice(newNums, func(i, j int) bool {
		return newNums[i].v < newNums[j].v
	})

	for i := 0; i < n; i++ {

		for ; i < j && newNums[i].v+newNums[j].v > target; j-- {
		}

		if i < j && newNums[i].v+newNums[j].v == target {
			return []int{newNums[i].i, newNums[j].i}
		}
	}
	return []int{}
}

type Num struct {
	v int
	i int
}

func MakeNums(nums []int) []Num {
	res := []Num{}
	for i, v := range nums {
		res = append(res, Num{v: v, i: i})
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/

