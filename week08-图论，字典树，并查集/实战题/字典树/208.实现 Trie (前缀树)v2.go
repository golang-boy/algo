/*
 * @lc app=leetcode.cn id=208 lang=golang
 * @lcpr version=20004
 *
 * [208] 实现 Trie (前缀树)
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type Trie struct {
	root *Node
}

type Node struct {
	count int
	// child []*Node
	child map[rune]*Node
}

func MakeNode() *Node {
	return &Node{
		count: 0,
		child: make(map[rune]*Node, 26),
	}
}

func Constructor() Trie {
	return Trie{
		root: MakeNode(),
	}

}

func (this *Trie) Insert(word string) {

	this.find(word, true, false)

}

func (this *Trie) Search(word string) bool {
	return this.find(word, false, false)
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix, false, true)

}

func (this *Trie) find(s string, isInsert, isPrefix bool) bool {
	cur := this.root

	// 从根出发，寻找每个字符
	for _, ch := range s {
		// 如果找不到，判断是否需要插入，或者返回
		if cur.child[ch] == nil {
			if isInsert {
				cur.child[ch] = MakeNode()
			} else {
				return false
			}
		}
		// 找到当前字符，继续看下一个字符
		cur = cur.child[ch]
	}
	// 循环完毕，如果是插入，则在当前节点上对count进行加一操作
	if isInsert {
		cur.count++
	}

	// 如果是判断前缀的，则返回true
	if isPrefix {
		return true
	}

	// 在找到的节点，判断count是否大于0
	//    如果大于0, 则说明有该词，否则没有
	return cur.count > 0
}

/*
总结：

	整个过程中就是查找单个字符的过程，找不到是看是否需要插入等
	    在trie树中层层判断字符串s中对应的字符, 使用数组更快一些, 但使用map针对不同类型的字符更方便一些

	16/16 cases passed (35 ms)
	Your runtime beats 7.72 % of golang submissions
	Your memory usage beats 5.03 % of golang submissions (30.2 MB)
*/

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end



