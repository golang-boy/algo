/*
 * @lc app=leetcode.cn id=911 lang=golang
 * @lcpr version=20004
 *
 * [911] 在线选举
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type TopVotedCandidate struct {
	persons []int
	times   []int
}

type Stat struct {
	votes int
	index int
}

func Constructor(persons []int, times []int) TopVotedCandidate {

	/*
	   输入：
	        1. persons和times, 第i张票是times[i]时刻，投给persons[i]的票
	        2. q() 为查询，查询t时刻，选举中领先的人的编号

	   输出：
	        t时刻，票最多的人

	   思路：
	        times中找到 val <=t 的对应的索引, 在persions中统计之前投票数

	*/

	stat := make(map[int]struct {
		votes int
		index int
	})

	return TopVotedCandidate{
		persons: persons,
		times:   times,
		stat:    stat,
	}

}

func (this *TopVotedCandidate) Q(t int) int {

	left := 0
	right := len(this.times) - 1

	for left < right {
		mid := (left + right + 1) / 2
		if this.times[mid] <= t {
			left = mid
		} else {
			right = mid - 1
		}
	}

	var maxVote int
	var latestIndex int
	var person int

	// 票数最多且最近的人   最近表示i最大
	for i := right; i >= 0; i-- {
		if v, ok := stat[this.persons[i]]; ok {
			stat[this.persons[i]] = Stat{votes: v.votes + 1, index: v.index}
		} else {
			stat[this.persons[i]] = Stat{votes: 1, index: i}
		}
	}

	for k, v := range stat {
		if v.votes > maxVote {
			maxVote = v.votes
			latestIndex = v.index
			person = k
		} else if v.votes == maxVote {
			if v.index > latestIndex {
				latestIndex = v.index
				person = k
			}
		}
	}

	return person

}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
// @lc code=end

/*
// @lcpr case=start
// ["TopVotedCandidate", "q", "q", "q", "q", "q", "q"][[[0, 1, 1, 0, 0, 1, 0], [0, 5, 10, 15, 20, 25, 30]], [3], [12], [25], [15], [24], [8]]\n
// @lcpr case=end

*/

