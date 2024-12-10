/*
 * @Author: 刘慧东
 * @Date: 2024-12-10 15:27:52
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-12-10 18:28:30
 */
/*
 * @lc app=leetcode.cn id=1482 lang=golang
 * @lcpr version=20004
 *
 * [1482] 制作 m 束花所需的最少天数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minDays(bloomDay []int, m int, k int) int {
	/*
		   输入：
		      bloomDay为一个数组，每个元素表示i朵花盛开需要的天数
			  m为要制作几束花
			  k为每束花需要k朵花，相邻的花朵才能用于制作花束

		   输出：摘m束花需要等待n天，找n最小的, 找不到返-1

		   思路：

		      每束花，需要k朵花，这k朵花的盛开天数累加求和，看这个和是否是最小

			  二分法
			     左右边界：
				    left ,right对应天数的边界
					     m*k总共需要这么多朵花，
						   每束1朵，最大天数就是最大值, 此为right
						           找最天数为最小值，此为left

						 left := min(bloomDay)   做1束花，最小天数
						 right := max(bloomDay)  做len(bloomDay)束花，最大天数

					mid代表需要验证的条件,如何验证mid符合要求
	*/

	if len(bloomDay) < m*k {
		return -1
	}

	left := min(bloomDay)
	right := max(bloomDay)

	for left < right {
		mid := (left + right) / 2
		if validate(mid, bloomDay, m, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func max(a []int) int {
	ans := 0
	for _, n := range a {
		if n >= ans {
			ans = n
		}
	}

	return ans
}

func min(a []int) int {
	ans := math.MaxInt
	for _, n := range a {
		if n <= ans {
			ans = n
		}
	}
	return ans
}

func validate(mid int, bloomDay []int, m, k int) bool {

	sum := 0
	cnt := 0

	for _, d := range bloomDay {
		if d <= mid {
			cnt++
			if cnt == k {
				sum++
				cnt = 0
			}
		} else {
			cnt = 0
		}
	}

	return sum >= m

}

// @lc code=end

/*
// @lcpr case=start
// [1,10,3,10,2]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,10,3,10,2]\n3\n2\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,12,7,7]\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// [1000000000,1000000000]\n1\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,10,2,9,3,8,4,7,5,6]\n4\n2\n
// @lcpr case=end

*/

