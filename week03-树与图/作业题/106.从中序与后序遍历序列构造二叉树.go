/*
 * @lc app=leetcode.cn id=106 lang=golang
 * @lcpr version=20004
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {

	/*
	   第一个参数为中序: 左-根-右
	   第二个为后序: 左-右-根
	   序列中值都不同

	   1. 先从后序序列中拿到root
	   2. 在中序数组中拿到root所在的位置，左边为左子树，右边为右子树
	   3. 在后续数组中，因为左右子树长度在俩数组中长度一样，在后序数组中找到下一次的左右序列
	   3. 递归处理下一个
	*/

	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]

	var mid = 0

	for i, v := range inorder {
		if v == rootVal {
			mid = i
			break
		}
	}

	root := &TreeNode{
		Val: rootVal,
	}

	// root = 3   mid = 1
	// lin   [ 9 ]
	// lpost [ 9 ]

	// rin   [ 15, 20, 7 ]  [2:]
	// rpost [ 15, 7, 20 ]  [1:]

	leftInorder := inorder[0:mid]
	leftPostorder := postorder[0:mid]
	rightInorder := inorder[mid+1:]
	rightPostorder := postorder[mid : len(postorder)-1]

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)

	return root
}

/*
总结：
    左右子树的边界得计算好
	要恢复树结构, 根本目的是要构建根节点，以及左右子树

*/

// @lc code=end

/*
// @lcpr case=start
// [9,3,15,20,7]\n[9,15,7,20,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/

