/*
 * @lc app=leetcode.cn id=327 lang=golang
 * @lcpr version=20004
 *
 * [327] 区间和的个数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countRangeSum(nums []int, lower int, upper int) int {
	/*
	   输入nums,以及lower和upper,
	   输出满足条件的区间和的数量

	      区间和 s(i,j) 在nums中从i到j元素之和, i<=j
	      区间和满足  upper >= s(i,j) >= lower

	    归并排序，左右两侧有结果时，判断是否满足条件，满足条件则结果累加
	*/

	ans := 0
	mergeSort(nums, 0, len(nums)-1, lower, upper, &ans)
	return ans
}

func mergeSort(nums []int, l, r, lower, upper int, ans *int) {
	if l >= r {
		sum := 0
		for k := l; k <= r; k++ {
			sum += nums[k]
		}
		if sum >= lower && sum <= upper {
			fmt.Println(l, r)
			*ans += 1
		}

		return
	}
	mid := (l + r) / 2

	calc(nums, l, mid, r, lower, upper, ans)
	mergeSort(nums, l, mid, lower, upper, ans)
	mergeSort(nums, mid+1, r, lower, upper, ans)
	merge(nums, l, mid, r)
}

func calc(nums []int, l, mid, r, lower, upper int, ans *int) {

	// 前缀和的影子
	for j := mid + 1; j <= r; j++ {
		// 右侧确定时，左侧累加
		for i := l; i <= mid; i++ {

			// 确定了前后索引后，累加中间的值，并判断
			// 问题是，左右相等的情况会漏掉
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum >= lower && sum <= upper {
				fmt.Println(i, j, sum, nums[i:j+1], nums)
				*ans += 1
			}
		}
	}
}

func merge(nums []int, l, mid, r int) {
	temp := make([]int, r-l+1)

	i := l
	j := mid + 1
	k := 0

	for ; i <= mid && j <= r; k++ {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		temp[k] = nums[i]
		k++
	}
	for ; j <= r; j++ {
		temp[k] = nums[j]
		k++
	}

	for i := range temp {
		nums[l+i] = temp[i]
	}
}

/*
总结： 大体逻辑基本程序，前缀和的思想不知道怎么实现

     排序后计算，则索引不对，排序前计算，非前缀和（上述方式）时间超限

*/

// @lc code=end

/*
// @lcpr case=start
// [-2,5,-1]\n-2\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,-3,-3,1,1,2]\n3\n5
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n0\n
// @lcpr case=end

*/

