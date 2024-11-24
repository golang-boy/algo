/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	/*
		 关键信息：
			  1. 找到与nums拥有相同度的最短子集
			      ans = min(len(subset), ans)
			  2. 返回子集长度 len(subset)
			  3. 最短连续子数组 ==>  连续的
		 思路：
		   1. 计算出nums中度对应的子集的起始和结束位置
		   2. 根据起始与结束位置，计算长度，返回

	*/

	q := map[int]int{}
	indexer := map[int][]int{}
	max := 0
	maxList := []int{}

	// 计算度,统计频数
	for i, n := range nums {
		q[n]++
		if l, ok := indexer[n]; ok {
			indexer[n] = append(l, i)
		} else {
			indexer[n] = append([]int{}, i)
		}
	}

	// 找到最大的频数，得出nums的度
	for k, v := range q {
		if v > max {
			max = v
			maxList = []int{k}
		}
		if v == max {
			maxList = append(maxList, k)
		}
	}

	// 找到相同度，最短长度的子集
	min := math.MaxInt32
	for _, num := range maxList {
		if length := getLength(indexer, num); length < min {
			min = length
		}
	}
	return min
}

func getLength(indexer map[int][]int, val int) int {
	// 定位起始和结束索引
	l := indexer[val]
	if len(l) == 1 {
		return 1
	}
	start := l[0]
	end := l[len(l)-1]
	return end - start + 1
}

/*
总结：
  1. 找到起始位置，和结束位置， 计算度
  2. 比较度与输入nums相同的子数组
  3. 找到最短的数组，返回长度

  上述实现，有些啰嗦，使用俩map来存储索引关系, 只需要一个map就行
*/

// @lc code=end

