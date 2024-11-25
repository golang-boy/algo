/*
 * @lc app=leetcode.cn id=113 lang=golang
 * @lcpr version=20003
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {

	/*
		思路：树的遍历

		从根节点开始递归

		路径类型为[]int, 深度遍历，每次递归时，+val， 进入下一级，直到最底层，判断是否等于目标值
	*/

	var res [][]int
	var chosen []int
	r(root, targetSum, &res, chosen)
	return res

}

func r(root *TreeNode, sum int, res *[][]int, chosen []int) {
	if root == nil { // 到了末尾
		return
	}

	sum -= root.Val

	// if sum < 0 {
	// 	return
	// }
	// 为啥这里不能剪枝？
	//   因为targetSum有正负数的原因,除非将正负数的情况进行分析

	//   不进行判断，下去后，chosen追加该值, 如果没到尾部，chosen就不用管

	chosen = append(chosen, root.Val)
	// 此处追加后的chosen和上一层的chosen不是一个chosen, 因此返回后，上一层没有root.Val的值
	// chosen是俩个指针，前面的部分指向相同的data,长度不同，因此，此处修改前面的数据时，上一层的也会改变
	// 值传递，一个长度变了，一个没变
	if root.Val == 7 {
		fmt.Println(chosen)
	}

	if root.Left == root.Right && root.Right == nil {
		if sum == 0 {
			temp := make([]int, len(chosen))
			copy(temp, chosen)
			*res = append(*res, temp)
		}
		return
	}

	r(root.Left, sum, res, chosen)
	r(root.Right, sum, res, chosen)
}

// @lc code=end

/*
// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,5,1]\n22\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

*/

