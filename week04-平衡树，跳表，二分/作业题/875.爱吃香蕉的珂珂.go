/*
 * @lc app=leetcode.cn id=875 lang=golang
 * @lcpr version=20004
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minEatingSpeed(piles []int, h int) int {

	/*
		    输入： piles[i] 一堆香蕉，h个小时
			输出： 求吃香蕉的速度k，k要最小的值

			  h小时内吃尽可能多的香蕉,

	*/

	left := 1
	right := max(piles)
	// right := math.MaxInt

	for left < right {
		mid := (left + right) / 2
		if validate(mid, piles, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func min(p []int) int {
	ans := math.MaxInt
	for _, pp := range p {
		if ans > pp {
			ans = pp
		}
	}
	return ans
}

func max(p []int) int {
	ans := 0
	for _, pp := range p {
		if ans < pp {
			ans = pp
		}
	}
	return ans
}

func validate(k int, piles []int, h int) bool {
	// accPiles := 0
	hCnt := 0

	// 每小时要吃mid根香蕉
	for _, p := range piles {

		hCnt += p / k
		if p%k != 0 {
			hCnt++
		}

		// // 累积的香蕉+当前堆的香蕉 > k
		// if accPiles+p > k {
		// 	accPiles += p - k
		// } else {
		// 	accPiles = 0
		// }
		// hCnt++
	}

	return hCnt <= h
}

/*

总结：
    存在的问题：
	    1. 审题不清楚, 以为没吃完的会放到下一堆，一起吃。实际情况是，没吃完的，下一个小时继续吃这堆，直到都吃完， 在去下一堆
		2. 起始条件的数量也是需要关注的一点

	总体来说，该题要求取最小值，最直接的思路就是二分

*/

// @lc code=end

/*
// @lcpr case=start
// [3,6,7,11]\n8\n
// @lcpr case=end

// @lcpr case=start
// [30,11,23,4,20]\n5\n
// @lcpr case=end

// @lcpr case=start
// [30,11,23,4,20]\n6\n
// @lcpr case=end

*/

