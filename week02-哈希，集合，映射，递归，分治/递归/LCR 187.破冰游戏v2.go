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
								但是全局状态 最大10w个，target有100w，太大了

								能不能分割子问题,子问题是不是重叠的?

								  f(num, target ) 时,删除 target-1 % num, 返回 target % num
								   f(num-1, target)时,删除 target-1 % num +1, 返回 target % num + 1

						              f(num-1, target) = x  num-1个人时，x位置存活, num时，存活的是哪个

				                      (x - target) %num = f(num, target)

			       f(num, target )      ((x-target) % num + num) %num
		                                (x-target) % num

	*/

	if num == 1 {
		return 0
	}
	x := iceBreakingGame(num-1, target)
	return (x + target) % num
}

/*

  总结：
    下一个起始位置
        nextIndex = index + target

    模多少呢？
        nextIndex % i
         i为减少人后的人员总数

    求什么呢？求的是当num为输入时的位置索引信息

     该题求的是索引, 如果剩余俩人，则在这俩人将数来数去

     同模取余操作，从头到尾和从尾到头一样，算的是一端的索引
     需要确定每次开始的位置

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

