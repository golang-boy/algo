
/*
 * @lc app=leetcode.cn id=433 lang=golang
 * @lcpr version=20004
 *
 * [433] 最小基因变化
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minMutation(startGene string, endGene string, bank []string) int {

	/*
		输入: 给定起始和结束基因，以及基因库
		    基因序列长度为8

		输出: 找最少变化次数, 从start到end
		    找不到返-1

		1. 每个位置由ACGT之一构成,依次广度优先搜索尝试
		2. 比较起始与结束的每个位置的基因，如果不一样，尝试使用end中对应的字符替换，到结尾时判断是否在基因库中，如果在，则记录结果。如果到结尾时不在，则需要回到替换的位置，尝试另一个字符进行替换，每次替换完毕后，需要在基因库中

		3. 状态空间位4的8次方


		如何程序化？
		    1. 处理基因库，将数组变为集合，用于判断
		    2.
	*/

	hashBank := map[string]struct{}{}

	// 用于记录变化的次数
	depth := map[string]int{}

	// 处理基因库
	for _, gene := range bank {
		hashBank[gene] = struct{}{}
	}

	if _, ok := hashBank[endGene]; !ok {
		return -1
	}

	// 起始基因入队
	q := []string{}
	genes := []byte{'A', 'C', 'G', 'T'}

	q = append(q, startGene)

	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		for i := 0; i < 8; i++ { // 在8位的状态空间中进行搜索,任意一个序列都是8位
			for j := 0; j < 4; j++ {
				if s[i] != genes[j] {
					// 建一个副本ns，尝试替换第i为基因
					ns := s[:i] + string(genes[j]) + s[i+1:]

					if _, ok := hashBank[ns]; !ok {
						continue
					}

					if _, ok := depth[ns]; ok {
						continue
					}
					// 每变化一次就累加, 层数加一，上一层的层数+1
					depth[ns] = depth[s] + 1
					q = append(q, ns)
					if ns == endGene {
						return depth[ns]
					}
				}

			}
		}
	}
	return -1
}

/*
总结：

    这个题是图吗？
	    不是图，不需要构建邻接表
	能否用广搜？
	    可以, 广搜的情况是没处理完一种情况后，构造新的状态入队。只要能构造新状态，就可以使用广搜
	    整个状态空间，可以看做是一个图，只是表示图时，不需要邻接表，只需要能构造新状态即可
	搜的是什么？
	    1. 搜索的是状态空间，搜索时，通过替换进行尝试，
	    2. 每找到一个符合条件的(在基因库), 就入队, 在此基础上进行下一轮的判定,直到找到末尾基因
	    3. 要找到变化次数最小的，就是要在广度优先搜索时，找层数最小的
*/

// @lc code=end

/*
// @lcpr case=start
// "AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]\n
// @lcpr case=end

// @lcpr case=start
// "AACCGGTT"\n"AAACGGTA"\n["AACCGGTA","AACCGCTA","AAACGGTA"]\n
// @lcpr case=end

// @lcpr case=start
// "AAAAACCC"\n"AACCCCCC"\n["AAAACCCC","AAACCCCC","AACCCCCC"]\n
// @lcpr case=end

*/

