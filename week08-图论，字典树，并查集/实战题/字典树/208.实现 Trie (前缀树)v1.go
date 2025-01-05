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
	child []*Node
}

func MakeNode() *Node {
	return &Node{
		count: 0,
		child: make([]*Node, 26),
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

	for _, ch := range s {
		if cur.child[ch-'a'] == nil {
			if isInsert {
				cur.child[ch-'a'] = MakeNode()
			} else {
				return false
			}
		}

		cur = cur.child[ch-'a']
	}
	if isInsert {
		cur.count++
	}
	if isPrefix {
		return true
	}
	return cur.count > 0
}

/*
总结：
	16/16 cases passed (35 ms)
	Your runtime beats 17.47 % of golang submissions
	Your memory usage beats 35.75 % of golang submissions (22 MB)
*/

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end



