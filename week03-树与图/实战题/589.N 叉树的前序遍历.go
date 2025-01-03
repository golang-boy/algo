/*
 * @lc app=leetcode.cn id=589 lang=golang
 * @lcpr version=20003
 *
 * [589] N 叉树的前序遍历
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	var ans []int
	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)
	for _, c := range root.Children {
		a := preorder(c)
		ans = append(ans, a...)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,null,3,2,4,null,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]\n
// @lcpr case=end

*/

