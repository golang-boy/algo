/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
// func isValidBST(root *TreeNode) bool {
//
// 	/*
// 		关键信息：
// 		   边界条件： root == nil
// 		   返回bool值
// 		   二叉搜索树：左子树 < 根节点 < 右子树
// 	*/
//
// 	if root == nil { // 等于nil时，返回true， 否则永远是false
// 		return true
// 	}
//
// 	if root.Left != nil && root.Left.Val >= root.Val {
// 		return false
// 	}
//
// 	if root.Right != nil && root.Right.Val <= root.Val {
// 		return false
// 	}
//
// 	left := isValidBST(root.Left)
// 	rgiht := isValidBST(root.Right)
//
// 	return left && rgiht
// }

// # 错误原因：未能理解二叉搜索树的概念: 左子树所有节点的值 < 根节点值 < 右子树所有节点值

// 重新开始
func isValidBST(root *TreeNode) bool {

	// 边界
	if root == nil {
		return true
	}

	// // 判断当前左右节点是否符合二叉搜索树(没必要)
	// if root.Left != nil && root.Left.Val >= root.Val {
	// 	return false
	// }
	// if root.Right != nil && root.Right.Val <= root.Val {
	// 	return false
	// }

	// 需要理清楚逻辑执行顺序：
	//    1. 先判断当前节点是否符合二叉搜索树
	//    2. 然后分别看左右子树是否符合

	// 如何保证左侧的所有节点值比根小？
	max := getLeftMax(root.Left)
	min := getRightMin(root.Right)

	// 左右根节点，都等于时，不是二叉搜索树
	if max != nil {
		if max.Val >= root.Val {
			return false
		}
	}

	if min != nil {
		if min.Val <= root.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

func getLeftMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	}
	return getLeftMax(root.Right)
}

func getRightMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return getRightMin(root.Left)
}

// @lc code=end

