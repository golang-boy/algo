/*
 * @lc app=leetcode.cn id=701 lang=golang
 * @lcpr version=20003
 *
 * [701] 二叉搜索树中的插入操作
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 二叉搜索树
	/*
	   1. 输入数据与已有值都不同
	   2. 与根节点比较，如果小于则在左子树，如果大于则在右子树
	   3. 插入的值，是否大约当前节点值，如果小于，则左子树找最大的

	*/

	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)

	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

/*
总结：

     树的操作，肯定都是递归
	    递归，需要找到递归边界，和递归子问题, 以及处理递归返回的结果

		上述三者确定后，该算法也就完成

		递归边界：

		    根据定义：在树中遍历，如果找到叶子节点后，且为nil时，说明找到了合适的位置
			   此时需要根据val值，构造一个节点插入

			子问题：由于有左右俩测的节点，递归是val小于当值，去左侧子树找，否则去右侧, 子树也是树

			返回结果处理：需要把新插入的点链接到树中，因此返回的新节点，要么在左，要么在右
*/

// @lc code=end

/*
// @lcpr case=start
// [4,2,7,1,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [40,20,60,10,30,50,70]\n25\n
// @lcpr case=end

// @lcpr case=start
// [4,2,7,1,3,null,null,null,null,null,null]\n5\n
// @lcpr case=end

*/

