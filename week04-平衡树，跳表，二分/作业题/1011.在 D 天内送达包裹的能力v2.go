/*
 * @lc app=leetcode.cn id=1011 lang=golang
 * @lcpr version=20004
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func shipWithinDays(weights []int, days int) int {

	var left = 1
	var right = 0

	left = max(weights)
	right = sum(weights)

	for left < right {
		mid := (left + right) / 2

		if validate(mid, weights, days) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func validate(mid int, weights []int, days int) bool {

	dayCnt := 1
	accWeight := 0

	// p := []string{}

	for _, w := range weights {
		if w > mid {
			return false
		} else if accWeight+w > mid {
			dayCnt++
			// fmt.Println(mid, dayCnt, strings.Join(p, ","))
			accWeight = w
			// p = []string{}
		} else {
			accWeight += w
			// p = append(p, fmt.Sprint(w))
		}
	}

	return dayCnt <= days
}

func sum(weights []int) int {
	s := 0
	for _, w := range weights {
		s += w
	}
	return s
}

func max(weights []int) int {
	ans := 0
	for _, w := range weights {
		if ans < w {
			ans = w
		}
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10]\n5\n
// @lcpr case=end

// @lcpr case=start
// [3,2,2,4,1,4]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,1]\n4\n
// @lcpr case=end

*/

