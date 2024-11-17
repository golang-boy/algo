/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {

	/*

				   关键信息：
				      1. 一个序列
					  2. 带下为k的滑动窗口
					  3. 每移动一次，求一次最大值

		           分析：
				     1. 遍历序列
					 2. 维护滑动窗口   ==》 怎么维护? 什么时候计算?
					     1. 左边一个指针，右边一个指针
						 2. 俩指针之间的距离为k，同时移动
						 3. 移动完毕后，计算一次最大值
						 4. 右指针移到末尾时，结束
					 3. 返回最大值


					 简单的做法,
					   1. 维护一个二叉大顶堆,大小为k
					   2. 每次移动时，将新元素加入堆中，同时将堆顶元素加入结果中
					      但是出队时，怎么清理元素呢？

						 这个做法不成熟!

					单调队列：
					  队列中维护一个递减的序列
					     当不满足递减时，将元素出队

	*/

	// 遍历序列
	maxList := []int{}

	q := list.New() // 存储下标, 队列里维护索引

	for i := 0; i < len(nums); i++ {

		// i 对应的right,  left = i-k+1
		// 左边移动
		for q.Len() != 0 {

			front := q.Front() // 队列头部
			index := front.Value.(int)

			if index > i-k { // i-k为left-1, 正好是左侧边界外的位置
				break
			}
			// 当对头的index <= i-k时， 把不在范围内的元素都出队, 移动左侧边界
			// 当队头的索引不在滑动窗口内时，队头出队
			q.Remove(front) // 左边的出队
		}

		// 右边移动
		for q.Len() != 0 {

			tail := q.Back() // 队列尾部
			index := tail.Value.(int)

			if nums[index] > nums[i] { // 1 > -3
				break
			}

			q.Remove(tail) // 出队
		}
		q.PushBack(i) // 入队

		// 移动完毕后，计算一次最大值
		if i >= k-1 {
			front := q.Front().Value.(int)
			maxList = append(maxList, nums[front])
		}
	}

	return maxList

}

/*
   总结：
      1. 队列中维护一个递减的序列, 队列中存储的是下标
	  2. 遍历数组时，i对应的右侧指针
	       移动右侧指针时，判断队尾与当前值的大小，如果队尾小于等于，队尾出队, 直到队尾大于当前值
	       移动左侧指针时，判断队头是否在滑动窗口内，不在则队头出队。
		       在不在滑动窗口内, 看对头的索引 是否小于left左侧边界, left = i-k+1。不在则队头出队

	  3. 通过移动左右指针, 队列中的元素始终优先，并且是递减的
	  4. 移动完毕后，计算一次

*/

// @lc code=end

