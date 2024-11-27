/*
 * @lc app=leetcode.cn id=105 lang=golang
 * @lcpr version=20003
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	/*

		前序遍历，root在前面，然后是l和r
		中序遍历，l，root, r

		思路：
		 1. 先拿到root节点，然后拿到l，在然后确定r
		 2. 从中序列表中，确定l是不是左，r是不是右
		 3. 确定完毕后，开始下一个节点的构造，逻辑相同，直到末尾
		 4. 每个都是子树,构造规则相同
		 5. 边界条件，输入数组都为空
	*/

	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	var mid = 0

	// 必定能找到
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			mid = i
			break
		}
	}

	// tree [1,nil, ]
	// pre  []
	// in   []

	root.Left = buildTree(preorder[1:mid+1], inorder[0:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}

/*

总结：

    1. 确定重叠子问题
          构造根节点，构造左子树，构造右子树
	2. 递归输入
	      构造根节点：从前序数组中拿
		  构造左子树：从之前的数组中，拿到构造左子树的前序,中序数组
		  构造左子树：从之前的数组中，拿到构造右子树的前序,中序数组
	3. 返回root


	问题：怎么拿到左子树和右子树的序列?

	   1. 前序中拿到根，中序中找根位置，左边为左子树需要的中序序列，右边为右子树需要的中序序列, 由此确定中序序列
	   2. 确定左子树的前序序列，根-左-右，长度相同，找到的位置长度，在前序序列中定位左子树的长度，由此确定递归时的前序序列

	   花了2小时, 思路不对, 钻入了细节，范围没控制住

*/

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,15,7]\n[9,3,15,20,7]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/

