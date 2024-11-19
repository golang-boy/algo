/*
 * @Author: 刘慧东
 * @Date: 2024-11-19 14:06:20
 * @LastEditors: 刘慧东
 * @LastEditTime: 2024-11-19 17:48:29
 */
/*
 * @lc app=leetcode.cn id=30 lang=golang
 * @lcpr version=20003
 *
 * [30] 串联所有单词的子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findSubstring(s string, words []string) []int {

	ans := []int{}

	length := len([]byte(words[0]))
	subLength := length * len(words)

	query := map[string]int{}

	for _, word := range words {
		if _, ok := query[word]; ok {
			query[word] += 1
		} else {
			query[word] = 1
		}
	}

	isValid := func(sub string) bool {

		// 子串分组
		// 构造子串集合
		mapp := map[string]int{}
		for i := 0; i < len(sub); i += length {
			word := sub[i : i+length]
			if _, ok := query[word]; !ok {
				return false
			}
			if _, ok := mapp[word]; ok {
				mapp[word] += 1
			} else {
				mapp[word] = 1
			}
		}

		// 判定俩集合是否相等
		for k, v := range mapp {
			if v != query[k] {
				return false
			}
		}
		for k, v := range query {
			if v != mapp[k] {
				return false
			}
		}
		return true
	}

	for i := 0; i+subLength <= len(s); i++ {
		sub := s[i : i+subLength]
		if isValid(sub) {
			ans = append(ans, i)
		}
	}
	return ans
}

/*
总结：
    1. 遍历字符串，每次为sublength长度
	2. 校验是否有效
	   1. 分组
	   2. 构造子串集合
	   3. 判定俩集合是否相等
	3. 有效记录，无效下一个


	老师说的部分枚举+滑动窗口 (没明白怎么回事？)
	   1. 枚举时，不用重新构造子串集合, 使用移动窗口
	   2. 优化的是内层循环
```
func findSubstring(s string, words []string) (ans []int) {
    ls, m, n := len(s), len(words), len(words[0])
    for i := 0; i < n && i+m*n <= ls; i++ {
        differ := map[string]int{}

        // 从当前位置i开始，循环m(words中词的数量)次，每次取长度为n(words中词的长度)的子串, 放入differ中,有重复时++
		// 相当于子串分组
        for j := 0; j < m; j++ {
            differ[s[i+j*n:i+(j+1)*n]]++
        }

		// 子串中的词与words匹配，匹配到的--，如果值为0，则删除
		// 看在不在words序列中
        for _, word := range words {
            differ[word]--
            if differ[word] == 0 {
                delete(differ, word)
            }
        }

		//  m * n +1 表示子串的长度
        for start := i; start < ls-m*n+1; start += n {
            if start != i {
                word := s[start+(m-1)*n : start+m*n] // 最后一个词？ 为什么？ 因为每次循环，都是取长度为n的子串, 窗口右侧
                differ[word]++
                if differ[word] == 0 {
                    delete(differ, word)
                }
                word = s[start-n : start] // 窗口左侧词，前一个
                differ[word]--
                if differ[word] == 0 {
                    delete(differ, word)
                }
            }

			// 第一次时，直接判断，后续判断differ是否为空，为空表示都匹配上了
            if len(differ) == 0 {
                ans = append(ans, start)
            }
        }
    }
    return
}
	   ````




	时间复杂度分析
	  1. 遍历字符串， O(len(s))
	  2. 每次遍历时，需要遍历子串，子串长度为words的长度
	  3. O(len(s)*len(words))

*/

// @lc code=end

/*
// @lcpr case=start
// "barfoothefoobarman"\n["foo","bar"]\n
// @lcpr case=end

// @lcpr case=start
// "wordgoodgoodgoodbestword"\n["word","good","best","good"]\n
// @lcpr case=end

// @lcpr case=start
// "wordgoodgoodgoodbestword"\n["word","good","best","word"]\n
// @lcpr case=end

// @lcpr case=start
// "barfoofoobarthefoobarman"\n["bar","foo","the"]\n
// @lcpr case=end

*/

