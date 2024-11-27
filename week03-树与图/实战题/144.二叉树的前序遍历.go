/*
 * @Author: 刘慧东
 * @Date: 2024-11-27 10:21:23
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-27 10:27:03
 */
/*
 * @lc app=leetcode.cn id=144 lang=golang
 * @lcpr version=20003
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {

	var ans []int

	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)

	left := preorderTraversal(root.Left)
	ans = append(ans, left...)

	right := preorderTraversal(root.Right)
	// ans = append(ans, left...)
	ans = append(ans, right...)
	return ans
}

/*
总结：
	除了根外，左子树的append能不能放在末尾？

	 经过实验，可以放在最下面，只需要保持放入队列的顺序就行
*/

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

