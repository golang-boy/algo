/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {

	// // 边界
	// // 当前节点为空时，深度为0
	// if root == nil {
	// 	return 0
	// }

	// // 递归方式：从顶向下定义，从底向上执行, 执行时累计深度

	// left := maxDepth(root.Left)
	// right := maxDepth(root.Right)

	// if left <= right {
	// 	return right + 1
	// } else {
	// 	return left + 1
	// }

	// 另一种是维护一个变量depth
	// 从上到下执行，每执行一步，维护一次状态
	// 直到最底层时完成结果赋值
	var ans, depth = 0, 0
	calc(root, &ans, &depth)
	return ans
}

func calc(root *TreeNode, ans, depth *int) {
	// 到了最下面
	if root == nil {
		// 已经到达最下面,把累计的深度进行赋值
		*ans = max(*depth, *ans)
		return
	}

	*depth++
	calc(root.Left, ans, depth) // 执行进入左边, 进入前加1
	// 左边完毕后，执行右边, 状态需要还原，因此depth--；
	// 后续需要进入右边，因此要++，两者抵消
	calc(root.Right, ans, depth)
	// 右侧计算完毕，回到上一层，depth--
	*depth--

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

