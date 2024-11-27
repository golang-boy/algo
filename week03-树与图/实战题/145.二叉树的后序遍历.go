
/*
 * @lc app=leetcode.cn id=145 lang=golang
 * @lcpr version=20003
 *
 * [145] 二叉树的后序遍历
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	left := postorderTraversal(root.Left)
	ans = append(ans, left...)
	right := postorderTraversal(root.Right)
	ans = append(ans, right...)
	ans = append(ans, root.Val)
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,null,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,null,8,null,null,6,7,9]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

