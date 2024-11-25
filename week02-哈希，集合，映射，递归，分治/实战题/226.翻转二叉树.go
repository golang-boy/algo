/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	/*
	  关键信息:
	    1. 一个二叉树
	    2. 翻转，左右子树交换，左右节点交换
	    3. 返回根节点
	  分析：
	    最小问题：一个节点，左右节点交换
	    其他: 子树交换
	    使用递归: 先判定最小问题，然后处理具体逻辑，最后递归
	*/

	// 边界
	if root == nil {
		return nil
	}

	//最小问题： 左右子节点翻转
	root.Left, root.Right = root.Right, root.Left

	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)

	return root
}

// @lc code=end

