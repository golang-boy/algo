/*
 * @Author: 刘慧东
 * @Date: 2024-11-27 10:42:32
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-27 12:14:29
 */
/*
 * @lc app=leetcode.cn id=429 lang=golang
 * @lcpr version=20003
 *
 * [429] N 叉树的层序遍历
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	var ans [][]int
	var q []*Node
	var q1 []*Node

	if root == nil {
		return ans
	}

	q = append(q, root)

	for {
		if len(q) == 0 {
			break
		}
		var temp []int

		q1 = q
		q = []*Node{}
		for _, n := range q1 {
			temp = append(temp, n.Val)
			q = append(q, n.Children...)
		}
		ans = append(ans, temp)
	}

	return ans

}

/*

总结：

   层序遍历：将每一层的值append后，开始下一层的操作, 直到最后一层完毕

   本题中，将同一层的值，放入一个数组，因此需要体现层的概念

   通常的解法是使用队列:
      1. 一个大循环，不停判断队列是否为空，不为空时，元素出队,加入结果集中
	  2. 为了体现层级关系，每次入队时，标明该节点的层级, 即队列元素，除了有节点信息外，额外有个层的信息

	上述解法没使用队列：
	  1.数组，没有入队出队的情况，或者说，叫双队列
	  2.将每一层放入队列，处理时，当层所有值出队，循环处理当前层
	  3.当前层处理时，下一层的依次进入另一个队列
	  4. 当下一次循环时，复制给前一个队列，情况当前队列（相当于全部出栈，去接收新的内容）


*/

// @lc code=end

/*
// @lcpr case=start
// [1,null,3,2,4,null,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]\n
// @lcpr case=end

*/

