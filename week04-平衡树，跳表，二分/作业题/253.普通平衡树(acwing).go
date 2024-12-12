package main

// https://www.acwing.com/problem/content/255/

/*

实现一个普通的平衡树, 提供一下操作
1. 插入数值 x。
2. 删除数值 x(若有多个相同的数，应只删除一个)。
3. 查询数值 x的排名(若有多个相同的数，应输出最小的排名)。
4. 查询排名为 x的数值。
5. 求数值 x的前驱(前驱定义为小于 x的最大的数)。
6. 求数值 x的后继(后继定义为大于 x 的最小的数)。

输入: n,表示操作个数

*/

type Node struct {
	left  *Node
	right *Node

	key, val  int
	cnt, size int
}

type Treap struct {
}

func (t *Treap) Insert(data int)             {}
func (t *Treap) Remove(data int)             {}
func (t *Treap) GetRankByVal(target int) int { return 0 }
func (t *Treap) GetValByRank(data int) int   {}
func (t *Treap) GetPre(target int) int       {}
func (t *Treap) GetNext(target int) int      {}

func main() {

	/*



	 */

}
