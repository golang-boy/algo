/*
 * @lc app=leetcode.cn id=297 lang=golang
 * @lcpr version=20003
 *
 * [297] 二叉树的序列化与反序列化
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

type Codec struct {
}

/*

    序列化： 该过程需要遍历树,遍历方式: 前,中,后序遍历

    反序列化： 树的反序列化过程中，只有前序没法知道左右子树的内容，只有中序，没法知道根节点，只有后序，和前序一样

    因此，反序列化时，要么前中，有么后中,此时能构造出树, 对应的序列化时,也必须前中，或者后中

	序列化后，如何保存？

	因为长度一样，前面是前序，后面是中序，依次拼接即可

*/

func Constructor() Codec {

	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	p := preorder(root)
	r := strings.Join(p, ",")
	return r
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	d := strings.Split(data, ",")

	var i int = 0
	return buildTree(d, &i)
}

func buildTree(p []string, i *int) *TreeNode {
	if p[*i] == "null" {
		(*i)++
		return nil
	}

	val, _ := strconv.Atoi(p[*i])

	root := &TreeNode{
		Val: val,
	}
	// 这里的p得共享，i为全局值，不停的往后移
	(*i)++

	root.Left = buildTree(p, i)

	root.Right = buildTree(p, i)
	return root
}

func preorder(root *TreeNode) []string {
	var ans []string
	if root == nil {
		ans = append(ans, "null")
		return ans
	}
	ans = append(ans, fmt.Sprintf("%d", root.Val))
	ans = append(ans, preorder(root.Left)...)
	ans = append(ans, preorder(root.Right)...)
	return ans
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

/*
总结：
   序列化: 使用任一种树遍历方法，上述使用了前序遍历
      特殊处理的地方是叶子节点为null,这样就知道了左右节点的边界

   反序列化：
      处理序列时，节点结束的条件为，遇到null时，return nil。同时索引加一，指向下一个

	  这里最重要的思想是，这个索引是要全局的，不论递归怎么递归，这个存储索引的变量需要全局


	  做的过程中，脑子不清楚，老纠结于代码细节

*/

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,null,4,5]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/

