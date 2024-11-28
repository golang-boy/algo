/*
 * @lc app=leetcode.cn id=235 lang=golang
 * @lcpr version=20003
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	/*

		输入：根节点root, p,q俩节点
		输出：最近的公共祖先

		必定能找到p,q

		子问题：

		如果都在左子树，递归后，返回左子树的结果
		   如果左子树没找到时，返回nil

		如果都在右子树，递归后，返回右子树的结果
		   如果右子树没找到时，返回nil

		如果在俩边，返回父节点
		     递归返回后，如果都不是nil， 说明俩边分别找到了p，q, 因此需要返回当前节点

		边界：
		   当前节点如果到了叶子，匹配后，返回当前节点
		   如果为nil，返回nil
		   如果不匹配，继续左右子树结果
	*/

	if root == nil {
		return nil
	}

	//  当前节点如果到了叶子，匹配后，返回当前节点
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l == nil {
		return r
	}

	if r == nil {
		return l
	}

	return root
}

/*

总结
     将该题当做了二叉树的最近公共祖先进行了求解, 审题不清
	    (后续做题，必须遵循，基本的做题原则:
		     输入，输出，处理逻辑, 变化的量，和不变的量等信息分析)

	做该题时，问题没分析清楚，依旧是在套模板

	  没有明确 递归函数的作用与意义

	     什么时候递归？
		 递归返回结果怎么处理？
		 递归什么时候返回?


*/

// @lc code=end

/*
// @lcpr case=start
// [6,2,8,0,4,7,9,null,null,3,5]\n2\n8\n
// @lcpr case=end

// @lcpr case=start
// [6,2,8,0,4,7,9,null,null,3,5]\n2\n4\n
// @lcpr case=end

*/

