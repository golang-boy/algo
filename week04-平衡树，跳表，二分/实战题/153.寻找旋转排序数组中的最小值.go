/*
 * @lc app=leetcode.cn id=153 lang=golang
 * @lcpr version=20004
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findMin(nums []int) int {
	/*

	  输入：
	   1. 长度为n，升序排列
	   2. 元素值不重复
	  输出：
	    找出数组中最小元素，并返回

	  找到原数组的起止位置，即为最小值

	   思路一：移动元素位置，每次把最前面的移到最后面，判断条件前面小于后面时停止
	         但是时间复杂度是O(n), 元素越多，移动次数越多
	   思路二：元素不动，移动索引位置, 判断左侧的是否小于右侧，小于则停.然后缩小范围
	         该方法也是O(n)
	*/

	left := 0
	right := len(nums) - 1
	ans := -5001

	for left <= right {

		// 小于右侧
		if nums[left] <= nums[right] {
			ans = nums[left]
			break
		}

		// 左侧大于右侧
		if nums[left] > nums[right] {
			left++
		}

	}
	return ans

	// left := 0
	// right := len(nums)
	// for left < right {
	// 	mid := (right + left) / 2

	// 	if nums[mid] >= nums[right] {
	// 		right = mid
	// 	} else {
	// 		left = mid + 1
	// 	}
	// }

	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (right + left) / 2

		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[right]
}

/*

总结：

   一个是判定的不等式，如果需要包含，则为right=mid

   套用模版时，
   一定要注意，该问题是否有解，有解 left，right初始化时为0，n-1

   如果无解，模板1.1为 0，n,  模版1.2为-1，n-1, 同时mid计算时，需要额外加个1

    总结：
	    思路整体正确，但是还是差一点，做法没符合要求，不是log(n), 只要出现找位置，且要求是log(n)必定考虑二分


*/

// @lc code=end

/*
// @lcpr case=start
// [3,4,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,7,0,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [11,13,15,17]\n
// @lcpr case=end

*/

