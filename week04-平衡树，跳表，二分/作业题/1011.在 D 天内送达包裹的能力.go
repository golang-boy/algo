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
	/*
		第 i 个包裹的重量为 weights[i]
		1. 按给出重量（weights）的顺序往传送带上装载包裹
		2. 我们装载的重量不会超过船的最大运载重量
		3. 在 days 天内将传送带上的所有包裹送达的船的最低运载能力


		输出: 找到给定天数时，运载能力最低的吨位

		思路：
		  二分查找

		    运载能力吨位从最小到最大, 看哪个吨位能满足days天内搞定事情

	*/

	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }

	// minWeight := func(ws []int) int {
	// 	ans := 0
	// 	for _, w := range ws {
	// 		ans = max(ans, w)
	// 	}
	// 	return ans
	// }

	// 最小的吨位
	left := 1
	right := 0

	// 最大的吨位
	for _, w := range weights {
		right += w
	}

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

	sum := 0
	dayCnt := 0

	for _, w := range weights {
		if w > mid {
			return false
		}
		if sum+w >= mid { // 大于等于时，表示装满了，今天运输这些
			dayCnt++ // 大于天数累加++
			fmt.Printf("\n第%d天 mid: %d sum: %d \n", dayCnt, mid, sum)
			sum = 0
		} else {
			// 小于船吨位时，sum累加, 说明船还没装满
			sum += w
			fmt.Print(w, " ")
		}
	}

	// 说明mid吨位的船，可以在小于等于days天内搞定这个事情
	return dayCnt <= days
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

