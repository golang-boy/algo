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
		      * 交换排序
		          冒泡，快速
		      * 插入排序
		          简单插入排序，希尔排序
				  每次从前到后，考虑每个未排序的数据，在已排序的序列中找位置插入
		      * 选择排序
		          简单选择排序，堆排序
		          每次从未排序的数据找最小值，放到以排序列表末尾
		      * 归并排序
		          二路归并，多路归并

		   2. 基于非比较的排序，时间复杂度受元素范围和分布等多种因素影响
		      * 计数排序
			     输入的数据必须要是确定范围的整数
			  * 基数排序
			     将数据切割为一位位数字，从低到高分别计数排序
			  * 桶排序
			     输入的数据服从均匀分布, 将数据分到有限数量的桶里，每个桶在排序
	*/

	minVal := min(nums)
	maxVal := max(nums)
	arrays := make([]int, maxVal-minVal+1)
	ans := []int{}

	for _, num := range nums {
		arrays[num-minVal]++
	}

	for num, cnt := range arrays {
		for cnt > 0 {
			ans = append(ans, num+minVal)
			cnt--
		}
	}
	return ans
}

func max(nums []int) int {
	ans := 0
	for _, num := range nums {
		if num > ans {
			ans = num
		}
	}
	return ans

}

func min(nums []int) int {
	ans := math.MaxInt
	for _, num := range nums {
		if num < ans {
			ans = num
		}
	}
	return ans
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

