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
	i := inorder(root)

	p = append(p, i...)

	r := strings.Join(p, ",")

	fmt.Println(r)
	return r

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	fmt.Println(data)
	d := strings.Split(data, ",")

	if len(d)%2 != 0 {
		// 非法输入
		return nil
	}

	pre := d[:len(d)/2]

	in := d[len(d)/2:]

	return buildTree(pre, in)
}

func buildTree(p []string, i []string) *TreeNode {

	if len(p) == 0 || len(i) == 0 {
		return nil
	}

	var mid = -1

	val, _ := strconv.Atoi(p[0])

	root := &TreeNode{
		Val: val,
	}

	for j, c := range i {
		if c == p[0] {
			mid = j
			break
		}
	}

	root.Left = buildTree(p[1:mid+1], i[:mid])
	root.Right = buildTree(p[mid+1:], i[mid+1:])
	return root
}

func preorder(root *TreeNode) []string {
	var ans []string
	if root == nil {
		return ans
	}
	ans = append(ans, fmt.Sprintf("%d", root.Val))
	ans = append(ans, preorder(root.Left)...)
	ans = append(ans, preorder(root.Right)...)
	return ans
}

func inorder(root *TreeNode) []string {

	var ans []string
	if root == nil {
		return ans
	}

	ans = append(ans, inorder(root.Left)...)
	ans = append(ans, fmt.Sprintf("%d", root.Val))
	ans = append(ans, inorder(root.Right)...)

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

  上述实现，无法处理元素相同的情况，任务失败


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

