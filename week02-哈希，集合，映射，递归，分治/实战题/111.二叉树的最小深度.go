/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {

	// // 边界：当前节点为空时，返回深度为0
	// if root == nil {
	// 	return 0
	// }

	// // 递归，定义从顶向下定义，执行从低向上执行

	// leftmin := minDepth(root.Left)
	// rightmin := minDepth(root.Right)

	// if leftmin == 0 && rightmin != 0 {
	// 	return rightmin + 1
	// }

	// if rightmin == 0 && leftmin != 0 {
	// 	return leftmin + 1
	// }

	// if leftmin == 0 && rightmin == 0 {
	// 	return 1
	// }

	// return min(leftmin, rightmin) + 1

	// 另一种方式

	var ans, depth = math.MaxInt, 1

	if root == nil {
		return 0
	}
	calcMin(root, &ans, &depth)

	return ans

}

func calcMin(root *TreeNode, ans, depth *int) {

	if root.Left == nil && root.Right == nil {
		*ans = min(*ans, *depth)
		return
	}

	if root.Right != nil {
		*depth++
		calcMin(root.Right, ans, depth)
		*depth--
	}

	if root.Left != nil {
		*depth++
		calcMin(root.Left, ans, depth)
		*depth--
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

