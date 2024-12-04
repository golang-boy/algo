/*
 * @lc app=leetcode.cn id=450 lang=golang
 * @lcpr version=20003
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	/*

	  1. 在左子树时，需要找到右子树最小的节点，与当前节点替换
	     找不到时，即右子树为nil, 左子树的根节点与当前节点替换

	  2. 在右子树时，需要找到右子树最小的节点，与当前节点替换
	     找不到时，即右子树为nil, 左子树的根节点与当前节点替换

	  3. 怎么替换

	*/

	if root == nil {
		return nil
	}

	if root.Val == key {

		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		n := min(root.Right)
		n.Right = deleteNode(root.Right, n.Val)
		n.Left = root.Left
		return n

	}

	// 左子树
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func min(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	m := min(root.Left)
	// root.Left = nil
	return m
}

// @lc code=end

/*
// @lcpr case=start
// [5,3,6,2,4,null,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [5,3,6,2,4,null,7]\n0\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/

