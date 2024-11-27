/*
 * @lc app=leetcode.cn id=94 lang=golang
 * @lcpr version=20003
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	/*

	  前序遍历：先根，然后左子树，最后右子树
	  中序遍历：先左子树，然后根，最后右子树
	  后序遍历：先左子树，然后右子树，最后根
	*/
	var res []int
	if root == nil {
		return res
	}

	left := inorderTraversal(root.Left)

	res = append(res, left...)
	res = append(res, root.Val)

	right := inorderTraversal(root.Right)
	res = append(res, right...)
	return res

}

/*

总结：
   1. 概念要清晰，二叉树不同遍历方式,只有根在移动, 左右子树始终从左到右
   2. 根在最前面处理，为前序遍历
   3. 根在中间位置，为中序遍历
   4. 根在最后，为后续遍历
   5. 结果的append位置也很重要


*/

// @lc code=end

/*
// @lcpr case=start
// [1,null,2,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

