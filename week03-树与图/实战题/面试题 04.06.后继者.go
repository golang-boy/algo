
/*
 * @lc app=leetcode.cn id=面试题 04.06 lang=golang
 * @lcpr version=20003
 *
 * [面试题 04.06] 后继者
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
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	/*

		   1. 在root中找到p,
		   2. 找到p的下一个节点（中序后继） 指比p大的最小值
		   3. 找不到是返回null


		   先找p，然后找p的后继

		   树结构需要递归
		     子问题： 先找左子树,后找右子树，中序后继的情况，需要找到左子树后，判断根的情况，然后在右子树

			 边界： 如果p.Val == root.Val, 找到p, 在p的右子树找最小的, 如果右子树为空，则返回root
	*/

	if root == nil {
		return nil
	}

	if p.Val < root.Val {
		l := inorderSuccessor(root.Left, p)
		if l != nil {
			return l
		}
		return root
	}

	return inorderSuccessor(root.Right, p)
}

/*
总结：
    子问题还是没分析清楚，分析子问题，思路不清晰

	子问题不清晰，边界也不会清晰

*/

// @lc code=end

/*
// @lcpr case=start
// [2,1,3]\n12/ \1   3\n
// @lcpr case=end

// @lcpr case=start
// [5,3,6,2,4,null,null,1]\n65/ \3   6/ \2   4/   1\n
// @lcpr case=end

*/

