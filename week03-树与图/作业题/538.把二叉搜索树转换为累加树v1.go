/*
 * @lc app=leetcode.cn id=538 lang=golang
 * @lcpr version=20004
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {

	/*

	   二叉树转累加树

	   当前节点的值 = 旧值+所有大于自身的节点值之和, 即当前节点值 = 当前节点值 + 右子树节点值之和 + 父节点值

	   1. 先计算右子树，然后计算根，然后计算左子树
	   2. 递归

	   3. 当为空时，返回0

	*/

	if root == nil {
		return nil
	}

	leftTemp := 0

	if root.Right != nil && root.Right.Left != nil {
		leftTemp = root.Right.Left.Val
	}

	// 右子树
	right := convertBST(root.Right)
	if right != nil {
		root.Val += right.Val
		if right.Left != nil {
			root.Val += leftTemp
		}
	}

	// 不能加，否则递归时，会重复累加
	cur := root.Left
	for cur != nil {
		if cur.Right == nil {
			cur.Val += root.Val
			break
		}
		cur = cur.Right
	}

	// 左子树
	left := convertBST(root.Left)
	if left != nil {
		left.Val += root.Val
	}

	root.Left = left
	root.Right = right

	return root
}

/*

总结：
   总体思路正确， 但涉及到跨层级累加时， 逻辑上存在错误

   右侧递归完毕后，左侧的节点都要加根节点的值 => 最右侧的节点要加根节点的值, 然后递归, 但是递归时会重复累加

   跨层级计算时，递归处理会出问题, 根据题解，这种情况需要全局变量处理, 递归时，全局变量参与累加

*/

// @lc code=end

/*
// @lcpr case=start
// [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]\n
// @lcpr case=end

// @lcpr case=start
// [0,null,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4,1]\n
// @lcpr case=end

*/

