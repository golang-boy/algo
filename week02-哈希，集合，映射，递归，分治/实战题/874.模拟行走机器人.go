/*
 * @Author: 刘慧东
 * @Date: 2024-11-19 10:04:02
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-19 13:39:33
 */
/*
 * @lc app=leetcode.cn id=874 lang=golang
 * @lcpr version=20003
 *
 * [874] 模拟行走机器人
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	/*
		关键信息：
		   1. 初始位置（0,0）, 朝北
		   2. 可接收的命令
		       左转90度 -2；右转90度 -1; 前进 x步 x >= 1  && x <= 9
		   3. 有障碍物坐标, 障碍物有多个，是一个坐标数组
		   4. 遇到障碍物等待指令
		   5. 返回远点到当前点的最大距离, 原点可能有障碍物
		        d*d := x*x + y*y

		分析：
		   1. 不论是左转还是右转，都是90度，即四个方向
		       坐标x,y, 不同方向，对应x或y的正负变化
			     默认朝北： +y   {0,1}     朝西：  -x  {-1, 0}
				 朝南：    -y   {0,-1}     朝东：  +x   {1, 0}
		   2. 前进多少步，表明在当前方向上前进
		         dir * step
		   3. 前进时，需要判断是否有障碍
		         每走一步需要判断一次
		   4. 有障碍，则坐标在前一个位置处，执行后续指令
		   5. 直到执行完毕指令进行结果计算
	*/

	ans := 0

	var n complex64 = 0 + 1i
	var s complex64 = 0 + -1i
	var w complex64 = -1 + 0i
	var e complex64 = 1 + 0i

	var robot complex64 = 0 + 0i
	dir := n

	getDir := func(d complex64, command int) complex64 {
		switch command {
		case -1:
			if d == n {
				d = e
			} else if d == e {
				d = s
			} else if d == s {
				d = w
			} else if d == w {
				d = n
			}
		case -2:
			if d == n {
				d = w
			} else if d == w {
				d = s
			} else if d == s {
				d = e
			} else if d == e {
				d = n
			}
		default:
		}
		return d
	}

	nextStep := func(p complex64, d complex64, step int) complex64 {
		return p + d*complex(float32(step), 0)
	}

	obstaclesMap := make(map[complex64]struct{})

	for i := 0; i < len(obstacles); i++ {
		o := complex(float32(obstacles[i][0]), float32(obstacles[i][1]))
		obstaclesMap[o] = struct{}{}
	}

	isObstacle := func(p complex64) bool {
		_, ok := obstaclesMap[p]
		return ok
	}

	for i := 0; i < len(commands); i++ {
		if commands[i] < 0 {
			dir = getDir(dir, commands[i])
		} else {
			step := commands[i]
			next := robot
			for j := 0; j < step; j++ {
				next = nextStep(next, dir, 1)
				if isObstacle(next) {
					break
				}
				robot = next // 前进
			}
		}
		// 每走完一步，返回的是距离原点最远的位置
		ans = max(ans, int(real(robot)*real(robot)+imag(robot)*imag(robot)))
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [4,-1,3]\n[]\n
// @lcpr case=end

// @lcpr case=start
// [4,-1,4,-2,4]\n[[2,4]]\n
// @lcpr case=end

// @lcpr case=start
// [6,-1,-1,6]\n[]\n
// @lcpr case=end

*/

