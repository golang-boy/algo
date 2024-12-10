/*
 * @lc app=leetcode.cn id=410 lang=golang
 * @lcpr version=20004
 *
 * [410] 分割数组的最大值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func splitArray(nums []int, k int) int {
	/*
		   输入：数组元素都为正数,  分组数k
		   输出：分割后每个分组求和，找到最大值,
		         这个最大值在所有情况的分组中最小


		   朴素的思路是，先分组，然后求每个分组的和，然后比较分组和的大小

		   题解使用二分查找， 条件判定为验证是否符合要求, 即转换为二分查找+条件判断

		   先确定左右边界和判断条件：

		   判定条件：
			   1. 分组是否等于k，和是不是等于某个值
			   2. 在这些值中找最小的值

			   有哪些值？
			       值的序列就是分组和, 假设left从0开始，right最大值怎么确定？
				       都加起来, left要不要从0开始？
					      得从最大的开始，最大分组时所有nums的和，最小分组为每个是一个

			  二分序列：
			     1.最小值左侧为序列长度每个为一个分组,因此找最大的，
				 2.最大值为整个序列一个分组
				   中间其他分组的值处于俩这之间, 不同分组情况和最大值不同，也可能相同，要找最小的，即为找左边界问题

				   假设mid为要找的值,在所有mid中，找个最小的。 数组中每累加一个数,看看小于mid的分组数多还是少
				   分组数大于k,说明这个mid过于小，需要变大，left往右移
				   分组数小于k,说明mid有点大，需要调小，right左移，

	*/

	left := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	right := 0
	for _, n := range nums {
		right += n
		left = max(left, n)
	}

	for left < right {
		mid := (left + right) / 2
		if validate(mid, nums, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func validate(mid int, nums []int, k int) bool {
	sum := 0
	groupCnt := 1
	for _, num := range nums {
		if sum+num > mid { // 超过了mid, 就是一个新分组
			sum = num
			groupCnt++
		} else {
			sum += num
		}
	}

	// 判断有分组数是否符合要求
	return groupCnt <= k
	// 可能有不同的分组, 小于时，继续往右找
}

/*
   总结：
*/

// @lc code=end

/*
// @lcpr case=start
// [7,2,5,10,8]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,4,4]\n3\n
// @lcpr case=end

*/

