/*
 * @lc app=leetcode.cn id=590 lang=golang
 * @lcpr version=20003
 *
 * [590] N 叉树的后序遍历
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

func postorder(root *Node) []int {

	var ans []int

	if root == nil {
		return ans
	}

	for _, c := range root.Children {
		r := postorder(c)
		ans = append(ans, r...)
	}

	ans = append(ans, root.Val)

	return ans

}

// @lc code=end

/*
// @lcpr case=start
// [1,null,3,2,4,null,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]\n
// @lcpr case=end

*/

