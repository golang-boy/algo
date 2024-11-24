/*
 * @lc app=leetcode.cn id=LCR 187 lang=golang
 * @lcpr version=20003
 *
 * [LCR 187] 破冰游戏
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func iceBreakingGame(num int, target int) int {

	/*
	  num成员, target目标值

	  围成圆圈
	  第target号为离开圆桌，编号从0开始

	  返回最后一位离开的编号

	  递归实现

	  成员编号不变, 编号从0开始, 总共num-1位

	  如果离开，则从下一位开始

	  得有个全局状态，记录剩下的
	*/

	var state = make([]bool, num)
	var left = -1

	breakingGame(0, num, target, &state)

	for i, ok := range state {
		if !ok {
			left = i
		}
	}
	return left
}

func breakingGame(i, num int, target int, state *[]bool) {
	if len(*state) == num-1 {
		return
	}

	for {
		if !(*state)[i%num] {
			break
		}
		i += 1
	}

	m := i + target - 1
	removed := m % num

	(*state)[removed] = true
	 := (removed + 1)
	fmt.Println(i)
	breakingGame(next, num, target, state)
}

/*

  总结：
     计算公式没想明白，就开始做

	   对某个变化的数的进行取模，同时还得知道索引

	下一个起始位置
	    nextIndex = index + target

	模多少呢？
	    nextIndex % i  i为减少人后的人员总数
	
	求什么呢？求的是当num为输入时的位置索引信息

	 该题求的是索引, 如果剩余俩人，则在这俩人将数来数去



*/

// @lc code=end

/*
// @lcpr case=start
// 7\n4\n
// @lcpr case=end

// @lcpr case=start
// 12\n5\n
// @lcpr case=end

*/

